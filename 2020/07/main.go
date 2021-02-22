package main

import (
    "bufio"
    "log"
    "os"
)

func main() {
    input := readData(true)
    log.Println(input)
}

//func countOuterBags(data []string) (sum int) {
//
//}

func readData(test bool) []string {
    var file *os.File
    var err error
    if test {
        file, err = os.Open("./test_input.txt")
        if err != nil {
            log.Fatal(err)
        }
    } else {
        file, err = os.Open("./input.txt")
        if err != nil {
            log.Fatal(err)
        }
    }

    scanner := bufio.NewScanner(file)

    var input []string

    for scanner.Scan() {
        text := scanner.Text()
        input = append(input, text)
    }

    return input
}
