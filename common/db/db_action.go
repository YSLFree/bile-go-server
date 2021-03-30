package db

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db db对象
var Db *sqlx.DB
var err error

var dbEnable bool

//Init ,
func Init() {
	Db, err = sqlx.Open(Driver, "root:123456@tcp(127.0.0.1:3306)/bile?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
	} else {
		dbEnable = true
	}
}

//pingDb ,ping
func PingDb() bool {
	if dbEnable {
		if Db.Ping() == nil {
			return true
		}
	}
	return false
}
