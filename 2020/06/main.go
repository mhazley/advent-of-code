package main

import (
    "bufio"
    "log"
    "os"
    "strings"
)

func main() {
    input := readData(false)
    log.Println(check(input))
}

func check(data [][]string) (sum int) {
    sum = 0
    for _, entry := range data {
        for _, char := range entry[0] {
            charPresent := true
            for i := 1; i < len(entry); i++ {
                if strings.ContainsRune(entry[i], char) {
                    continue
                } else {
                    charPresent = false
                }
            }
            if charPresent {sum++}
        }
    }
    return
}

func readData(test bool) [][]string {
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

    var input [][]string
    var entry []string

    for scanner.Scan() {
        text := scanner.Text()
        if len(text) > 0 {
            entry = append(entry, text)
        } else {
            input = append(input, entry)
            entry = []string{}
        }
    }
    // Last input needs to append!
    input = append(input, entry)

    return input
}
