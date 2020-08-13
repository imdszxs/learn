package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)
//斐波那契数列
func fibonacci() func() int {
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(
	p []byte) (n int,err error){
	next := g()
	if next > 1000 {
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}

func printFibContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fibonacci()
	printFibContents(f)

	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
}