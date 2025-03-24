package main

import ( 
	"fmt"
	"strconv"
	"strings"
) 

func main() {
	// test()
	// testTwo()
	goVars()
	// funcForSlices()
	// stringsFunc()
}


// Variable creation

func goVars() {
	var a = "apple"
	fmt.Println(a)
	var bread, milk string = "Bread", "Milk"
	fmt.Println(bread, milk)
	var empty int
	fmt.Println(empty)
	water := "Water"
	fmt.Println(water)
}


// Interface

type Book struct {
	Title string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}


// Function

func test() {
	fmt.Println("test")
}

func printParam(s string) {
	fmt.Printf(`Printed: %s`, s)
}

// Slice

var newSlice []int

// possible to create this way only in functions

// func funcForSlices() {
	// newSlice := []string{"First value", "Second value", "Third value"}
// 	makeNewSlice := make([]int, 5)
// 	makeNewSliceWithCapacity := make([]int, 5, 3)

	// newSlice = append(newSlice, "Fourth")
	// fmt.Println(newSlice)

// 	arr := [5]int{1, 2, 3, 4, 5}
// 	arrSlice := arr[1:4]

// 	matrixSlice := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 	}
// }


// Strings

func stringsFunc() {
	var str = "AaAaA"
	fmt.Println(strings.ToLower(str))

	if strings.Contains(str, "B") {
		fmt.Println("B is here!")
	}
}

