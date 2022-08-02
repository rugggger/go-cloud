package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"

	"github.com/rugggger/go-cloud/src/pkg/fakerequest"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//fakeRequest := fakerequest.CreateFakeRequest()

	go server()

	for {

	}
	// for {
	// 	res, err := fakeRequest()
	// 	if err != nil {
	// 		fmt.Printf("Request failed with... %s\n", err)
	// 	} else {
	// 		fmt.Printf("Request... %s\n", res)
	// 	}
	// 	time.Sleep(500 * time.Millisecond)
	// }

}

func server() {
	r := make(chan string)
	resChannel := make(chan fakerequest.FakeResponse)
	go fakerequest.RunRequestServer(r, resChannel)
	reqNumber := 1
	for {
		fmt.Printf("Send request %d\n", reqNumber)
		select {
		case resp := <-resChannel:
			fmt.Println("got response")
			res, err := resp.Res, resp.Err
			if err != nil {
				color.Red("Error is %s\n", err)
			} else {
				color.Green("Response is %s\n", res)
			}
		default:

		}
		//r <- fmt.Sprintf("%d", reqNumber)
		fmt.Printf("after sleep")

		// resp := <-resChannel
		// res, err := resp.Res, resp.Err
		// if err != nil {
		// 	color.Red("Error is %s\n", err)
		// } else {
		// 	color.Green("Response is %s\n", res)
		// }
		time.Sleep(300 * time.Millisecond)

		reqNumber += 1
	}
}
