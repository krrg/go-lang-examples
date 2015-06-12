package main;

import (
  "fmt"
)

func even(n int) bool {
  return n % 2 == 0
}

func main() {
  fmt.Print("Enter an integer: ")
  var x int
  fmt.Scanf("%dl", &x)

  var evenString string
  if even(x) {
    evenString = "even"
  } else {
    evenString = "odd"
  }

  fmt.Printf("The number %d is %s\n", x, evenString)
}
