// publish2.go

package main

import (
  "fmt"
  "time"
)

func main() {
  var ch <-chan struct{}
  ch = Publish("hello", 5)

  fmt.Println(<-ch)
}

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
  ch := make(chan struct{})
  go func() {
    time.Sleep(delay)
    fmt.Println("Breaking News:", text)
    close(ch)
  }()

  return ch
}
