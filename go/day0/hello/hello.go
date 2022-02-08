package main

import (
	"fmt"
	"log"
	"tleroy/greetings"
)

func main() {
	//var (
	//	name  string
	//	name2 string
	//)

	// A slice of name
	names := []string{"Theophile", "Bob", "Dylan", "Mélanie"}

	// Getting name from user
	//fmt.Printf("Enter your name:")
	//fmt.Scanf("%v %v",&name,&name2)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Getting greetings and printing it
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
