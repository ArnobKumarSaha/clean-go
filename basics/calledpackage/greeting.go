package calledpackage

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	tmp := randomFormat()
	message := fmt.Sprintf(tmp, name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("init() called in singular")
}

// TODO : multiple init() function. Which one to call first ?
