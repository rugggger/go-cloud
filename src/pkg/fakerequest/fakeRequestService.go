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
		select {
		case r := <-req: //received request
			fmt.Println("Received request:", r)
			processingTime := time.Duration((50 + rand.Intn(200)) * int(time.Millisecond))
			time.Sleep(processingTime)
			fmt.Printf("processed for %dms\n", processingTime/time.Millisecond)
			err := raisedError()
			fmt.Println(err)
			response := FakeResponse{
				Res: fmt.Sprintf("%s!!!", r),
				Err: err,
			}
			select {
			case res <- response:
				fmt.Println("sent")
			default:
			}

		default:
			fmt.Println("Dont block")
			time.Sleep(50 * time.Millisecond)
		}

	}

}
