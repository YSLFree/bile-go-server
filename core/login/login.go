package login

import (
	"bile-go-server/code"
	"bile-go-server/common/datautil"
	"bile-go-server/common/db"
	"bile-go-server/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

//LoginHandle 登录
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3*time.Second)
	fmt.Println("id=",GetGoroutineID())
	var data = new(entity.BaseEntity)
	enc := json.NewEncoder(w)
	r.ParseForm()
	userToken := r.Header.Get("usertoken")
	deviceToekn := r.Header.Get("devicetoken")
	if db.GetMysqlCon() == nil {
		data.Code = code.StatusFailed
		data.State = false
		data.Data = "服务端数据库异常"
		enc.Encode(data)
		return
	}

	if r.Method == http.MethodPost {
		//types := r.Header.Get("Content-Type")
		if len(userToken) > 0 { //token存在，就用token登录
			if len(deviceToekn) < 0 {
				data.Code = code.StatusIdentityExpired
				data.State = true
				data.Data = "无效身份"
				enc.Encode(data)
				return
			} else {
				//根据token查找redis中是否缓存用户信息，有就直接去除返回给用户，没有就在mysql中查询
				result, err := redis.String(db.GetRedisCon().Do("GET", userToken))
				if err == nil && len(result) > 0 { //查询到了用户信息
					userInfo := new(entity.UserInfos)
					json.Unmarshal([]byte(result), userInfo)
					data.Code = code.StatusSuccess
					data.State = true
					data.Data = userInfo
					enc.Encode(data)
					db.GetRedisCon().Close()
					return
				}
				//redis中没有查询到用户信息，在mysql中查询
				sqlQuer := "select * from logininfo where usertoken=? and devicetoken=?"
				var queryData = new(entity.Login)
				err = db.GetMysqlCon().Get(queryData, sqlQuer, userToken, deviceToekn)
				if err == nil { //查询到用户信息
					sqlUser := "select * from user_info where user_id=?"
					var userInfo = new(entity.UserInfos)
					err := db.GetMysqlCon().Get(userInfo, sqlUser, queryData.User_Id)
					if err == nil {
						data.Code = code.StatusSuccess
						data.State = true
						data.Data = userInfo
						enc.Encode(data)
						byt, e := json.Marshal(userInfo)
						s := string(byt)
						fmt.Println(s)
						if e == nil {
							db.GetRedisCon().Do("SET", userToken, s)
						}
						db.GetRedisCon().Close()
						return
					} else {
						data.Code = code.StatusQueryDataError
						data.State = true
						data.Data = "查询数据异常"
						enc.Encode(data)
						return
					}
				} else { //未查询到用户信息
					data.Code = code.StatusTokenInvalid
					data.State = true
					data.Data = "身份过期请重新登录"
					enc.Encode(data)
				}
			}
		} else { //用户密码登录
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
				strings.Contains(account, " ") ||
				strings.Contains(account, "or") ||
				strings.Contains(account, "==") ||
				strings.Contains(password, "#") ||
				strings.Contains(password, " ") ||
				strings.Contains(password, "or") ||
				strings.Contains(password, "==")

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
			if er == nil { //查询到用户
				sqlUser := "select * from user_info where user_id=?"
				var userInfo = new(entity.UserInfos)
				err := db.GetMysqlCon().Get(userInfo, sqlUser, queryData.User_Id)
				if err == nil {
					data.Code = code.StatusSuccess
					data.State = true
					data.Data = userInfo
					enc.Encode(data) //查询成功后，将用户信息返回
					//更新token到用户登录表中
					tx, _ := db.GetMysqlCon().Begin()
					updateLoinInfo := "update logininfo set usertoken=? ,devicetoken=? where user_id=?"
					userToken = datautil.CreateMD5(account + time.Now().String())
					res, err := tx.Exec(updateLoinInfo, userToken, deviceToekn, userInfo.User_Id)
					if err != nil {
						data.Code = code.StatusFailed
						data.State = true
						data.Data = err.Error()
						enc.Encode(data)
						tx.Rollback()
						fmt.Scan()
						return
					}
					res.RowsAffected()
					tx.Commit()
					return
				}
			} else { //未查询到yong hu
				data.Code = code.StatusAccountError
				data.State = true
				data.Data = "用户名或密码错误"
				enc.Encode(data)
			}
		}
	} else {
		data.Code = code.StatusFailed
		data.State = true
		data.Data = "request type error"
		enc.Encode(data)
	}
}
func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
