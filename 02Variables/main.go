package main

import "fmt"


// Variables can be declared outside a function
var x int = 123
var y = 234   // Type inference for the win

func main() {
  fmt.Println("Here is the value of x: ", x)
  fmt.Println("Here is the value of y: ", y)
}
