package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// maybe we'd better change the name of the file to something else like problem solver. the same with the main function name

func main() {
	fmt.Println("===========================================================================================")
	fmt.Println("Linear Algebra Equations Solver by Amirhossein Shakeri (amirhossein.shakeri.work@gmail.com)")
	fmt.Println("===========================================================================================")

	// get user input to read the file name
	filePath := Input("Please enter the file path: ", "./3x2.equation.csv")

	// read the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening %s: %v", filePath, err.Error())
	}
	defer file.Close()

	// parse data
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Error parsing %s: %v", filePath, err.Error())
	}
	// fmt.Printf("Parsed Data: %+v\n", data)

	// Adapt data
	matrix := StringToFloat(data)

	// Solve the problem
	SolveEquation(matrix)

	// show the answer
}

func SolveEquation(m [][]float64) []float64 {
	/* x=2, y=1
	[
		[3,3,9],
		[2,4,8],
	]
	*/
	// i, j := 0, 0
	// ValidateEquation(m) //? if # of rows & cols not valid, panic...
	output := m // FIXME: change m to output in the code
	var p *float64
	for i, row := range output {
		p = &m[i][i]
		pValue := *p
		// if the m[i][i] != 1, then divide the entire line/row by m[i][i]
		if m[i][i] != 1 {
			for j, val := range row {
				m[i][j] /= pValue
				fmt.Printf("%.2f -> %.2f\n", val, m[i][j])
			}
		}
		fmt.Printf("Pivot is now %.2f\n", *p)

		// check if the bottom or elements of 1 are not 0, enter a loop to make them 0
		for j, val := range row {
			fmt.Printf("Passing x=%d y=%d val=%.2f \n", i, j, val)
			// ?if iterating over the last column, skip? I think it's not gonna happen if we validate the # of rows & cols
		}
	}
	return []float64{0, 0}
}

func StringToFloat(arr [][]string) [][]float64 {
	// var output [len(arr)][len(arr[0])]float64 // error: must be constant
	// output := make([][]float64, len(arr), len(arr[0])) // index error
	output := make([][]float64, len(arr))
	for i := range output {
		output[i] = make([]float64, len(arr[0]))
	}
	for i, row := range arr {
		for j, val := range row {
			r, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("Error converting string to float!", err.Error())
			}
			output[i][j] = r
		}
	}
	return output
}
