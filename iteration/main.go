package main

import "fmt"

func Repeat(s string, count int) string {
	rs := ""
	for i := 0; i < count; i++ {
		rs += s
	}
	return rs
}

func main() {
	fmt.Println("Hello, World!")
}
