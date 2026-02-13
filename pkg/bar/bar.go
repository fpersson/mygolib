package bar

import (
	"fmt"
	"log/slog"
)

func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %s.", name)
	return msg
}

func LogSomething(something string) {
	slog.Info(something)
}
