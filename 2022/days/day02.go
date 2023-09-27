package days

import (
	"bufio"
	"fmt"
	"os"
)

const (
  _rock = 1
  _paper = 2
  _scissors = 3
  _win = 6
  _draw = 3
)

var _pointsMap = map[string]int{
  "A": _rock,
  "B": _paper,
  "C": _scissors,
}


func computePoints() {
  readFile, err := os.Open("./inputs/input02.txt")
  defer readFile.Close()

  if err != nil {
    fmt.Println(err)
  }

  fileReader := bufio.NewReader(readFile)
  totalPoints := 0

  movesMap := map[string]string{
    "X": "A",
    "Y": "B",
    "Z": "C",
  }



  for {
    var oppMove, ourMove string
    _, err := fmt.Fscanf(fileReader, "%v %v\n", &oppMove, &ourMove)

    // break when EOF is reached
    if err != nil {
      break
    }
    
    ourMove = movesMap[ourMove]
    
    // determine how many points to give based on match result
    // draw: 3, win: 6, lose: 0
    if  ourMove == oppMove {
      totalPoints += _draw
    } else if (ourMove == "A" && oppMove == "C") ||
    (ourMove == "B" && oppMove == "A") || (ourMove == "C" && oppMove == "B") {
      totalPoints += _win
    }

    // give points for the shape selected
    // rock: 1, paper: 2, scissors: 3
    totalPoints += _pointsMap[ourMove]
  }

  fmt.Printf("using the strategy guide will give a total of %d points.\n", 
  totalPoints)
}

func correctComputePoints() {
  readFile, err := os.Open("./inputs/input02.txt") 
  defer readFile.Close()

  if err != nil {
    fmt.Println(err)
  }

  fileReader := bufio.NewReader(readFile)
  totalPoints := 0

  for {
    var oppMove, wantResult string
    _, err := fmt.Fscanf(fileReader, "%v %v\n", &oppMove, &wantResult)

    if err != nil {
      break
    }

    totalPoints += computeMatch(oppMove, wantResult)

  }

  fmt.Printf("total points for part2 are %d.\n", totalPoints)
}

func computeMatch(oppMove, result string) int {
  winMap := map[string]int{
    "A": _paper,
    "B": _scissors,
    "C": _rock,
  }

  loseMap := map[string]int{
    "A": _scissors,
    "B": _rock,
    "C": _paper,
  }


  switch result {
  case "X":
    return loseMap[oppMove]
  case "Y":
    return _pointsMap[oppMove] + _draw
  case "Z":
    return winMap[oppMove] + _win
  default:
    fmt.Println("invalid oppMove provided.")
    return -1
  }

}


func Day02(part string) {
  switch part {
  case "1":
    computePoints() 
  case "2":
    correctComputePoints()
  default:
    fmt.Println("invalid part provided.")
  }
}


