package main

import (
	"flag"
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)



func main() {
	filePath := flag.String("filePath", "./input.txt","yeet")
	flag.Parse()
	reports := parseFile(*filePath," ")
	safeCount := 0;
	shitCount := 0;
	for _, report := range reports {
		if (ReportIsSafe(report,1,3)) {
			safeCount = safeCount + 1
		} else {
			shitCount = shitCount + 1
		}
	}
	
	fmt.Println(safeCount,shitCount)

}

func ReportIsSafe(report []int, minDist,maxDist int) bool {

	if !(IsAscVec(report)) && !(IsDescVec(report)) {
		return false
	}
	if !(ValidAdjacentDiffs(report, minDist,maxDist)) {
		return false
	}


	return true
}

func ValidAdjacentDiffs(vec []int, minDist, maxDist int) bool {
	for i := 0; i < len(vec) -1; i++ {
		dist := IntAbsDist(vec[i],vec[i+1])
		if (dist < minDist) || (dist > maxDist) {
			return false
		}
	}
	
	return true
}

func IsAscVec(vec []int) bool {
	for i := 0; i < len(vec) -1 ; i++ {
		if vec[i] > vec[i+1] {
			return false
		}
	}

	return true
}

func IsDescVec(vec []int) bool {
	for i := 0; i < len(vec) -1 ; i++ {
		if vec[i] < vec[i+1] {
			return false
		}
	}

	return true
}

func IntAbsDist(x,y int) int {
	if x >= y{
		return x - y
	} else {
		return y - x
	}
}

func parseRow(row []string) []int {
	x := make([]int,len(row))
	for i := 0; i < len(row); i++ {
		fmt.Println(row[i])
		val, err := strconv.Atoi(row[i])
		if err != nil {
			log.Fatal(err)
		}
		x[i] = val
	}
	return x
}

func parseFile(filePath string, sep string) ([][]int){
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)
	
	reports := [][]int{}
	
    for scanner.Scan() {
        reports = append(reports,parseRow(strings.Split(scanner.Text(),sep) ))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


	return reports
}

