package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sync"
)

// //////////////////////////
// ///// INTERFACE //////////
// //////////////////////////
type Animal interface {
	Speak()
} // any struct with speak can implement it

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() {
	fmt.Print("vow vow!")
}

func (c Cat) Speak() {
	fmt.Print("meow!")
}

func testSound(a Animal) {
	a.Speak()
}

// //////////////////////////
// ///// GOROUTINE //////////
// //////////////////////////
func worker(wg *sync.WaitGroup) {
	runtime.LockOSThread()
	defer wg.Done()

	for i := 0; i <= 10000000; i++ {
		fmt.Printf("working: %v", i)
		println()
	}
}

// /////////////////////////////////
// //////// CHANNELS ///////////////
// /////////////////////////////////
func channels() chan string {
	ctx := context.Background()

	fmt.Println(ctx.Deadline())

	ch := make(chan string)

	go func() {
		for i := 0; i <= 3; i++ {
			ch <- "hello"
		}
		close(ch)
	}()

	return ch
}

type Pod struct {
	Name string `json:"name"`
}

// JSON PARSING
func parsingJson() {
	data := []byte(`{"name":"nginx"}`)

	var p Pod

	json.Unmarshal(data, &p)

	fmt.Println(p.Name)

	_, err := os.ReadFile("abc.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	d := Dog{}
	c := Cat{}
	testSound(d)
	testSound(c)

	// var wg sync.WaitGroup

	// wg.Add(1)
	// // go worker(&wg)
	// wg.Wait()

	ch := channels()

	for v := range ch {
		fmt.Println()
		fmt.Println(v)
	}

	parsingJson()
}
