package main

import "fmt"

func main(){

  //var sum int = 0 //general style of initialization.
   sum := 0 //short form

  //for var i int = 0; i<=10;i++ { //cannot use var declaration in the for initializer
    for i:=0; i<=10; i++{
      sum = sum + i
    }

  fmt.Println(sum)

}
