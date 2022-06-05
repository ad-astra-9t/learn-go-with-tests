package main

import "fmt"

var helloPrefix = map[string]string{
	"en": "Hello, ",
	"fr": "Bonjour, ",
	"zh": "你好，",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%v%v", helloPrefix[language], name)
}

func main() {
	fmt.Println(Hello("Tim", "en"))
}
