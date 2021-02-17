package dao

import (
	"github.com/didi/gendry/scanner"
	"go_web_starter/database"
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

type PersonDao struct {
	Db *database.DbConnection
}

var personDaoInstance *PersonDao = nil

// 单例
func NewPersonDao() *PersonDao {
	if personDaoInstance == nil {
		personDaoInstance = &PersonDao{
			Db: database.Db,
		}
	}
	return personDaoInstance
}

func (p *PersonDao) GetPerson() []Persons {
	var data []Persons
	rows, err := p.Db.MySQL.Query("select * from Persons")
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	scanner.Scan(rows, &data)
	return data
}
