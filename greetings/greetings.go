package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"rsc.io/quote"
)

func Hello(userName string) (string, error) {
	if userName == "" {
		return "", errors.New("empty userName")
	}
	
	message := fmt.Sprintf(randomFormatHello(), userName);
	return message, nil
}

func Goodbye(userName string) (string, error) {
	if userName == "" {
		return "", errors.New("empty userName")
	}

	var message string
	message = fmt.Sprintf(randomFormatGoodbye(), userName);
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormatHello returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormatHello() string {
	// A slice of message formats.
	formats := []string{
			"Hi, %v. Welcome!",
			"Great to see you, %v!",
			"Hail, %v! Well met!",
			"To %v: " + quote.Go(),
			"To %v: " + quote.Hello(),
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// randomFormatHello returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormatGoodbye() string {
	// A slice of message formats.
	formats := []string{
			"Peace! %v.",
			"See ya, %v! Until we meet again!",
			"Goodbye, %v. Thank you for your time!",
			"To %v: " + quote.Glass(),
			"To %v: " + quote.Opt(),
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}