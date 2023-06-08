package main

import (
	"log"
	"os"
)

func main() {
	f1, err := os.Open("no-file.txt")
	if err != nil {
		log.Println(err)
	}

	defer f1.Close()

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}

	defer f2.Close()
}
