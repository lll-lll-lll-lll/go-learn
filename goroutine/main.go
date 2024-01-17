package main

import (
	"fmt"
	"sync"
	"time"
)

func download(u string) {
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()
	// goroutineの生成量を10個に抑えることで、サーバへの負荷を抑える
	limit := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		// goroutineの中で追加してしまうと、任意のタイミングで起動するgoroutineが起動する前にメインのgoroutineが
		//最後に到達してしまい、処理が終わってしまう。なので、goroutineに入る前に値を追加する。
		wg.Add(1)
		i := i
		go func() {
			limit <- struct{}{}
			defer wg.Done()
			u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
			download(u)
			<-limit
		}()
	}
	wg.Wait()
	//処理時間計測
	fmt.Printf("%v\n", time.Since(before))
}
