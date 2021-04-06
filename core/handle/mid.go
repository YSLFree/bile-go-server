package handle

import (
	"net/http"
)

type AuthMidServer struct {
	Next http.Handler
}

func (authSer *AuthMidServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if authSer.Next == nil {
		authSer.Next = http.DefaultServeMux
	}

	auth := r.Header.Get("auth")
	if auth != "" {
		authSer.Next.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
