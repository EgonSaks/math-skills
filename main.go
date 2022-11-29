package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read the data in
	readFile, _ := os.ReadFile("data.txt")
	// convert bytes to string with new line
	slice := strings.Split(string(readFile), "\n")
	f := sliceStringToSliceFloat(slice)

	fmt.Printf("Average: %v\n", average(f))
	fmt.Printf("Median: %v\n", median(f))
	fmt.Printf("Variance: %.f\n", variance(f))
	fmt.Printf("Standard Deviation: %v\n", standardDeviation(f))
}

// sliceStringToSliceFloat convert slice string to slice float64
func sliceStringToSliceFloat(s []string) []float64 {
	var f []float64

	for i := 0; i < len(s)-1; i++ {
		temp, _ := strconv.Atoi(s[i])
		f = append(f, float64(temp))
	}
	return f
}

// average get an average from an array
func average(a []float64) float64 {
	length := float64(len(a))
	sum := float64(0)

	for _, value := range a {
		sum = sum + float64(value)
	}
	return math.RoundToEven((sum / length))
}

// median get a median from an array
func median(a []float64) float64 {
	sort.Float64s(a)

	var median float64
	length := len(a)

	if length == 0 {
		return 0
	} else if length%2 == 0 {
		median = (a[length/2-1] + a[length/2]) / 2
	} else {
		median = a[length/2]
	}
	return math.RoundToEven(median)
}

// variance get a variance from an array
func variance(a []float64) float64 {
	avg := average(a)
	deviations := []float64{}
	sum := float64(0)

	for i := 0; i < len(a); i++ {
		deviations = append(deviations, (math.Pow((float64(a[i]) - avg), 2)))
		sum = sum + float64(deviations[i])
	}
	return math.RoundToEven((sum / float64(len(a))))
}

// standardDeviation get standard deviation from an array
func standardDeviation(a []float64) float64 {
	squareRoot := math.Sqrt(float64(variance(a)))
	return math.RoundToEven(squareRoot)
}
