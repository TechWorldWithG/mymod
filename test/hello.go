package test

import "fmt"

func Hello(greet, name string) string {
	return fmt.Sprintf("%v %v", greet, name) // "greet name" // greet = hello // name=jay
}
