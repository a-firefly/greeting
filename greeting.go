package greeting

import "fmt"

func Greet(name string, words string) string {
	return fmt.Sprintf("Hello, %s. %s!", name, words)
}
