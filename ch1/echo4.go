package ch1

import "fmt"

func Echo4(args []string) {

	for i, arg := range args[1:] {
		fmt.Println(i+1, arg)
	}

}
