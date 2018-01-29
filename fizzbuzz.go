package main

import (
  "fmt"
)

func main(){
  for i := 1; i < 100; i++ {
    if i%15 == 0 {
      fmt.Printf("fizzbuzz\n")
    } else if i%5 == 0 {
      fmt.Printf("buzz\n")
    } else if i%3 ==0{
      fmt.Printf("fizz\n")
    } else {
      fmt.Printf("%d\n", i)
    }
  }
}
