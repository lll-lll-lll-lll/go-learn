package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sw sync.WaitGroup
	var m sync.Map
	sw.Add(1)
	sw.Add(1)
	m.Store(1, 1)
	go func() {
		defer sw.Done()
		fmt.Println("go func 1")
		if val, ok := m.Load(1); ok {
			fmt.Println("success load", val)
		}
		if isSwapped := m.CompareAndSwap(1, 1, 2); !isSwapped {
			fmt.Println("feaild swap")
		}

	}()
	go func() {
		defer sw.Done()
		fmt.Println("go func 2")
		time.Sleep(time.Second * 3)
		if val, ok := m.Load(1); ok {
			fmt.Println("success load 2", val)
		}
	}()

	sw.Wait()
	fmt.Println("終了")
}

func downLoadURL(u string) error {

}
