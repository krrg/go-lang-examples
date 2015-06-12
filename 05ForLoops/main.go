package main;

import (
  "fmt"
)

func main() {
  for i := 0; i <= 10; i++ {
    fmt.Printf("Counting %d\n", i)
  }

  fmt.Println("End of loop, look carefully for the next one!")

  for i := 0; i < 10; i++ {
    fmt.Printf("Counting %d\n", i);
  }
}
