package main

import (
	"fmt"
	"regexp"
)

const text = `my email is ccmouse@gmail.com
email1 is yydgs@abc.top
email2 is    kkk@qq.com
email3 is   ddd@abc.com.cn`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindString(text)
	match_name := re.FindAllStringSubmatch(text, -1)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
	fmt.Println(match_name)
}
