package days

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"

    "golang.org/x/exp/slices"
)

func Day05(part string) {
    switch part {
    case "1":
        file, err := os.Open("./inputs/input05.txt")
        defer file.Close()

        if err != nil {
            log.Fatal(err)
        }

        fileReader := bufio.NewReader(file)

        crateStacks := readFile(fileReader)

        moveOneEach(fileReader, crateStacks)

        case "2": 
        file, err := os.Open("./inputs/input05.txt")
        defer file.Close()

        if err != nil {
            log.Fatal(err)
        }

        fileReader := bufio.NewReader(file)

        crateStacks := readFile(fileReader)

        moveMulti(fileReader, crateStacks)
    default:
        fmt.Println("invalid part provided")
    }
}

func readFile(fileReader *bufio.Reader) [][]rune {

    crates := make([][]rune, 9)

    for {
        crateLine, err := fileReader.ReadString('\n')
        if err != nil || len(crateLine) == 1 {
            break
        }

        for pos, item := range crateLine {
            if unicode.IsLetter(item) {
                // fmt.Printf("%c at position %d", item, (pos ) / 4)
                crates[pos / 4] = append(crates[pos / 4], item)
            }
        }
    }

    for _, stack := range crates {
        slices.Reverse(stack)
    }

    return crates
}

func moveOneEach(fileReader *bufio.Reader, crates [][]rune) {
    for {
        var crateAmount, fromPos, toPos int 
        _, err := fmt.Fscanf(fileReader, "move %d from %d to %d\n",
        &crateAmount, &fromPos, &toPos)

        if err != nil {
            break 
        }
        fromPos--
        toPos--

        for i := 0; i < crateAmount; i++ {
            stack := crates[fromPos]
            crate := stack[len(stack) - 1]
            crates[fromPos] = stack[:len(stack) - 1]
            crates[toPos] = append(crates[toPos], crate)
        }

    }

    fmt.Println("the crates on top of each stack are:")
    for _, crate := range crates {
        fmt.Printf("%c", crate[len(crate) - 1])
    }
    fmt.Println()
}

func moveMulti(fileReader *bufio.Reader, crates [][]rune) {
    for {
        var amount, fromPos, toPos int
        _, err := fmt.Fscanf(fileReader, "move %d from %d to %d\n", &amount, &fromPos, &toPos)
        
        if err != nil {
            break
        }

        fromPos--
        toPos--

        stack := crates[fromPos]
        mvCrates := stack[len(stack) - amount :]
        crates[fromPos] = stack[:len(stack) - amount]
        crates[toPos] = append(crates[toPos], mvCrates...)
    }

    fmt.Println("the crates on top of each stack are:")
    for _, crate := range crates {
        fmt.Printf("%c", crate[len(crate) - 1])
    }
    fmt.Println()
}
