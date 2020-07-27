package main

import (
	"fmt"
	"learn/retriever/mooc"
	"learn/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}
func main() {
	var r Retriever = mooc.Retriever{"this is fake imooc.com"}
	inspect(r)

	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// type assertion
	if moocRetriever, ok := r.(mooc.Retriever); ok {
		fmt.Println(moocRetriever.Contents)
	} else {
		fmt.Println("not a mooc retriever")
	}
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mooc.Retriever:
		fmt.Println("contents:", v.Contents)
	case real.Retriever:
		fmt.Println("userAgent:", v.UserAgent)
	}
}
