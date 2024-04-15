package server

import "fmt"

const prefix = "Hello"

func print(message string) string {
	if message == "" {
		message = "World"
	}

	return fmt.Sprintf("%s %s", prefix, message)
}
