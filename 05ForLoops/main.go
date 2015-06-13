package main;

import (
  "fmt"
)

func main() {
  for i := 0; i <= 10; i++ {
    fmt.Printf("Counting %d\n", i)
  }

  fmt.Println("End of loop, look carefully for the next one!")

  // Notice no parens
  for i := 0; i < 10; i++ {
    fmt.Printf("Counting %d\n", i);
  }

  // No such thing as while loops, just modify the for loop:
  keepGoing := true
  for keepGoing {
    fmt.Print("Do you want to keep looping? (y/n) ")
    var _going string
    fmt.Scanf("%s", &_going)
    if _going == "n" {
      keepGoing = false
    } else if _going != "y" {
      fmt.Printf("%s is not a valid option...", _going)
    }
  }
}
