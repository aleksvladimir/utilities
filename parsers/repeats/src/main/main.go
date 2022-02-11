package main

import (
	"os"
)

func main() {
	for {
		data := make([]byte, 8)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			parseAndPrint(data)
		} else {
			break
		}
	}
}

func parseAndPrint(data []byte) {
	// todo
}
