
package greeting

import (
	"errors"
	"fmt"
)

// Greeting returns a greeting for the named person/organisation.
func Greeting(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("The name is Empty ")
	}
	msg := fmt.Sprintf("Hello world from %s.", name)
	return msg, nil
}

