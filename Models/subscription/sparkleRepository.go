package subscription

import (
	"log"
	database "github.com/engajerest/auth/utils/dbconfig"
)

const (
	getAllSparkleQuery = ""
)

func GetAllSparkles() ([]Category, []SubCategory, []Packages) {
	stmt, err := database.Db.Prepare(getAllSparkleQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return nil,nil,nil
}