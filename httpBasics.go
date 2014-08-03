// httpBasics.go

package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/", timeHandler)
  http.ListenAndServe(":3000", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, time.Now())
}
