package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countFullyContain() {
  file, err := os.Open("./inputs/input04.txt")
  defer file.Close()

  if err != nil {
    log.Fatal(err)
  }
  
  fileReader := bufio.NewReader(file)
  totalOverlap := 0

  for {
    var x1, x2, y1, y2 int

    _, err := fmt.Fscanf(fileReader, "%d-%d,%d-%d\n", &x1, &x2, &y1, &y2)

    if err != nil {
      break
    }


    if ( x1 <= y1 && x2 >= y2 ) || ( y1 <= x1 && y2 >= x2) {
      totalOverlap++
    }
  }

  fmt.Printf("the number of assignment pairs where one range fully contains the other is %d.\n", totalOverlap)

}

func countOverlap() {
  file, err := os.Open("./inputs/input04.txt")
  defer file.Close()

  if err != nil {
    log.Fatal(err)
  }

  fileReader := bufio.NewReader(file)
  totalOverlap := 0

  for {
    var x1, x2, y1, y2 int

    _, err := fmt.Fscanf(fileReader, "%d-%d,%d-%d\n", &x1, &x2, &y1, &y2)

    if err != nil {
      break
    }

    if x1 <= y2 && y1 <= x2 {
      totalOverlap++
    }

  }

  fmt.Printf("the number of assignment pairs where ranges overlaps %d.\n", totalOverlap)
}

func Day04(part string) { 
  switch part {
  case "1":
    countFullyContain()
  case "2":
    countOverlap()
  default:
    fmt.Println("invalid part provided.")
  }
}
