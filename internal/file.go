package internal

import (
	"log"
	"os"
)

func Open(loc string) string {
	input, err := os.ReadFile(loc)
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	return string(input)
}
