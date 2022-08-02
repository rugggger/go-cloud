package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ERROR_RATE = 0.2

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Printf("Request... %d\n", randInt(10, 100))
		time.Sleep(1 * time.Second)
	}

}

func randInt(min int, max int) (int, error) {

	return min + rand.Intn(max-min), nil
}
