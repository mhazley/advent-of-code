package main

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    input := readData(false)
    log.Println("start")
    log.Println(check(input))
}

func check(data []string) int {
    requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
    var count = 0
    for _, line := range data {
        //log.Println(line)
        r, _ := regexp.Compile("(\\S*):(\\S*)*")

        keyValuePairs := r.FindAllString(line, -1)

        m := make(map[string]string)
        valid := true

        for _, pair := range keyValuePairs {
            keyValue := strings.Split(pair, ":")
            m[keyValue[0]] = keyValue[1]
        }

        // Check keys are present before bothering to parse
        for _, key := range requiredKeys {
            _, ok := m[key]
            if !ok {
                valid = false
                break
            }
        }

        // ok, valid so far... parse
        if valid {
            for _, key := range requiredKeys {
                val, _ := m[key]

                switch key {
                case "byr":
                    if len(val) == 4 {
                        if intVal, err := strconv.Atoi(val); err == nil {
                            if (intVal < 1920) || (intVal > 2002) {
                                valid = false
                                break
                            }
                        }
                    } else {
                        valid = false
                        break
                    }
                case "iyr":
                    if len(val) == 4 {
                        if intVal, err := strconv.Atoi(val); err == nil {
                            if (intVal < 2010) || (intVal > 2020) {
                                valid = false
                                break
                            }
                        }
                    } else {
                        valid = false
                        break
                    }
                case "eyr":
                    if len(val) == 4 {
                        if intVal, err := strconv.Atoi(val); err == nil {
                            if (intVal < 2020) || (intVal > 2030) {
                                valid = false
                                break
                            }
                        }
                    } else {
                        valid = false
                        break
                    }
                case "hgt":
                    unit := val[len(val)-2:]
                    if intVal, err := strconv.Atoi(val[0:len(val)-2]); err == nil {
                        if unit == "cm" {
                            if (intVal < 150) || (intVal > 193) {
                                valid = false
                                break
                            }
                        } else if unit == "in" {
                            if (intVal < 59) || (intVal > 76) {
                                valid = false
                                break
                            }
                        }
                    } else {
                    valid = false
                    break
                }
                case "hcl":
                    r, _ := regexp.Compile("^#[0123456789a-fA-F]{6}$")
                    if !r.MatchString(val) {
                        valid = false
                        break
                    }
                case "ecl":
                    validEyes := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
                    found := false
                    for _, listVal := range validEyes {
                        if val == listVal {
                            found = true
                            break
                        }
                    }
                    if !found {
                        valid = false
                        break
                    }
                case "pid":
                    r, _ := regexp.Compile("^[0123456789]{9}$")
                    if !r.MatchString(val) {
                        valid = false
                        break
                    }
                }
            }
        }
        if valid {
            count++
        }
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
    var line = ""

    for scanner.Scan() {
        text := scanner.Text()
        if len(text) > 0 {
             line = line + " " + text
        } else {
            input = append(input, line)
            line = ""
        }
    }
    // Last input needs to append!
    input = append(input, line)

    return input
}
