package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nombre vacio")
	}
	message := fmt.Sprintf(randonFormat(), name)
	return message, nil
}

func randonFormat() string {
	formats := []string{"Hola, %v. Bienvenido",
		"Que bueno verte, %v.",
		"Saludo, %v. Encantado de conocerte"}
	return formats[rand.Intn(len(formats))]
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
