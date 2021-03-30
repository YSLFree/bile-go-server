package Login

import (
	"bile-go-server/common/db"
	"bile-go-server/common/nets"
	users_info "bile-go-server/users"
	_ "database/sql"
	"encoding/json"
	_ "errors"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
var err error
var sqls *db.SQL
var dbEnable bool

//LoginHandler 登录
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data users_info.ResultData
	if er := r.ParseForm(); er != nil {
		data.Code = nets.ParamsFormErrorFormClient
		data.State = false
		data.Data = "服务器解析来自客户端的参数时错误"
		_, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(err.Error()))
			return
		}

	} //解析参数，默认不会解析
	if r.Method == "GET" { //客户端get请求
		account := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("from GET %v,%T", account, account)
		Login(account, password)
	} else if r.Method == "POST" { //客户端post请求
		account := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("from POST %v,%T", account, account)
		Login(account, password)
	}
	//验证用户名密码，如果成功则header里返回session，失败则返回StatusUnauthorized状态码

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("服务器连接成功\n"))
}

//Login 用户登陆
func Login(account string, password string) (interface{}, error) {

	res, err := Db.Exec(sqls.RegisterSQL(), account, password)
	if err != nil {
		fmt.Println("Register Error: " + err.Error())
		return nil, err
	}
	id, _ := res.LastInsertId()
	if id > 0 {

	}
	return nil, nil
}