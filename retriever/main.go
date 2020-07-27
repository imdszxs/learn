package main

import (
	"fmt"
	"learn/retriever/mooc"
	"learn/retriever/real"
	"time"
)

const url = "https://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(p Poster)  {
	p.Post(url,
		map[string]string{
			"name":"imdszxs",
			"course":"golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mooc.Retriever:
		fmt.Println("contents:", v.Contents)
	case real.Retriever:
		fmt.Println("userAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever = &mooc.Retriever{"this is fake imooc.com"}
	inspect(r)

	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// type assertion
	if moocRetriever, ok := r.(*mooc.Retriever); ok {
		fmt.Println(moocRetriever.Contents)
	} else {
		fmt.Println("not a mooc retriever")
	}

	fmt.Println("try a session")
	s := mooc.Retriever{"this is fake imooc.com"}
	fmt.Println(session(&s))
}


