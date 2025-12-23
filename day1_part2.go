package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var totalCount int = 50
var password int = 0

func rotationLogic(direction byte, distance int) {
	
	switch  {
		case distance > 900:
			distance = distance - 900
			password += 9
		case distance > 800:
			distance = distance - 800
			password += 8
		case distance > 700:	
			distance = distance - 700
			password += 7
		case distance > 600:
			distance = distance - 600
			password += 6
		case distance > 500:
			distance = distance - 500
			password += 5
		case distance > 400:
			distance = distance - 400
			password += 4
		case distance > 300:
			distance = distance - 300
			password += 3
		case distance > 200:
			distance = distance - 200
			password += 2
		case distance > 100:
			distance = distance - 100	
			password += 1	
	}
			
	

	switch direction {
	case 'L':
		previousCount := totalCount
		totalCount -= distance

		if previousCount == 0 {
			totalCount = 100 + totalCount
		}
		
		if totalCount < 0 {
			totalCount = 100 + totalCount
			password += 1
		}
		if totalCount == 0 {
			password += 1
		}

	case 'R':
		previousCount := totalCount
		totalCount += distance


		if totalCount == 100 {
			totalCount = totalCount - 100
		}	

		if totalCount > 100 && previousCount != 100 {
			totalCount = totalCount - 100
			password += 1
		}
		if totalCount == 0 {
			password += 1
		}
				
	default:
		fmt.Println("Invalid direction")
		os.Exit(1)
	}			

	fmt.Printf("Rotating %c by %d units -> %d\n", direction, distance, totalCount)
}

func fileProcessing(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != io.EOF {
			line = strings.Trim(line, "\r\n")

			b := line[0]  // must be the direction of rotation
			n, _ := strconv.Atoi(line[1:])  // must be the distance
			rotationLogic(b, n)
		} else {
			if len(line) != 0 {
				b := line[0]  // must be the direction of rotation
				n, _ := strconv.Atoi(line[1:])  // must be the distance
				rotationLogic(b, n)
			}
			break
		}	
	}	
	return nil
}

func main() {
	err := fileProcessing("largeinput.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Password:", password)
}
