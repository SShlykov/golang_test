package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	printTest()
	cycleTest()
	arrMultTest()
	recursionTest()
	factoryTest()
	bubleSortTest()
	mapOffuncs()
	regexTest()
	importTest()

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("All operations took %s\n", delta)
}
