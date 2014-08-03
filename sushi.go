// sushi.go

package main

import (
  "fmt"
)

type Sushi struct {
  name string
}

func main() {
  var ch <-chan Sushi
  ch = Producer()
  for s := range ch {
    fmt.Println("Consumed", s.name)
  }
}

func Producer() <-chan Sushi {
  ch := make(chan Sushi)

  go func() {
    ch <- Sushi{"hello"}
    ch <- Sushi{"goodbye"}
    close(ch)
  }()
  return ch
}
