package main

import (
	"fmt"
)

func main() {
	byt := []byte{123, 34, 105, 100, 24, 58, 49, 44, 34, 117,97,97,97}
	str := string(byt)
	x:= []byte(`\`)
	// s1,e1:= redis.String(con.Do("set", "name" ,"ysl "))
	fmt.Println(str)
	fmt.Println(x)

}
