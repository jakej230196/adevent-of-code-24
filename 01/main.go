package main

import (
	"fmt"
	"bufio"
	"flag"
	"os"
	"log"
	"strings"
	"strconv"
	"slices"
)

func parseRow(row []string) (int, int){
	leftVal, err := strconv.Atoi(row[0])
	if err != nil {
		log.Fatal(err)
	}
	rightVal, err := strconv.Atoi(row[1])
	if err != nil {
		log.Fatal(err)
	}
	return leftVal,rightVal
}

func parseFile(filePath string, sep string) ([]int,[]int,error){
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)
	
	leftList := []int{}
	rightList := []int{}
	
    for scanner.Scan() {
        leftVal, rightVal := parseRow(strings.Split(scanner.Text(),sep))
        leftList = append(leftList,leftVal)
        rightList = append(rightList,rightVal)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


	return leftList,rightList,nil
}

func calculateAbsDistanceVec(x,y []int) ([]int,int) {
	distanceVec := make([]int,len(x))
	sum := 0 
	for i := 0 ; i < len(x); i++ {
		if x[i] >= y[i] {
			distanceVec[i] = x[i] - y[i]
		} else {
			distanceVec[i] = y[i] - x[i]
		}
		sum = sum + distanceVec[i]
	}
	return distanceVec,sum
	
}

func ListToOccurenceMap(x []int) map[int]int{
	occurenceMap := make(map[int]int)
	for i := 0; i < len(x); i++ {
		val, exists := occurenceMap[x[i]]
		if !exists {
			occurenceMap[x[i]] = 1
		} else {
			occurenceMap[x[i]] = val + 1
		}
	}
	return occurenceMap
}

func CalculateListSimilarity(left,right []int) int {
	rightOC := ListToOccurenceMap(right)
	similarityScore := 0
	for i := 0; i < len(left) ; i++ {
		mul, exists := rightOC[left[i]]
		if !exists {
			continue
		} else {
			similarityScore = similarityScore + left[i] * mul
		}
	
	}
	return similarityScore
	
}


func main() {
	filePath := flag.String("filePath", "./input.txt","yeet")
	flag.Parse()
	
	leftList,rightList,err := parseFile(*filePath,"   ")
	if err != nil {
		log.Fatal(err)
	}
	slices.Sort(leftList)
	slices.Sort(rightList)
	
	_,sum := calculateAbsDistanceVec(leftList,rightList)
	fmt.Println(sum)
	// end p1
	
	fmt.Println(CalculateListSimilarity(leftList,rightList))
}

































