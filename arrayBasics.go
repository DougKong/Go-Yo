// arrayBasics.go

package main

import (
  "fmt"
)

func main() {
  var blah [5]float32
  blah[0] = 6

  blah[3] = 6.44

  blah[1] = 6.76543

  fmt.Println(blah)

  // for i := 0; i < len(blah); i++ {
  //   var temp int
  //   blah[len(blah)-i]
  // }

  // var sum float32
  // sum = 0

  sum := float32(0)

  for _, value := range blah {
    sum += value
  }
  fmt.Println("total: ", sum)
  var average float32
  var length float32
  length = float32(len(blah))
  average = sum / length
  fmt.Println("Average: ", average)

}
