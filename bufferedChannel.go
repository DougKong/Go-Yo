package main

import "fmt"

func main() {
  c := make(chan int, 2)
  c <- 1
  c <- 2
  // c <- 3   // produces a wait error
  fmt.Println(<-c)
  fmt.Println(<-c)
  // fmt.Println(<-c)   // produces a goroutines are asleep error
}
