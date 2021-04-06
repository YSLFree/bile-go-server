package db

import (
	_ "database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

var mysqlDB *sqlx.DB
var pool *redis.Pool

//Init ,
func GetMysqlCon() *sqlx.DB {
	// Db, err = sqlx.Open(Driver, MYSQL_DNS) //检查DNS语句是否正确
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// if err = Db.Ping(); err != nil { //Ping是真正带入用户名密码去登录的
	// 	return err
	// }
	// Db.SetMaxOpenConns(20)//连接池最大连接数
	// Db.SetMaxIdleConns(2)//连接池最大空闲连接数
	// return nil
	if mysqlDB == nil {
		var err error
		mysqlDB, err = sqlx.Connect(Driver, MYSQL_DNS) //检查DNS语句,并且执行Ping方法
		if err != nil {
			fmt.Println(err)
			return nil
		}
		mysqlDB.SetMaxOpenConns(20)                  //连接池最大连接数
		mysqlDB.SetMaxIdleConns(2)                   //连接池最大空闲连接数
		mysqlDB.SetConnMaxIdleTime(time.Second * 30) //连接最大空闲时间
		mysqlDB.SetConnMaxLifetime(time.Second * 15) //最大连接时间
	}
	return mysqlDB
}

func GetRedisCon() redis.Conn {
	if pool == nil {
		pool = &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial(
					"tcp", "127.0.0.1:6379",
					//redis.DialUsername("username"),
					//redis.DialPassword("123456"),
					redis.DialConnectTimeout(5*time.Second), //连接超时时间
					redis.DialReadTimeout(5*time.Second),    //读超时时间
					redis.DialWriteTimeout(5*time.Second),   //写超时时间
				)
			},
			MaxIdle:         1,
			MaxActive:       50,
			IdleTimeout:     time.Minute * 10,
			Wait:            true,
			MaxConnLifetime: 0,
		}
	}
	if pool != nil {
		return pool.Get()
	}
	return nil
}
