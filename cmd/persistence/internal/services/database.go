package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/persistence/internal/config"
	"github.com/FelipeMandelli/ProjetoInterativoV/pkg/database"
	entities "github.com/FelipeMandelli/ProjetoInterativoV/pkg/entities"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(provider *Provider) error {

	db, err := gorm.Open(mysql.Open(createDBConnString(
		viper.GetInt(config.DBPortKey),
		viper.GetString(config.DBHostKey),
		viper.GetString(config.DBUserKey),
		viper.GetString(config.DBPassKey),
		viper.GetString(config.DBNameKey),
	)), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&entities.Attendance{},
		&entities.Professor{},
		&entities.Student{},
		&entities.Subject{},
	)
	if err != nil {
		return err
	}

	provider.DB = db
	provider.DbIsON = true

	return nil
}

func PersistAtendance(p *Provider, teacherTag string, studentTags []string, receivedTime time.Time) error {
	var errs error

	professor, err := database.FindProfessorByID(p.DB, teacherTag)
	if err != nil {
		return fmt.Errorf("could not find specified professor: [%w]", err)
	}

	attendedStudents := make(map[string]*entities.Student)

	for _, t := range studentTags {
		s, err := database.FindStudentByID(p.DB, t)
		if err != nil {
			errors.Join(errs, fmt.Errorf("could not find student with tag %s [%w]", t, err))
		}

		attendedStudents[s.IDBiometry] = s
	}
	if err != nil {
		return fmt.Errorf("could not find specified strudents: [%w]", err)
	}

	p.Log.Sugar().Infof("checking for subject using professor [%s], year [%s], weekday [%s], schedule [%s]", professor.IDBiometry, strconv.Itoa(receivedTime.Year()), strconv.Itoa(getWeekDay(receivedTime)), strconv.Itoa(getSchedule(receivedTime)))

	subject, err := database.FindSubjectByProfessorAndWeekdayAndSchedule(p.DB, professor.IDBiometry, strconv.Itoa(receivedTime.Year()), strconv.Itoa(getWeekDay(receivedTime)), strconv.Itoa(getSchedule(receivedTime)))
	if err != nil {
		return fmt.Errorf("could not find specified subject for teacher [%s] and receivedTime [%s]: [%w]", teacherTag, receivedTime.String(), err)
	}

	att, ok, err := database.FidExistentAttendace(p.DB, professor.IDBiometry, receivedTime.Format(config.DateFormater), strconv.Itoa(getSchedule(receivedTime)))
	if err != nil {
		return fmt.Errorf("could not check existing attendace for teacher [%s] and receivedTime [%s]: [%w]", teacherTag, receivedTime.String(), err)
	}

	if !ok {
		att.Date = receivedTime.Format(config.DateFormater)
		att.Schedule = subject.Schedule
		att.SubjectID = strconv.Itoa(int(subject.ID))
		att.ProfessorID = professor.IDBiometry
	}

	studentsEnrolled := strings.Split(subject.StudentsEnrolled, ",")

	var attendanceList []string
	var absenceList []string

	for _, s := range studentsEnrolled {

		if _, found := attendedStudents[s]; found {
			attendanceList = append(attendanceList, s)

			continue
		}

		absenceList = append(absenceList, s)
	}

	if ok {
		existingAttendanceList := strings.Split(att.AttendanceList, ",")

		for i, abs := range absenceList {
			for _, att := range existingAttendanceList {
				if abs == att {
					attendanceList = append(attendanceList, absenceList[i])
					absenceList = append(absenceList[:i], absenceList[i+1:]...)
				}
			}
		}
	}

	attendanceListStr := strings.Join(checkDuplicates(attendanceList), ",")
	absenceListStr := strings.Join(checkDuplicates(absenceList), ",")

	att.AttendanceList = attendanceListStr
	att.AbsenceList = absenceListStr

	err = database.SaveAttendance(p.DB, att)
	if err != nil {
		return fmt.Errorf("could not save attendance [%+v]: [%w]", att, err)
	}

	return nil
}

func PersistStudentRegistry(p *Provider, reg entities.Registry) error {
	student, err := reg.ToStudent()
	if err != nil {
		return fmt.Errorf("could not parse to student: [%w]", err)
	}

	err = database.SaveStudent(p.DB, student)
	if err != nil {
		return fmt.Errorf("error saving professor: [%w]", err)
	}

	return nil
}

func PersistProfessorRegistry(p *Provider, reg entities.Registry) error {
	professor, err := reg.ToProfessor()
	if err != nil {
		return fmt.Errorf("could not parse to student: [%w]", err)
	}

	err = database.SaveProfessor(p.DB, professor)
	if err != nil {
		return fmt.Errorf("error saving professor: %w", err)
	}

	return nil
}

func PersistSubjectRegistry(p *Provider, subReg entities.SubjectRegistry) error {
	err := database.SaveSubject(p.DB, subReg.ToSubject())
	if err != nil {
		return fmt.Errorf("error saving subject: %w", err)
	}

	return nil
}

func createDBConnString(port int, host, username, password, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, name)
}

func getSchedule(r time.Time) (s int) {
	startSchedule1 := time.Date(r.Year(), r.Month(), r.Day(), 19, 0, 0, 0, r.Location())
	endSchedule1 := time.Date(r.Year(), r.Month(), r.Day(), 21, 0, 0, 0, r.Location())
	startSchedule2 := time.Date(r.Year(), r.Month(), r.Day(), 21, 0, 0, 1, r.Location())
	endSchedule2 := time.Date(r.Year(), r.Month(), r.Day(), 23, 0, 0, 0, r.Location())

	if r.After(startSchedule1) && r.Before(endSchedule1) {
		s = 1

		return
	}

	if r.After(startSchedule2) && r.Before(endSchedule2) {
		s = 2

		return
	}

	return
}

func getWeekDay(r time.Time) int {
	switch r.Weekday() {
	case time.Sunday:
		return 1
	case time.Monday:
		return 2
	case time.Tuesday:
		return 3
	case time.Wednesday:
		return 4
	case time.Thursday:
		return 5
	case time.Friday:
		return 6
	case time.Saturday:
		return 7
	default:
		return 0
	}
}

func checkDuplicates(slice []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, val := range slice {
		if !encountered[val] {
			encountered[val] = true
			result = append(result, val)
		}
	}

	return result
}
