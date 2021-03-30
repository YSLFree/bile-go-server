package datautil

import (
	"math/rand"
	"time"
)

//GetRandomString 随机生成指定位数的大写字母,小写字母和数字的组合
//
func GetRandomString(strOutPutLengh int, dev string, devPostion int) string {
	str := `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ|`
	bytes := []byte(str)
	result := []byte{}
	bsa := []byte(dev)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < strOutPutLengh; i++ {
		if i%devPostion == 0 && i != 0 {
			result = append(result, bsa[0])
		} else {
			result = append(result, bytes[r.Intn(len(bytes))])
		}
	}
	return string(result)
}
