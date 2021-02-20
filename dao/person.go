package dao

import (
	"github.com/didi/gendry/scanner"
	"go_web_starter/database"
	"go_web_starter/util"
	"time"
)

type Persons struct {
	Id         int    `ddb:"id"`
	UserName   string `ddb:"username"`
	Password   string `ddb:"password"`
	CreateTime time.Time `ddb:"create_time"`
	UpdateTime time.Time `ddb:"update_time"`
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
	rows, err := p.Db.MySQL.Query("select * from user")
	defer rows.Close()
	if err != nil {
		util.Log.Error(err)
	}
	scanner.Scan(rows, &data)
	return data
}
