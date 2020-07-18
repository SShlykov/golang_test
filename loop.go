package main

import "fmt"

func loop(len int){
	for len < 20 {
		fmt.Printf("This is %d;\n", len)
		len++
	}
}
