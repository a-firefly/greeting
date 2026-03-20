package greeting

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s. Nice to meet you!", name)
}
