package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)


func FindMostValuablePath(arr [][]int) int {
	if len(arr) == 0 {
		return 0
	}

	// Initialize a 2D slice to store the maximum values for each element in the array
	maxValues := make([][]int, len(arr))
	for i := range maxValues {
		maxValues[i] = make([]int, len(arr[i]))
	}

	// Initialize the first element of maxValues with the first element of arr
	maxValues[0][0] = arr[0][0]

	// Calculate the maximum values for each element in maxValues
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			maxAbove := 0
			if i-1 >= 0 && j < len(arr[i-1]) {
				maxAbove = maxValues[i-1][j]
			}

			maxLeft := 0
			if i-1 >= 0 && j-1 >= 0 {
				maxLeft = maxValues[i-1][j-1]
			}

			// Update the current max value with the sum of the current element and the maximum of maxAbove and maxLeft
			maxValues[i][j] = arr[i][j] + max(maxAbove, maxLeft)
		}
	}

	// Find the maximum value in the last row of maxValues
	lastRow := maxValues[len(maxValues)-1]
	maxValue := 0
	for _, v := range lastRow {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	// arr := [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }

	// result := FindMostValuablePath(arr)
	// fmt.Println("The most valuable path total:", result)

	// Open the JSON file
	file, err := os.Open("files/hard.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Read the JSON data from the file
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse the JSON data into a 2D integer slice
	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Find the most valuable path
	result := FindMostValuablePath(triangle)
	fmt.Println("The most valuable path total:", result)
}