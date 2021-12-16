package example

import "fmt"

type Demo struct{}

func (d Demo) Hello() {}

// Hello prints out hello to the person provided
func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}
