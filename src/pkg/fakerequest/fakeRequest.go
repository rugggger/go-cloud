package fakerequest

import (
	"errors"
	"math/rand"
)

const ERROR_PERCENT_RATE = 20

type FakeRequest struct {
}
type RequestFunction func() (string, error)

func CreateFakeRequest() RequestFunction {

	return func() (string, error) {
		err := raisedError()
		if err != nil {
			return "", err
		}
		return "response", nil
	}

}

func raisedError() error {

	rand := rand.Intn(100)
	if ERROR_PERCENT_RATE > rand {
		return errors.New("Raised Error")
	}
	return nil

}
