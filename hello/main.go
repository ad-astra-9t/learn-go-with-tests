package main

import "fmt"

const helloPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%v %v", helloPrefix, name)
}

func main() {
	fmt.Println(Hello("Tim"))
}
