package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {
  var file, err = ioutil.ReadFile("blah.text")

  fmt.Println(err)

  myString := string(file)
  myStringArray := strings.Fields(myString)
  fmt.Println(myStringArray)
  var slice = strings.Join(Insert(myStringArray, 3, "5555555"), " ")
  fmt.Println(slice)

  ioutil.WriteFile("output.text", []byte(slice), 444)
}

func Insert(array []string, index int, value string) []string {
  // Grow the slice by one element.
  // fmt.Println("yo")
  // var sliceLen int
  // sliceLen = len(array) + 1
  // fmt.Println("Slice length: ", sliceLen)
  // var slice []string
  var slice = make([]string, len(array)+1)
  fmt.Println(array)
  copy(slice, array)

  fmt.Println(slice)

  slice = slice[0:len(slice)]
  // Use copy to move the upper part of the slice out of the way and open a hole.
  fmt.Println("hey")
  copy(slice[index+1:], slice[index:])
  // Store the new value.
  slice[index] = value
  // Return the result.
  return slice
}
