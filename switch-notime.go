package main

import(
  "fmt"
  "time"
)

func main(){
  time_now:= time.Now()

  switch{
  case time_now.Hour() <12: fmt.Println("Good Morning")
  case time_now.Hour() < 17: fmt.Println("Good Afternoon")
  default: fmt.Println("Good evening")
  }
}
