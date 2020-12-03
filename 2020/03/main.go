package main

import (
    "bufio"
    "log"
    "os"
)

func main() {
    input := readData(false)
    result1 := navigate(input, 3, 1)
    result2 := navigate(input, 1, 1) *
        navigate(input, 3, 1) *
        navigate(input, 5, 1) *
        navigate(input, 7, 1) *
        navigate(input, 1, 2)

    log.Println(result1)
    log.Println(result2)
}

func navigate(data []string, rightAdvance int, downAdvance int) int {
    var col = 0
    var tree = uint8('#')
    var count = 0
    for row := downAdvance; row < len(data); row=row+downAdvance {
        col += rightAdvance
        if data[row][col%len(data[row])] == tree { count++ }
    }
    return count
}

func readData(test bool) []string {
    var fileName string
    if test {
        fileName = "test_input.txt"
    } else {
        fileName = "input.txt"
    }

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    var input []string

    for scanner.Scan() {
        text := scanner.Text()
        input = append(input, text)
    }

    return input
}