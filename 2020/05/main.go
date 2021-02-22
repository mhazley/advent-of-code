package main

import (
    "bufio"
    "log"
    "os"
)

func main() {
    input := readData(false)
    log.Println(analyseTickets(input))
}

func analyseTickets(data []string) (max int, myId int) {
    max = 0
    seats := [128*8]bool{false}
    for _, listString := range data {
        row := binarySearch(listString[0:7], 'B', 128)
        seat := binarySearch(listString[7:], 'R', 8)
        id := (row * 8) + seat
        seats[id] = true
        log.Printf("%d - Row: %d, Seat: %d", id, row, seat)

        if id > max {max = id}
    }

    for i, val := range seats {
       if val == false && i != 0 {
           if seats[i-1] && seats[i+1] {
               myId = i
               break
           }
       }
    }

    return
}

func binarySearch(instr string, downChar uint8, size int) int {
    start := 0
    end := size - 1
    result := 0

    for i := 0; i < len(instr); i++ {
        if instr[i] == downChar {
            start = start + ((end - start) / 2) + 1
        } else {
            end = start + ((end - start) / 2)
        }
        //log.Printf("%d (%s): %d, %d", i, string(instr[i]), start, end)
        if start == end {
            result = start
        }
    }
    return result
}


func readData(test bool) []string {
    if test {
        test := []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
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