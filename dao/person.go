package dao

import (
	"github.com/didi/gendry/scanner"
	"log"
)

type Persons struct {
	Id        string `ddb: "Id_P"`
	FirstName string `ddb: "FirstName"`
	LastName  string `ddb: "LastName"`
	City      string `ddb: "City"`
	Columns   int    `ddb: "column"`
	Address   string `ddb: "Address"`
}

func GetPerson() []Persons {
	var data []Persons
	rows, err := Db.Query("select * from Persons")
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	scanner.Scan(rows, &data)
	return data
}
