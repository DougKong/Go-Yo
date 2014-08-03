package main

import "fmt"

func main() {
  fmt.Println("Please enter a number:")
  var num int

  fmt.Scanf("%d", &num)
  num = num * 2
  fmt.Println(num)
}
