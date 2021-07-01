package main

import (
	"bile-go-server/core/handle"
	"bile-go-server/core/login"
	"bile-go-server/core/register"
	"encoding/json"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/register", register.RegisterHandle)
	http.HandleFunc("/login", login.LoginHandle)
	http.HandleFunc("/a", testH)
	http.ListenAndServe("127.0.0.1:8080", new(handle.AuthMidServer))

}

func testH(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var syn sync.WaitGroup
	syn.Add(1)
	mm := r.Form.Get("a")
	go func() {
		if mm == "a" {
			enc := json.NewEncoder(w)
			enc.Encode("sddsdsdasdsdasdsadsddsada")
			w.Write([]byte("ssaddsadsdsad"))
		}
		syn.Done()
	}()
	syn.Wait()
}
