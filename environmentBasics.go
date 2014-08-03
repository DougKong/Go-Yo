// environmentBasics.go

package main

import (
  "fmt"
  "os"
)

func main() {

  os.Setenv("Hey", "You")
  fmt.Println(os.Getenv("Hey"))
  fmt.Print("**************")

  var environmentVariables = os.Environ()
  for _, value := range environmentVariables {
    fmt.Println(value)
  }

}
