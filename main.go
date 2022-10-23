package main

import (
	"amirhossein-shakeri/go-linear-algebra/matrix"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// TODO: GO fractinal/rational number(search it)
// maybe we'd better change the name of the file to something else like problem solver. the same with the main function name

func main() {
	fmt.Println("==========================================================================================")
	fmt.Println("Linear Algebra Equation Solver by Amirhossein Shakeri (amirhossein.shakeri.work@gmail.com)")
	fmt.Println("==========================================================================================")

	// get user input to read the file name
	// filePath := "./3x2.equation.csv"
	filePath := Input("Please enter the file path: ", "./3x2.equation.csv")

	// read the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening %s: %v", filePath, err.Error())
	}
	defer file.Close()

	mtx := matrix.LoadFromFile(file)

	mtx.Reduce()

	msg, answers := mtx.Solve()
	fmt.Println(msg)
	if len(answers) > 0 {
		fmt.Println("Answers: ", answers)
	}

	// mtx.Print()
	// parse data
	// csvReader := csv.NewReader(file)
	// data, err := csvReader.ReadAll()
	// if err != nil {
	// 	log.Fatalf("Error parsing %s: %v", filePath, err.Error())
	// }
	// fmt.Printf("Parsed Data: %+v\n", data)

	// Adapt data
	// matrix := StringToFloat(data)

	// Solve the problem
	// answers := SolveEquation(matrix)

	// show the answer
	// fmt.Println("Answers: ", answers)
}

func SolveEquation(m matrix.Matrix) []float64 {
	// ValidateEquation(m) //? if # of rows & cols not valid, panic?...
	mtx := m
	var p *float64
	for i, row := range mtx {
		p = &mtx[i][i]
		pValue := *p
		if pValue != 1 {
			for j, val := range row {
				mtx[i][j] /= pValue
				fmt.Printf("%.2f -> %.2f\n", val, mtx[i][j])
			}
		}
		fmt.Printf("Pivot is now %.2f at {%d,%d}\n", *p, i, i)
		// fmt.Println(output)

		// check if the bottom or elements of 1 are not 0, enter a loop to make them 0
		// if i != 0 {
		// resetAbove = true
		for k, targetRow := range mtx {
			if k == i { // skip the pivot
				continue
			}
			ratio := targetRow[i] / mtx[i][i]
			for j := i; j < len(targetRow); j++ {
				fmt.Printf("-%.2fx{%d,%d} -> {%d,%d}: %.2f - %.2f = %.2f | Matrix: \n", ratio, j, i, j, k, targetRow[j], ratio*mtx[i][j], targetRow[j])
				targetRow[j] -= ratio * mtx[i][j]
				// fmt.Printf("%v\n", matrix)
				mtx.Print()
			}
		}
	}

	answers := make([]float64, len(mtx))
	for order, answer := range mtx {
		answers[order] = math.Round(answer[len(answer)-1])
	}
	return answers
}

func StringToFloat(arr [][]string) matrix.Matrix {
	// var output [len(arr)][len(arr[0])]float64 // error: must be constant
	// output := make(matrix.Matrix, len(arr), len(arr[0])) // index error
	output := make(matrix.Matrix, len(arr))
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
