package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	con := "DATABASE=bludb;HOSTNAME=98538591-7217-4024-b027-8baa776ffad1.c3n41cmd0nqnrk39u98g.databases.appdomain.cloud;PORT=30875;PROTOCOL=TCPIP;UID=rqx78693;PWD=1CY68qga6Ses37LI;SECURITY=SSL;"
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	fmt.Println("FUNCIONOU BUCETA")
}
