package bar

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %s.", name)
	return msg
}
