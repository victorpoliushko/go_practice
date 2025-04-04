package main

import (
	"fmt"
	// "log"
	"strconv"
	"strings"
	"unicode/utf8"
) 



func main() {
	// test()
	// testTwo()
	// goVars()
	// funcForSlices()
	// stringsFunc()
	// mapsFunc()
	// countDistinctWords([]string{"WTS wtb", "inv lfg", "INV lfm"})

	// friendships := map[string][]string{
  //   "Alice":   {"Bob", "Charlie"},
  //   "Bob":     {"Alice", "Charlie", "David"},
  //   "Charlie": {"Alice", "Bob", "David", "Eve"},
  //   "David":   {"Bob", "Charlie"},
  //   "Eve":     {"Charlie"},
	// }
	// findSuggestedFriends("Alice", friendships)

	// 	fmt.Println(logAndDelete(map[string]user{
	// 		"laura": {name: "laura", number: 4355556023, admin: false},
	// 		"dale":  {name: "dale", number: 8015558937, admin: true},
	// 	}, "dale"))


	// Pointers
	a := 1
	fmt.Println("Init: ", a)

	basicVar(a)
	fmt.Println("basic: ", a)

	zeroVar(&a)
	fmt.Println("zeroVar: ", a)

	fmt.Println("pointer:", &a)
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

	fmt.Println("Rune count: ", utf8.RuneCountInString(str))
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
	// map[string]map[int]map[string]int

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

// Fill nested maps
func fillNestedMaps(names []string) map[rune]map[string]int {
	newMap := make(map[rune]map[string]int);

	for _, name := range names {
		runes := []rune(name)
		if newMap[runes[0]] == nil {
			newMap[runes[0]] = make(map[string]int)
		}
		newMap[runes[0]][name]++
	}
	return newMap
}

// Maps + Slices
func countDistinctWords(messages []string) int {
	wordsMap := make(map[string]int)
	for _, message := range messages {
		wordsSlice := strings.Fields(message)
		for _, word := range wordsSlice {
			wordsMap[strings.ToLower(word)]++
		}
	}
	return len(wordsMap)
}

// Search in a map of slices 
func findSuggestedFriends(username string, friendships map[string][]string) []string {
	var directFriendsSet = make(map[string]bool) 
	for _, friend := range friendships[username] {
		directFriendsSet[friend] = true
	}

	var suggestions = make(map[string]bool)

	for _, friend := range friendships[username] {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend != username && !directFriendsSet[friendOfFriend] {
				suggestions[friendOfFriend] = true
			}
		}
	}

	var res []string
	// sug is a key
	for sug := range suggestions {
		res = append(res, sug)
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

// Defer 

// const (
// 	logDeleted = "user deleted"
// 	logNotFound = "user not found"
// 	logAdmin = "user admin deleted"
// )

// type user struct {
// 	name string
// 	number int
// 	admin bool
// }

// func logAndDelete(users map[string]user, name string) (log string) {
// 	defer delete(users, name)

// 	user, ok := users[name]

// 	if !ok {
// 		return logNotFound
// 	}

// 	if user.admin {
// 		return logAdmin
// 	}

// 	return logDeleted
// }


// Pointer
func basicVar(a int) {
	a = 0
}

func zeroVar(aPointer *int) {
	*aPointer = 0
}
