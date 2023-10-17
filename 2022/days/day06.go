package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day06(part string) {
    switch part {
        case "1": 
        result := findStartOfPacketMarker(4)
        fmt.Println("the amount of characters to proccess before start-of-packer marker are", result)
    case "2":
        result := findStartOfPacketMarker(14)
        fmt.Println("the amount of characters to proccess before start-of-packer marker are", result)
    } 
}

func findStartOfPacketMarker(distinct int) int {
    file, err := os.Open("./inputs/input06.txt")

    if err != nil {
        log.Fatal(err)
    }

    fileReader := bufio.NewReader(file)
    signal, _, err := fileReader.ReadLine()

    if err != nil {
        log.Fatal(err)
    }

    size := len(signal)
    for i := range signal {
        marker := make(map[byte]struct{}) 

        for j := i; j < i + distinct && j < size; j++ {
            marker[signal[j]] = struct{}{}
        }

        if len(marker) == distinct {
            return i + distinct
        }
    }

    return -1
}
