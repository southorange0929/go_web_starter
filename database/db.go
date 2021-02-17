package database

import "database/sql"

type DbConnection struct {
	MySQL *sql.DB
}

var Db = &DbConnection{
	MySQL: nil,
}

func Init() {
	// mysql初始化
	mysqlConnection := MySQLStart()
	Db.MySQL = mysqlConnection
}
