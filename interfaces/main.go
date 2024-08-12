package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// can ommit instance if not uses
func (englishBot) getGreeting(text string, n int) (string, error) {
	return fmt.Sprintf("Hello %s, %i", text, n), nil
}

func (sp spanishBot) getGreeting(text string, n int) (string, error) {
	return fmt.Sprintf("Hola %s, %i", text, n), nil
}
