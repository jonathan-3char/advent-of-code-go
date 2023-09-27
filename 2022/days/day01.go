package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)


func readCalories() []int {
  readFile, err := os.Open("./inputs/input01.txt")
  
  if err != nil {
    fmt.Println(err)
  }

  defer readFile.Close()
  elves := make([]int, 0)
  fileScanner := bufio.NewScanner(readFile)
  calorieCount := 0

  for fileScanner.Scan() {
    line := fileScanner.Text()

    if line == "" {
      elves = append(elves, calorieCount)
      calorieCount = 0
    } else {
      num, err := strconv.Atoi(line)

      if err != nil {
        fmt.Println(err)
      }

      calorieCount += num
    }
  }

  slices.SortFunc(elves, func(a, b int) int {
    return b - a
  })

  return elves
}

func highestCalories() {
  highestCalories := readCalories()[0]

  fmt.Printf(
    "the elf carrying the most calories has a total of %d calories.\n",
    highestCalories)
}

func topThreeHighestCalories() {
  topThreeCalories := readCalories()[:3]
  total := 0

  for _, calories := range topThreeCalories {
    total += calories
  }

  fmt.Printf("the top three elves carrying the most calories have a combined total of %d calories\n.", total) 
}

func Day01(part string) {
  switch part {
  case "1":
    highestCalories()
  case "2":
    topThreeHighestCalories()
  default:
    fmt.Println("invalid part provided.")
  }
}
