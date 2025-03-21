package main

import ( 
	"fmt"
	"strconv"
) 

func main() {
	test()
	testTwo()
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

func callFunc(func (int) int) {
	
}


// Slice

var newSlice []int

// possible to create this way only in functions

func funcForSlices() {
	newSlice := []string{"First value", "Second value", "Third value"}
	makeNewSlice := make([]int, 5)
	makeNewSliceWithCapacity := make([]int, 5, 3)

	arr := [5]int{1, 2, 3, 4, 5}
	arrSlice := arr[1:4]

	newSlice = append(newSlice, )

	matrixSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
}


// Loop (slice)

func loopFunction() {
	messages := []string{"first message", "Second", "Third message"}
	for index, value := range messages {
		fmt.Println(index)
		printParam()
	}
}



