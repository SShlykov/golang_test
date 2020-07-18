package main

import (
	"fmt"
	"github.com/SShlykov/golang_test/uc"
	"os"
	"regexp"
	"strconv"
)

func printTest() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)
	errHandler("Can not print message", err)
	fmt.Printf("%d bytes written.\n", n)
}
func cycleTest(){
	cycle(0, 1, printNum)
}
func arrMultTest() {
	arr := []int{1,3,4,5}
	num := mult(arr)
	fmt.Printf("multiply of %v = %d\n", arr, num )
}
func recursionTest() {
	fibN := 7
	fib := fibonacci(fibN)
	fmt.Printf("fibonacci from %d = %d\n", fibN, fib)

	fac := factorial(fibN)
	fmt.Printf("factorial from %d = %d\n", fibN, fac)
}
func factoryTest() {
	inpArr := []int{1,3,54,23,12,5,43,983}
	odd,even := filterFactory(isOdd)(inpArr)
	arrWriter("even", inpArr, even)
	arrWriter("odd", inpArr, odd)

	more,less := filterFactory(compWith5)(inpArr)
	arrWriter("more than 5: ", inpArr, more)
	arrWriter("less or equal than 5: ", inpArr, less)
}
func bubleSortTest()  {
	inpArr := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	bubbleSort(inpArr)
	fmt.Printf("after sort: %v\n", inpArr)
}
func mapOffuncs()  {
	createMapOfFuncs()
	createWeekMap()
}
func regexTest()  {
	str := "param 1 is 7987234.233"
	re, _ := regexp.Compile("[0-9]+.[0-9]+")
	f := func (s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v * 2, 'f', 2, 32)
	}
	str2 := re.ReplaceAllStringFunc(str, f)
	fmt.Println(str2)
}
func importTest()  {
	fmt.Printf("imported uppercase %s\n", uc.UpperCase("dfskjhfds"))
}
func httpTest() {
	addr := "https://jsonplaceholder.typicode.com/posts/1"
	postHandler(addr)
}
func loopTest(){
	len := 10
	loop(len)
}