package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// 如果这里不传参的话， 那么就相当于是闭包， 外部的i就是自由变量， 在main这个goroutine结束之后
		// i的值为10， 此时再执行 a[i]++ 会 out of range
		// race condition i
		go func(i int) { // 主程序还在往下面跑， 只是并发地开了一个函数
			for {
				//fmt.Printf("Hello from "+"goroutine %d\n", i)
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
