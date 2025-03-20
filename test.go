package main

import ( 
	"fmt"
	"strconv"
) 

func main() {
	test()
	testTwo()
}

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

func test() {
	fmt.Println("test")
}