package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initiazlize() {
	once.Do(func() {
		fmt.Println("Initializing......")
	})
}

func main() {
	for i := range 5 {
		initiazlize()
		fmt.Printf("app %d is connected to db\n", i)
	}
}