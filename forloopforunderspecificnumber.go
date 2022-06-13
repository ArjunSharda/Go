package main

import "fmt"

func main() {
  stopat := 500
  startat := 20
  add := 10
  for startat < stopat {
    startat += add
  if startat < stopat {
    fmt.Printf("Your number is under the number of stopat. \n Your stopat number: %d \n Number currently: %d \n", stopat, startat)
  } else {
    fmt.Printf("Done! At the number of %d", stopat)
  }
    }
  }
