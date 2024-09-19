package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var dataLength int

func main() {
	myArgs := os.Args[1:]

	if len(myArgs) != 1 {
		fmt.Println("enter the file name")
		return
	}

	file, err := os.ReadFile(myArgs[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(file), "\n")
	myNbrs := []int{}

	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			nbr, _ := strconv.Atoi(lines[i])
			myNbrs = append(myNbrs, nbr)
		}
	}
	dataLength = len(myNbrs)
	fmt.Println("Average : ", math.Round(getAverage(myNbrs)))
	fmt.Println("Median : ", getMedian(myNbrs))
	fmt.Printf("Variance : %.f\n", getVariance(myNbrs))
	fmt.Println("Deviation : ", getDeviation(getVariance(myNbrs)))
}

func getAverage(nbrs []int) float64 {
	var total float64
	for _, v := range nbrs {
		total += float64(v)
	}
	return total / float64(len(nbrs))
}

func getMedian(nbrs []int) int {
	Median := 0.0
	middle := 0

	for i := 0; i < dataLength-1; i++ {
		for j := i + 1; j < dataLength; j++ {
			if nbrs[i] > nbrs[j] {
				temp := nbrs[i]
				nbrs[i] = nbrs[j]
				nbrs[j] = temp
			}
		}
	}
	//fmt.Println("dataLength: ", dataLength)
	if dataLength%2 == 0 {

		middle = (dataLength / 2)

		Median = float64((nbrs[middle] + nbrs[middle-1]) / 2)

		// fmt.Printf("middle: %d",   int(math.Round(float64((nb0 + nb1) / 2))))
	} else { 

		Median = float64(nbrs[dataLength / 2])
	}

	return int(math.Round(float64(Median)))
}

func getVariance(nbrs []int) float64 {
	var Variance float64
	myAvr := float64(getAverage(nbrs))

	for i := 0; i < len(nbrs); i++ {
		Variance += (float64(nbrs[i]) - myAvr) * (float64(nbrs[i]) - myAvr)
	}

	return math.Round(Variance / float64(len(nbrs)))
}

func getDeviation(varince float64) float64 {
	return math.Round(math.Sqrt(varince))
}
