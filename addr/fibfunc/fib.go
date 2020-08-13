package main

import (
	"bufio"
	"fmt"
	"io"
	"learn/addr/fib"
	"strings"
)

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
	var f intGen = fib.Fibonacci()
	printFibContents(f)

	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
}