package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var totalSum64_part2 int64 = 0

func checkIfAllDigitsSame(num int64) bool {
	dight := num % 10
	for num > 0 {
		if num % 10 != dight {
			return false
		}
		num /= 10
	}
	return true
}

func checkIfOneDigit(num int64) bool {
	if num >=0 && num <=9 {
		return true
	}
	return false
}	

func idFinder64_part2(start, end int64) {
	for i := start; i <= end; i++ {
		if checkIfOneDigit(i) {
			continue
		}

		if checkIfAllDigitsSame(i) {
			totalSum64_part2 += i
			continue
		}

		str := strconv.FormatInt(i, 10)

		for j := 1; j <= len(str)/2; j++ {
			isIDValid := true
			var partition int

			if len(str) % j != 0 {
				continue
			} else {	
				partition = j
			}

			for k := 1; k <= ((len(str)/partition) - 1); k++ {
				if str[0:partition] == str[partition*k:partition*(k+1)] {
					if k == ((len(str)/partition) - 1) {
						totalSum64_part2 += i
						goto LABEL
					}
					continue
				} else {
					isIDValid = false
					break
				}
			}
			if isIDValid == false{
				continue
			}else {
				totalSum64_part2 += i
			}
		}	
	  LABEL:
	}	
}
func dataProcessing64_part2(path string) error {
	fmt.Println("Processing string at path:", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(',')
		line = strings.Trim(line, ",")

		if err != io.EOF {
			rangeSlice := strings.Split(line, "-")
			start, _ := strconv.ParseInt(rangeSlice[0], 10, 64)
			end, _ := strconv.ParseInt(rangeSlice[1], 10, 64)
			idFinder64_part2(start, end)
			//fmt.Println(rangeSlice)
			//fmt.Println("Processing range:", start, "to", end)
		} else {
			if len(line) > 0 {
				rangeSlice := strings.Split(line, "-")
				start, _ := strconv.ParseInt(rangeSlice[0], 10, 64)
				end, _ := strconv.ParseInt(rangeSlice[1], 10, 64)
				idFinder64_part2(start, end)
				//fmt.Println(rangeSlice)
				//fmt.Println("Processing range:", start, "to", end)
			}
			break
		}
	}

	return nil
}

func main() {
	err := dataProcessing64_part2("indi.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Total Sum =", totalSum64_part2)
}
