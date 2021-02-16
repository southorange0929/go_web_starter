package dao

import (
	"database/sql"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"go_web_starter/config"
	"log"
	"strconv"
)

var Db *sql.DB

func Init() {
	port, err := strconv.Atoi(config.Config.MySQL.Port)
	if err != nil {
		log.Print(err)
	}
	db, err := manager.New(config.Config.MySQL.Database, config.Config.MySQL.User, config.Config.MySQL.Password, config.Config.MySQL.Host).Set().Port(port).Open(true)
	if err != nil {
		log.Print(err)
	}
	Db = db
}
