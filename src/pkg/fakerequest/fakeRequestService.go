package fakerequest

import (
	"fmt"
	"math/rand"
	"time"
)

type FakeResponse struct {
	Res string
	Err error
}

func RunRequestServer(req chan string, res chan FakeResponse) {
	for {
		r := <-req
		fmt.Println("Received request:", r)
		processingTime := time.Duration((50 + rand.Intn(200)) * int(time.Millisecond))
		time.Sleep(processingTime)
		fmt.Printf("processed for %dms\n", processingTime/time.Millisecond)
		err := raisedError()

		res <- FakeResponse{
			Res: fmt.Sprintf("%s!!!", r),
			Err: err,
		}

	}

}
