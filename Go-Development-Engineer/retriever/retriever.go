package main

import (
	"Learn_go/Go-Development-Engineer/retriever/mock"
	"Learn_go/Go-Development-Engineer/retriever/real"
	"fmt"
	"time"
)

// Retriever 定义了一个接口，实现者需要有 Get方法
type Retriever interface {
	Get(url string) string
}

// 直接通过接口调用Get方法
func download(r Retriever) string {
	return r.Get("https://www.acwing.com")
}

func main() {
	var r Retriever = &mock.Retriever{Contents: "This is a fake url"}
	inspect(r)
	//fmt.Println(download(r))
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock Retriever!")
	}
	fmt.Println(realRetriever.TimeOut)
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
