package main

import (
	"log"
)

func main() {
	_, err := handlers.newServer(54321)
	if err != nil {
		log.Fatal(err)
	}

}
