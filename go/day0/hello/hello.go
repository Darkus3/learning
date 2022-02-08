package main

import (
  "fmt"
  "log"
  "tleroy/greetings"
)

func main() {
  var (
    name string
    name2 string
  )
  // Getting name from user
  fmt.Printf("Enter your name:")
  fmt.Scanf("%v %v",&name,&name2)

  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  // Getting greetings and printing it
  message, err := greetings.Hello(name+name2)
  if err != nil { log.Fatal(err) }

  fmt.Println(message)
}
