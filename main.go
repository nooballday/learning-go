package main

import (
	"fmt"
	"sort"
)

func main() {
	comment()
	variables()
	dataType()
	structInterface()
	functions()
	objects()
	arrayAndSlice()
	makeFunction()
	newFunction()
}

func arrayAndSlice() {
	var arrLength int = 3
	var arr1 [3]string = [3]string{"a", "b", "c"}
	var sliceArr = make([]string, arrLength)
	sliceArr[0] = "Hello"
	sliceArr[1] = ""
	elem := sliceArr[0]
	fmt.Println(arr1, sliceArr, elem)

	numbers := []int{100, 123, 90, 22, 1, 25, 8802, 34289}
	sort.Ints(numbers)
	fmt.Println(numbers)

	fmt.Println(numbers[0:5]) // silce from index 0 (included) to index 5
}

func makeFunction() {
	var arrOfString []string = make([]string, 2, 100)
	arrOfString[0] = "1"
	arrOfString[1] = "2"
	// arrOfString[2] = "3" //error because the returned slice length is 2
	fmt.Println(arrOfString)
}

func newFunction() {
	newValue := new(string)
	*newValue = "this is weird"
	fmt.Println(*newValue)
}

func structInterface() {

}

func objects() {
	panic("unimplemented")
}

func functions() {
	panic("unimplemented")
}

func comment() {
	// comment
	/**
	test comment
	*/
}

func variables() {
	var hello string = "wow"
	hello2 := hello
	var (
		hello3 string = hello2
		hello4 int    = 1
	)
	fmt.Print(hello3, hello4)
}

func dataType() {
	var a bool = true
	var b int = 5
	var c float32 = 32.1
	var d string = "Hi!"
	fmt.Print(a, b, c, d)
}
