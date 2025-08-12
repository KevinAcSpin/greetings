package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hello, %s! Welcome aboard \U0001F60E",
		"Greetings, %s! We're glad to have you here \U0001F60A",
		"Hi, %s! It's great to see you \U0001F60D",
	}
	return formats[rand.Intn(len(formats))]
}
