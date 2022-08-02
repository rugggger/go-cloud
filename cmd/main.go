package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rugggger/go-cloud/src/pkg/fakerequest"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fakeRequest := fakerequest.CreateFakeRequest()

	for {
		res, err := fakeRequest()
		if err != nil {
			fmt.Printf("Request failed with... %s\n", err)
		} else {
			fmt.Printf("Request... %s\n", res)
		}
		time.Sleep(1 * time.Second)
	}

}
