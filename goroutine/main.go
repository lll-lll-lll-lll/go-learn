package main

import (
	"fmt"
	"time"
)

func download(u string) {
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()
	for i := 1; i <= 100; i++ {
		u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
		download(u)
	}
	fmt.Printf("%v\n", time.Since(before))
}
