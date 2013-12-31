package main

import (
	"github.com/tsukanov/steaminfo-go/tracker"
	"log"
)

func main() {
	log.Println("Recording app usage...")
	err := tracker.RecordHistory()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("History is recorded!")
}
