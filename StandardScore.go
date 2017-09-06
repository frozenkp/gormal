package gormal

import (
	"fmt"
	"os"

	"github.com/montanaflynn/stats"
)

// StandardScoreSingle normalizes the first row with "Standard Score".
//
// (usually only one row in input)
func (d CustomData) StandardScoreSingle() []float64 {
	col := make([]float64, 0)
	out := make([]float64, 0)
	for i := 0; i < d.RawData.ColSize(); i++ {
		col = append(col, d.RawData.Value(0, i))
	}
	mean, err := stats.Mean(col)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	stderv, err := stats.StandardDeviationPopulation(col)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < d.RawData.ColSize(); i++ {
		out = append(out, (d.RawData.Value(0, i)-mean)/stderv)
	}
	return out
}

// StandardScore normalizes every columns with "Standard Score".
func (d CustomData) StandardScore() [][]float64 {
	out := make([][]float64, d.RawData.RowSize())
	for k := 0; k < len(out); k++ {
		out[k] = make([]float64, d.RawData.ColSize())
	}
	for k := 0; k < d.RawData.ColSize(); k++ {
		row := make([]float64, 0)
		for i := 0; i < len(out); i++ {
			row = append(row, d.RawData.Value(i, k))
		}
		mean, err := stats.Mean(row)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		stderv, err := stats.StandardDeviationPopulation(row)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for i := 0; i < len(out); i++ {
			out[i][k] = (d.RawData.Value(i, k) - mean) / stderv
		}
	}
	return out
}

// StandardScoreSingle normalizes the first row with "Standard Score".
//
// (usually only one row in input)
func StandardScoreSingle(data []float64) []float64 {
	col := make([]float64, 0)
	out := make([]float64, 0)
	for i := 0; i < len(data); i++ {
		col = append(col, data[i])
	}
	mean, err := stats.Mean(col)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	stderv, err := stats.StandardDeviationPopulation(col)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < len(data); i++ {
		out = append(out, (data[i]-mean)/stderv)
	}
	return out
}

// StandardScore normalizes every columns with "Standard Score".
func StandardScore(data [][]float64) [][]float64 {
	out := make([][]float64, len(data))
	if len(data) == 0 {
		return out
	}
	for k := 0; k < len(out); k++ {
		out[k] = make([]float64, len(data[0]))
	}
	for k := 0; k < len(data[0]); k++ {
		row := make([]float64, 0)
		for i := 0; i < len(out); i++ {
			row = append(row, data[i][k])
		}
		mean, err := stats.Mean(row)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		stderv, err := stats.StandardDeviationPopulation(row)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for i := 0; i < len(out); i++ {
			out[i][k] = (data[i][k] - mean) / stderv
		}
	}
	return out
}
