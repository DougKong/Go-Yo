package main

import (
  "fmt"
)

func main() {
  ch := make(chan string)

  go func() {
    ch <- "Hello!"
    ch <- "ddd!"
    ch <- "aaa!"
    ch <- "sss!"
    ch <- "ffff!"
    close(ch)
  }()
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  v, ok := <-ch
  fmt.Println(v)
  fmt.Println(ok)
}
