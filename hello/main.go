package main

import "fmt"

const helloPrefix = "Hello,"

func Hello(name string) string {
	return fmt.Sprintf("%s %s", helloPrefix, name)
}

func main() {
	fmt.Println(Hello("Tim"))
}
