package main

import (
	"bufio"
	"fmt"
	"learn/addr/fib"
	"os"
)

func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error")
	return
	fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f := fib.Fibonacci()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
