/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 20:33
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
		fmt.Println(value)

		/*
			这段代码可能正常输出1,可能触发panic,因为go select在没有default情况下会选择执行
		*/
	}
}
