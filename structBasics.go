package main

import (
  "fmt"
)

type Rectangle struct {
  width  int
  height int
}

func (r Rectangle) setWith() {
  r.width = 5
}
func (r Rectangle) getWidth() int {
  return r.width
}

func main() {
  var rect = Rectangle{5, 3333}
  fmt.Println(rect.getWidth())
}
