package days

import (
	"bufio"
	"fmt"
	"os"
)

func itemPrioity(item byte) int {
  if item >= 'a' && item <= 'z' {
    return int( item - 'a' + 1 )
  }

  return int( item - 'A' + 27 )
}

func sumCompartments() {
  file, err := os.Open("./inputs/input03.txt")
  defer file.Close()

  if err != nil {
    fmt.Println(err)
  }

  fileReader := bufio.NewReader(file)
  totalPriorities := 0

  for {
    var rucksack string
    _, err := fmt.Fscanf(fileReader, "%s\n", &rucksack)

    if err != nil {
      break
    }

    firstCmpMap := make(map[byte]bool)
    cmpLen := len(rucksack)
    bytesCmp  := []byte(rucksack)

    for i := 0; i < cmpLen / 2; i++ {
      firstCmpMap[bytesCmp[i]] = true 
    }

    for i := cmpLen / 2; i < cmpLen; i++ {
      _, ok := firstCmpMap[bytesCmp[i]]

      if ok {
        totalPriorities += itemPrioity(bytesCmp[i])
        break
      }
    }
  }

  fmt.Printf("the sum of the priorities of those item types is %d\n", totalPriorities)
}

func sumGroupsOfThree() {
  file, err := os.Open("./inputs/input03.txt")
  defer file.Close() 

  if err != nil {
    fmt.Println(err)
  }
  
  totalPriorities := 0
  fileReader := bufio.NewReader(file)

  for {
    var elf1, elf2, elf3 []byte 
    _, err := fmt.Fscanf(fileReader, "%s\n%s\n%s\n", &elf1, &elf2, &elf3)

    if err != nil {
      break
    }

    itemFrequency := make(map[byte]int)

    for _, char := range elf1 {
      itemFrequency[char] = 1
    }

    for _, char := range elf2 {
      _, ok := itemFrequency[char]

      if ok {
        itemFrequency[char] = 2
      }
    }

    for _, char := range elf3 {
      item, ok := itemFrequency[char]

      if ok && item == 2 {
        totalPriorities += itemPrioity(char)
        break
      }
    }
  }

  fmt.Printf("part 2's sum of the priorities of those item types is %d.\n", totalPriorities)

}

func Day03(part string) {
  switch part {
  case "1":
    sumCompartments()
  case "2":
    sumGroupsOfThree()
  default:
    fmt.Println("invalid part provided.") 
  }
}
