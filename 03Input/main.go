package main;

import (
  "fmt"
)

func main() {
  fmt.Print("Enter integer: ")
  var x int
  fmt.Scanf("%d", &x)

  fmt.Print("Enter double: ")
  var y float64
  fmt.Scanf("%f", &y)

  fmt.Printf("Here is your integer: %d\n", x)
  fmt.Printf("Here is your double: %f\n", y)
}
