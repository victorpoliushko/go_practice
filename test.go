package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
) 

func main() {
	// test()
	// testTwo()
	// goVars()
	// funcForSlices()
	// stringsFunc()
	mapsFunc()
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


// Maps

func mapsFunc() {
	ages := make(map[string]int)
	ages["johny"] = 25
	ages["bill"] = 40
	// overrides bill
	ages["bill"] = 51

	fmt.Println(len(ages))

	tickers := map[string]string{
		"GOOG": "google",
		"HOOD": "robin hood",
	}

	fmt.Println(tickers)


	m := make(map[string]int)
	// insert element
	m["BMW"] = 1

	// get an element
	var elem int
	elem = m["BMW"]
	fmt.Println(elem)

	// delete an element
	delete(m, "BMW")

	// check if a key exists
	elem, ok := m["BMW"]
	fmt.Println(elem, ok)

	// map of maps
	hits := make(map[string]map[string]int)
	fmt.Println(hits)

	// Count instance
	missingAges := []string{}
	if _, ok := ages["tony"]; !ok {
		missingAges = append(missingAges, "tony")
	}
	fmt.Println(missingAges)

	// comma ok, used to distinguish a missing entry from a zero value
	// var seconds int
	// var ok bool
	// if seconds, ok := timezone[tz]; ok {
	// 	return seconds
	// }
	// log.Println("unknown timezone: ", tz)
	// return 0
	
}
