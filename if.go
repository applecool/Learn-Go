package main

import (
    "fmt"
    "math"
)

func squarert(x float64) string {

  if x<0 {
    return squarert(-x) + "i"
  }

    return fmt.Sprint(math.Sqrt(x))
}


func main(){
   fmt.Println(squarert(5), squarert(-9))
}
