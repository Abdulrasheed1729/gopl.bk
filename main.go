package main

import (
	"fmt"
	"os"

	ch1 "github.com/Abdulrasheed1729/gopl.bk/ch1"
)

func main() {
	ch1.Echo1()
	ch1.Echo2(os.Args)
	ch1.Echo4(os.Args)
	fmt.Println(ch1.Echo3(os.Args))
}
