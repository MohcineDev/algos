package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("nbrs.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(file), "\n")
	myNbrs := []int{}

	for i := 0; i < len(lines); i++ {
		nbr, _ := strconv.Atoi(lines[i])
		myNbrs = append(myNbrs, nbr)
	}

	fmt.Println("Average : ", getAverage(myNbrs, len(myNbrs)))
	fmt.Println("Median : ", getMedian(myNbrs))
	fmt.Println("Variance : ", getVariance(myNbrs))
}

func getAverage(nbrs []int, length int) int {
	total := 0
	for i := 0; i < len(nbrs); i++ {
		total += nbrs[i]
	}
	return total / length
}

func getMedian(nbrs []int) int {
	Median := 0
	for i := 0; i < len(nbrs)-1; i++ {
		for j := i + 1; j < len(nbrs); j++ {
			if nbrs[i] > nbrs[j] {
				temp := nbrs[i]
				nbrs[i] = nbrs[j]
				nbrs[j] = temp
			}
		}
	}
	if len(nbrs)%2 == 0 {

		middle := len(nbrs) / 2
		nb0 := nbrs[middle]
		nb1 := nbrs[middle-1]
		Median = (nb0 + nb1) / 2
	}
	return Median
}

func getVariance(nbrs []int) float64{
	var Variance float64
	myAvr := float64(getAverage(nbrs, len(nbrs)))
	for i := 0; i < len(nbrs); i++ {
		Variance += (float64(nbrs[i]) - myAvr) * (float64(nbrs[i]) - myAvr)

	}
	return math.Round(Variance / float64(len(nbrs)))
}
