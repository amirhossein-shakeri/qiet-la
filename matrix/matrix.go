package matrix

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
)

type Matrix [][]float64

func (m Matrix) ToSlice() [][]float64 {
	return m
}

func (m Matrix) Print() {
	for _, row := range m {
		fmt.Println(row)
	}
}

func (m Matrix) Reduce() {
	// reduce row echelon
	fmt.Println("Reducing Matrix")
	for i, row := range m {
		originalValue := m[i][i]
		if originalValue != 1 {
			for j, val := range row {
				m[i][j] /= originalValue
				fmt.Printf("Pivot reduction: %4.2f -> %4.2f\n", val, row[j])
			}
		}
		fmt.Printf("Pivot is now %.2f at {%d,%d}\n", m[i][i], i, i)
		m.Print()
		for k, targetRow := range m {
			if k == i { // skip the pivot
				continue
			}
			ratio := targetRow[i] / m[i][i]
			for j := i; j < len(targetRow); j++ {
				fmt.Printf("-(%4.2f)x{%d,%d} -> {%d,%d}\n", ratio, j, i, j, k)
				// fmt.Printf("-(%4.2f)x{%d,%d} -> {%d,%d}: %4.2f - %4.2f = %4.2f\n", ratio, j, i, j, k, targetRow[j], ratio*m[i][j], targetRow[j])
				targetRow[j] -= ratio * m[i][j]
				m.Print()
			}
		}
	}
	fmt.Println("Reduction Complete")
}

func LoadFromFile(file io.Reader) Matrix {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Error parsing the file: %v", err.Error())
	}
	m := GenerateFromStringArray(data)
	fmt.Println("Loaded matrix data from csv file:")
	m.Print()
	return m
}

func (m Matrix) Solve() (string, []float64) {
	// fmt.Printf("Rows=%d, Cols=%d\n", len(m), len(m[0]))
	if len(m)+1 < len(m[0]) {
		return "The provided matrix, has infinite answers", []float64{}
	}
	m.Reduce()
	// todo: check if there is anything like 8=0 or something like that and then tell that there is no answer
	answers := make([]float64, len(m))
	for order, answer := range m {
		answers[order] = math.Round(answer[len(answer)-1])
	}
	return "Matrix was solved", answers
}

func GenerateFromStringArray(arr [][]string) Matrix {
	output := make(Matrix, len(arr))
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
