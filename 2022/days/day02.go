package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func computePoints() {
  readFile, err := os.Open("./inputs/input02.txt")
  defer readFile.Close()

  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  totalPoints := 0

  movesMap := map[string]string{
    "X": "A",
    "Y": "B",
    "Z": "C",
  }

  pointsMap := map[string]int{
    "A": 1,
    "B": 2,
    "C": 3,
  }


  for fileScanner.Scan() {
    line := strings.Split(fileScanner.Text(), " ")
   
    if len(line) == 1 {
      break
    }

    // determine how many points to give based on match result
    // tie: 3, win: 6, lose: 0
    if ourMove := movesMap[line[1]]; ourMove == line[0] {
      totalPoints += 3
    } else if (ourMove == "A" && line[0] == "C") ||
    (ourMove == "B" && line[0] == "A") || (ourMove == "C" && line[0] == "B") {
      totalPoints += 6
    }

    // give points for the shape selected
    // rock: 1, paper: 2, scissors: 3
    totalPoints += pointsMap[movesMap[line[1]]]
  }

  fmt.Printf("using the strategy guide will give a total of %d points.\n", 
  totalPoints)
}

func Day02(part string) {
  switch part {
  case "1":
    computePoints() 
  case "2":
    fmt.Println("not done yet...")
  default:
    fmt.Println("invalid part provided.")
  }
}


