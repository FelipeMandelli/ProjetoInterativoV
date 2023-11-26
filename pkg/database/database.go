package database

import (
	"fmt"

	"github.com/FelipeMandelli/ProjetoInterativoV/pkg/entities"
	"gorm.io/gorm"
)

func SaveStudent(db *gorm.DB, student *entities.Student) error {
	err := db.Save(student).Error
	if err != nil {
		return fmt.Errorf("error saving student: [%w]", err)
	}

	return err
}

func SaveProfessor(db *gorm.DB, professor *entities.Professor) error {
	err := db.Save(professor).Error
	if err != nil {
		return fmt.Errorf("error saving professor: [%w]", err)
	}

	return err
}

func SaveSubject(db *gorm.DB, subject *entities.Subject) error {
	err := db.Save(subject).Error
	if err != nil {
		return fmt.Errorf("error saving subject: [%w]", err)
	}

	return err
}

func SaveAttendance(db *gorm.DB, attendance *entities.Attendance) error {
	err := db.Save(attendance).Error
	if err != nil {
		return fmt.Errorf("error saving attendance: [%w]", err)
	}

	return err
}

func FindStudentByID(db *gorm.DB, id string) (*entities.Student, error) {
	var student entities.Student

	err := db.Find(&student, "id_biometry", id).Error
	if err != nil {
		return nil, fmt.Errorf("could not find Student by id: [%w]", err)
	}

	return &student, nil
}

func FindProfessorByID(db *gorm.DB, id string) (*entities.Professor, error) {
	var professor entities.Professor

	err := db.Find(&professor, "id_biometry", id).Error
	if err != nil {
		return nil, fmt.Errorf("could not find Professor by id: [%w]", err)
	}

	return &professor, nil
}

func FindSubjectByProfessorAndWeekdayAndSchedule(db *gorm.DB, professor, year, weekday, schedule string) (*entities.Subject, error) {
	var subject entities.Subject

	err := db.Where("professor_id = ? AND week_day = ? AND (schedule = ? OR ?) AND reference_year = ?", professor, weekday, schedule, "3", year).Find(&subject).Error
	if err != nil {
		return nil, fmt.Errorf("could not find subject by given info: [%w]", err)
	}

	return &subject, nil
}

func FindExistentAttendace(db *gorm.DB, professorID, date, schedule string) (*entities.Attendance, bool, error) {
	var att entities.Attendance

	result := db.Where("professor_id = ? AND schedule = ? AND date = ?", professorID, schedule, date).First(&att)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, false, fmt.Errorf("could not check the db for existance: [%w]", result.Error)
	}

	if result.RowsAffected == 1 {
		return &att, true, nil
	}

	if result.RowsAffected > 1 {
		return nil, true, fmt.Errorf("founded more than one attendace for teacher [%s] date [%s] and schedule [%s]. Will not process", professorID, date, schedule)
	}

	return &entities.Attendance{}, false, nil
}
