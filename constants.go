package main

import "fmt"

const Pi = 3.14

func main(){
  const World = "Awesome"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi ,"Day")
  const Lie = false

  fmt.Println("I suck?", Lie)
}
