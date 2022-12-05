package main

import (
	"log"
)

func doStuff() string {
	return "I do stuff!"
}

func doStuff1() string {
	return "I do stuff!"
}

func doStuff2() string {
	return "I do stuff!"
}

func main() {
	log.Printf("starting service sonar ...")
	log.Printf(doStuff())
	log.Printf(doStuff1())
	log.Printf(doStuff2())
}
