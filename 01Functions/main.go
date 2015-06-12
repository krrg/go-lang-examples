package main

import "fmt"

func main() {
  fmt.Println(bigCompute())
  fmt.Println(biggerCompute())
  fmt.Println(giantCompute(747))
  fmt.Println(giantessCompute(747))  // Will not print 747!
}

// Notice that a return type is necessary!
func bigCompute() string {
  return "Hello World"
}

func biggerCompute() (int, int) {
  return 1, 2  //You can even return multiple values
}

// Explicitly named return values;  at the end of the function,
//  whatever the value of y is will be returned.
func giantCompute(x int) (y int) {
  y = x
  return // Gathers up the value of y and returns it.
}

// But this can be overridden!
func giantessCompute(x int) (y int) {
  y = x
  return 800  // Does not return `y` in this case
}
