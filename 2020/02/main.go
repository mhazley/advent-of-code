package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    input := readData(false)
    // log.Printf("Input Data: %v\n", input)
    // count := analyseStringsOne(input)
    count := analyseStringsTwo(input)
    log.Printf("Num passwords passed: %d", count)
}

func analyseStringsOne(data []string) int {
    var count = 0
    for i, listString := range data {

        // Split and find the char we are searching for
        splitString := strings.Split(listString, " ")
        searchChar := string(splitString[1][0])

        // Check the search char is even in there before we test
        if strings.ContainsAny(splitString[2], searchChar) {

            // Get the lower and upper limits
            limitString := strings.Split(splitString[0], "-")
            lowerLimit, _ := strconv.Atoi(limitString[0])
            upperLimit, _ := strconv.Atoi(limitString[1])

            countInString := strings.Count(splitString[2], searchChar)

            if (lowerLimit <= countInString) && (upperLimit >= countInString) {
                log.Printf("%d: Pass!", i)
                count++
            } else {
                log.Printf("%d: Fail!", i)
            }
        } else {
            log.Printf("%d: Fail! Not in string.", i)
        }
    }
    return count
}

func analyseStringsTwo(data []string) int {
    var count = 0
    for i, listString := range data {

        // Split and find the char we are searching for
        splitString := strings.Split(listString, " ")
        searchChar := splitString[1][0]

        // Check the search char is even in there before we test
        if strings.ContainsAny(splitString[2], string(searchChar)) {

            // Get the indexes
            limitString := strings.Split(splitString[0], "-")
            indexOne, _ := strconv.Atoi(limitString[0])
            indexTwo, _ := strconv.Atoi(limitString[1])

            // Account for index zero being 1
            indexOne--
            indexTwo--

            if ( splitString[2][indexOne] == searchChar ) && ( splitString[2][indexTwo] == searchChar ) {
                log.Printf("%d: Fail! Both are char", i)
            } else if ( splitString[2][indexOne] == searchChar ) || ( splitString[2][indexTwo] == searchChar ) {
                log.Printf("%d: Pass!", i)
                count++
            } else {
                log.Printf("%d: Fail! Neither are char", i)
            }
        } else {
            log.Printf("%d: Fail! Not in string.", i)
        }
    }
    return count
}

func readData(test bool) []string {
    if test {
        test := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
        return test
    } else {
        file, err := os.Open("./input.txt")
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
}