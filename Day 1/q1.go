package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Row  int     `json:"row"`
	Col  int     `json:"col"`
	Data [][]int `json:"data"`
}

func newMatrix(row, col int) *Matrix {
	m := &Matrix{
		Row:  row,
		Col:  col,
		Data: make([][]int, row),
	}
	for i := range m.Data {
		m.Data[i] = make([]int, col)
	}
	return m
}

func (m *Matrix) GetRow() int {
	return m.Row
}

func (m *Matrix) GetCol() int {
	return m.Col
}

func (m *Matrix) Set(i, j, val int) {
	m.Data[i][j] = val
}

func (m *Matrix) Add(n *Matrix) *Matrix {
	if m.GetRow() != n.GetRow() || m.GetCol() != n.GetCol() {
		fmt.Println("Sizes are different")
		return nil
	}
	res := newMatrix(m.Row, m.Col)
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			res.Data[i][j] = m.Data[i][j] + n.Data[i][j]
		}
	}
	return res
}

func (m *Matrix) PrintAsJSON() {
	jsonData, err := json.Marshal(m) // no indentation
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func main() {
	m1 := newMatrix(2, 2)
	m1.Set(0, 0, 1)
	m1.Set(0, 1, 2)
	m1.Set(1, 0, 3)
	m1.Set(1, 1, 4)

	m2 := newMatrix(2, 2)
	m2.Set(0, 0, 5)
	m2.Set(0, 1, 6)
	m2.Set(1, 0, 7)
	m2.Set(1, 1, 8)

	result := m1.Add(m2)

	fmt.Println("Matrix 1:")
	m1.PrintAsJSON()

	fmt.Println("Matrix 2:")
	m2.PrintAsJSON()

	fmt.Println("Result (M1 + M2):")
	if result != nil {
		result.PrintAsJSON()
	}
}
