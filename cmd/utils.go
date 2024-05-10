package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func GetStinData() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\r')

	if err != io.EOF && err != nil {
		log.Fatalf("Error - Utils: %s", err)
	}

	return text
}
