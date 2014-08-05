package main

import "fmt"

func f(from string) {
  for i := 0; i < 10; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  go f("go yo!")
  f("direct")

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
