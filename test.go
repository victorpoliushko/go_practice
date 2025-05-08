package main

import (
	"fmt"
	"sync"

	// "log"
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/key"
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
	// a := 1
	// fmt.Println("Init: ", a)

	// basicVar(a)
	// fmt.Println("basic: ", a)

	// zeroVar(&a)
	// fmt.Println("zeroVar: ", a)

	// fmt.Println("pointer:", &a)


	// Replace string for pointer
	// p := "English, motherfubber, do you speak it? shiz"
	// replaceStrings(&p)
	// fmt.Println(p)


	// Pointers: update original value
	// new()

	// Pointers: What is the value of *y after the code on the left executes?
	// pointerTest()

	// Pointers: What is the value of x after the entire code block on the left executes?
	// pointerTestTwo()

	// Concurrency
	// concurrencyTestEmail("Hello there Kaladin!")
	// concurrencyTestEmail("Hi there Shallan!")
	// concurrencyTestEmail("Hey there Dalinar!")

	// Channels
	pingPongTest(3)
	pingPongTest(4)
	pingPongTest(2)
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


// strings.ReplaceAll with pointer
func replaceStrings(message *string) {
	badWords := []string{"fub", "shiz"}

	for _, word := range badWords {
		replacement := ""
		for i := 0; i < len(word); i++ {
			replacement += "*"
		}

		*message = strings.ReplaceAll(*message, word, replacement)
	}
}

// Pointers: update original value
func increment(i *int) {
	*i++;
	fmt.Println("Increment func: ", *i)
}

func new() {
	i := 5
	increment(&i)
	fmt.Println("New func:", i)
}

// Pointers: 
// What is the value of *y after the code on the left executes?
func pointerTest() {
	var x int = 50
	var y *int = &x
	*y = 100
	fmt.Println("*y =", *y)
}

// Pointers: What is the value of x after the entire code block on the left executes?
func pointerTestTwo() {
	var x int = 50
	var y *int = &x
	*y = 100
	fmt.Println("x =", x)
}

// Pointers: before using a pointer it's value should be checked to avoid panic
func emptyPointer(a *int) {
	if a == nil {
		return
	}

	// rest of the function
}

// Pointers: pointer reciever
func (e *email) sendMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message string
	fromAddress string
	toAddress string
}

// Pointers: performance vise just variables are created faster taht pointers because they go into stack while pointer go into heap


// Pointers: update struct values
type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance (c *customer, tx transaction) error {
	switch tx.transactionType {
	case transactionWithdrawal:
		if c.balance < tx.amount {
			return errors.New("insufficient funds")
		}
		c.balance -= tx.amount
	case transactionDeposit:
		c.balance += tx.amount
	default:
		return errors.New("unknown transaction type")
	}
	return nil
}

// Packages: naming convension
// repo: github.com/mailio/rand
// package name: package rand


// Packages:
// Each repo is a singe module
// each repo contains one or more packages
// each repo consists of one or more Go sources files in a singe directory
// The path to a package's directory determines its import path and where it can be downloaded from

// Concurrency
func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func concurrencyTestEmail(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("=======================")
}

// Channels
type channelsEmail struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]channelsEmail) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

func sendIsOld(isOldChan chan <- bool, emails [3]channelsEmail) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}


// Channels 2

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

// Channels: Buffers
// Complete the addEmailsToQueue function. It should create a buffered channel with a buffer 
// large enough to store all of the emails it's given. It should then write the emails to the 
// channel in order, and finally return the channel.
func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, email := range emails {
		ch <- email
	}
	return ch
}

// Channels: Close channel
func countReports(numSentCh chan int) int {
	total := 0
	for {
		v, ok := <- numSentCh
		if !ok {
			return total
		}
		total += v
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

// simultaneous assignment
// x, y = y, x + y
// same as
// tempX := x
// tempY := y
// x = tempY
// y = tempX + tempY

// Channels: range

func concurrentFib(n int) []int {
	ch := make(chan int)
	go fibonacci(n, ch)
	newSlice := []int{}
	for item := range ch {
		newSlice = append(newSlice, item)
	}
	return newSlice
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

// Channels: select
// func logMessages(chEmails, chSms chan string) {
// 	for {
// 		select {
// 		case email, ok := <-chEmails:
// 			if !ok {
// 				return
// 			}
// 			logEmail(email)
// 		case sms, ok := <-chSms:
// 			if !ok {
// 				return
// 			}
// 			logSms(sms)
// 		}
// 	}
// }

// Channels: select default
// func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
// 	for {
// 		select {
// 		case <-snapshotTicker:
// 			takeSnapShot(logChan)
// 		case <-saveAfter:
// 			saveSnapshot(logChan)
// 			return
// 		default:
// 			waitForData(logChan)
// 			time.Sleep(500 * time.Microsecond)
// 		}
// 	}
// }

// Channels
func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	go ponger(pings, pongs)
	go pinger(pings, numPings)
	func() {
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
	}()
}

func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i:= 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %b\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func pingPongTest(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===Game Over===")
}

// Mutex
type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	// allow readers to read
	// sc.mu.RLock()
	// defer sc.mu.RUnlock()
	return sc.slowVal(key)
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

