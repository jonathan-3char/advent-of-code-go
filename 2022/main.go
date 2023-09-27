package main

import (
	"fmt"
	"log"

	"github.com/jonathan-3char/advent-of-code/2022/days"
)

func main() {
  fmt.Print("input day: ")

  var day, part string
  _, err := fmt.Scan(&day)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Print("input part: ")


  _, err = fmt.Scan(&part)

  if err != nil {
    log.Fatal(err)
  }
  
  switch day {
  case "1":
    days.Day01(part)
  case "2":
    days.Day02(part)
  case "3":
    days.Day03(part)
  default:
    fmt.Println("invalid day provided.")
  }

}
