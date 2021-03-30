package register

import (
	"bile-go-server/common/db"
	"bile-go-server/common/nets"
	users_info "bile-go-server/users"
	_ "database/sql"
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


//RegisterHandler 。。
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//验证用户名密码，如果成功则header里返回session，失败则返回StatusUnauthorized状态码
	w.WriteHeader(http.StatusOK)
	var result users_info.ResultData = new(users_info.ResultData)
	if result.Code == 1 {
		return
	}
	// var code int32
	// var state bool
	// user := r.Form.Get("user")
	// pwd := r.Form.Get("pass")

	// if formatutils.VerifyMobileFormat(user) || formatutils.VerifyEmailFormat(user) {

	// }
	if (r.Form.Get("user") == "admin") && (r.Form.Get("pass") == "888") {
		w.Write([]byte("hello,验证成功！"))
	} else {
		w.Write([]byte("hello,验证失败了！"))
	}
}


//Register 用户注册
func Register(account string, password string) int32 {
	
	var user  users_info.LoginInfo
	e1 := Db.Get(&user,sqls.SearchSQL(), account, password)
	if e1 == nil { //用户存在
		return nets.UserExist
	}
	res, err := Db.Exec(sqls.RegisterSQL(), account, password)
	if err != nil {
		fmt.Println("register failed! error=",err.Error())
		return nets.SqlActionError
	}
	id, _ := res.LastInsertId()
	fmt.Println("register successed! ->id: ", id)
	return 0
}
