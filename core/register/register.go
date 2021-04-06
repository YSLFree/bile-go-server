package register

import (
	"bile-go-server/code"
	"bile-go-server/common/datautil"
	"bile-go-server/common/db"
	"bile-go-server/entity"
	"encoding/json"
	"net/http"
	"strings"
)

//RegisterHandle 。。
func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	var data = new(entity.BaseEntity)
	if db.GetMysqlCon() == nil {
		data.Code = code.StatusFailed
		data.State = false
		data.Data = "服务端数据库异常"
		enc.Encode(data)
		return
	} 
	r.ParseForm()
	if r.Method == http.MethodPost {
		//types := r.Header.Get("Content-Type")
		account := r.Form.Get("account")
		password := r.Form.Get("password")
		if datautil.VerifyMobileFormat(account) {
			data.Code = code.StatusFailed
			data.State = true
			data.Data = "请输入合法的手机号码"
			enc.Encode(data)
			return
		}
		if len(password) < 6 {
			data.Code = code.StatusFailed
			data.State = true
			data.Data = "密码长度不能小于6"
			enc.Encode(data)
			return
		}
		isV := strings.Contains(account, "#") ||
			strings.Contains(account, "or") ||
			strings.Contains(account, "==") ||
			strings.Contains(account, " ") ||
			strings.Contains(password, "#") ||
			strings.Contains(password, "or") ||
			strings.Contains(password, "==") ||
			strings.Contains(password, " ")
		if isV {
			data.Code = code.StatusFailed
			data.State = true
			data.Data = "请不要输入非法字符"
			enc.Encode(data)
			return
		}
		sqlQuer := "select * from logininfo where account=? and password=?"
		var queryData = new(entity.Login)
		er := db.GetMysqlCon().Get(queryData, sqlQuer, account, password)
		if er == nil { //用户存在
			data.Code = code.StatusFailed
			data.State = true
			data.Data = "该账号用户已经存在，请直接登录"
			enc.Encode(data)
			return
		}
		tx, ter := db.GetMysqlCon().Begin()
		if ter != nil {
			data.Code = code.StatusFailed
			data.State = true
			data.Data = ter
			enc.Encode(data)
			return
		}
		sqlInsert := "insert into logininfo(account,password) values(?,?)"
		res, err := tx.Exec(sqlInsert, account, password)
		if err != nil {
			data.Code = code.StatusFailed
			data.State = true
			data.Data = err.Error()
			enc.Encode(data)
			tx.Rollback()
			return
		}
		res.LastInsertId()
		userId := datautil.CreateMD5(account + password)
		updateLoginInfoSql := "update  logininfo set user_id=? where account=? and password=?"
		res, err = tx.Exec(updateLoginInfoSql, userId, account, password)
		if err != nil {
			data.Code = code.StatusFailed
			data.State = true
			data.Data = err.Error()
			enc.Encode(data)
			tx.Rollback()
			return
		}
		res.LastInsertId()

		updateUserInfoSql := "insert into user_info(user_id) values(?)"
		res, err = tx.Exec(updateUserInfoSql, userId)
		if err != nil {
			data.Code = code.StatusFailed
			data.State = true
			data.Data = err.Error()
			enc.Encode(data)
			tx.Rollback()
			return
		}
		res.LastInsertId()
		tx.Commit()
		data.Code = code.StatusSuccess
		data.State = true
		data.Data = "注册成功"
		enc.Encode(data)
	} else {
		data.Code = code.StatusFailed
		data.State = true
		data.Data = "request type error"
		enc.Encode(data)
	}
}
