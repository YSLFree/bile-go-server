package datautil

import (
	"crypto/md5"
	"fmt"
	"io"
)

func CreateMD5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
