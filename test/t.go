package main

import (
	"fmt"
	"net/http"
)

// 创建一个新的logger实例。可以创建任意多个。
//var log = logrus.New()

func main() {
	fmt.Println("main--start !")
	http.HandleFunc("/", test)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}

}
func test(w http.ResponseWriter, r *http.Request) {
	if r != nil {
		r.ParseForm()
		if r.Method == http.MethodPost {
			w.Write([]byte("method=post  return data=" + getStr()))
		} else if r.Method == http.MethodGet {
			w.Write([]byte("method=get  return data=" + getStr()))
		}
	}
}

func getStr() string {
	return "{\n    \"result\": 0,\n    \"content\": \"测试服务器，用于个人学习、开发使用\"\n}"

}
