package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var totalSum64 int64

func idFinder64(start, end int64) {
	// Placeholder for ID finding logic
	//fmt.Println("Finding IDs in range:", start, "to", end)

	for i := start; i <= end; i++ {
		// Simulate processing each ID
		str := strconv.FormatInt(i, 10)
		//_ = str // Use str as needed

		//fmt.Println("ID = ", str, "length =", len(str))

		if len(str) % 2 == 0 {
			if str[0:len(str)/2] == str[len(str)/2:] {
				fmt.Println("Found matching ID:", str)
				totalSum64 += i
			}
		}
	}	
}

func dataProcessing64(path string) error {
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
			idFinder64(start, end)
			//fmt.Println(rangeSlice)
		} else {
			if len(line) > 0 {
				rangeSlice := strings.Split(line, "-")
				start, _ := strconv.ParseInt(rangeSlice[0], 10, 64)
				end, _ := strconv.ParseInt(rangeSlice[1], 10, 64)
				idFinder64(start, end)
				//fmt.Println(rangeSlice)
			}
			break
		}
	}

	return nil
}

func main() {
	err := dataProcessing64("indi.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Total Sum =", totalSum64_part2)
}
