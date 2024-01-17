package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	fmt.Println(buf)
	r := io.Reader(os.Stdin)
	n, err := r.Read(buf)
	fmt.Println(buf)

	if n > 0 {
		fmt.Println(n)
		//読み込んだデータ分を出力
		fmt.Println(buf[:n])
	} else {
		fmt.Println("No data read")
	}
	if err != nil {
		fmt.Println("err faild to read buf")
		return
	}
}
