package main

import "fmt"

const (
	french = "French"
	spanish = "Spanish"

  englishHelloPrefix = "Hello "
  frenchHelloPrefix = "Bonjour "
  spanishHelloPrefix = "Hola "
)

func Hello(name string, language string) string {
		if name == "" {
			name = "world"
		}
		return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
		switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
		default:
			prefix = englishHelloPrefix
		}
		return
}

func main(){
	fmt.Println(Hello("world", ""))
}
