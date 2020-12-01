package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main()  {
	input := readData(false)
	log.Printf("Input Data: %v\n", input)

	result := findNumsForTarget(2020, 2, input)
	log.Printf("Result One: %d\n", result)

	result = findNumsForTarget(2020, 3, input)
	log.Printf("Result Two: %d\n", result)
}

func findNumsForTarget(target int, numVals int, data []int) int {
	var result = 0
	for i, listNum := range data {
		if numVals > 1 {
			nextListNum := findNumsForTarget(target-listNum, numVals-1, data[i+1:])
			if nextListNum == 0 {
				continue
			} else {
				result = listNum * nextListNum
				break
			}
		} else {
			for _, listNum := range data {
				if listNum == target {
					result = listNum
					break
				}
			}
		}
	}
	return result
}

func readData(test bool) []int {
	if test {
		test := []int{1721, 979, 366, 299, 675, 1456}
		return test
	} else {
		file, err := os.Open("./input.txt")
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)

		var input []int

		for scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			input = append(input, num)
		}

		return input
	}
}