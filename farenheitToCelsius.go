// farenheitToCelsius.go

package main

import (
  "fmt"
)

func main() {
  var farenheitNumber float64
  fmt.Println("Enter in number in Farenheit")
  fmt.Scanf("%f", &farenheitNumber)

  var celsius float64
  celsius = (farenheitNumber - 32) * 5 / 9

  // don't do the following.  makes celsius = 0
  // this is stupid.  Int/Int == Int.
  // celsius = (farenheitNumber - 32) * (5 / 9)

  fmt.Println(celsius)
  fmt.Printf("Celsius %f", celsius)
}
