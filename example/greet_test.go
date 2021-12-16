// Package example is a package for demonstrating examples in source code.
package example

import "fmt"

func ExampleHello_name() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}
