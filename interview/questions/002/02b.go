package main

import "fmt"

//查找字符串是否所有字符串不相同

func IsRepeated(s string) bool {

	if len(s) > 3000 {
		return false
	}
	for key, value := range s {
		if value > 256 {
			return false
		}

		for kk, vv := range s {
			if vv == value && kk != key {
				fmt.Println(string(vv))
				return true
			}
		}
	}
	return false

}
func main() {
	s := "1dr2jala"
	if IsRepeated(s) {
		fmt.Println("重复")
	} else {
		fmt.Println("不重复")
	}

}
