package main

import (
	"ComponentPlacement/ga"
	"fmt"
)

var (
	matrix = [][]int{{1, 2, 2, 4}, {1, 2, 2, 4}, {1, 2, 2, 4}, {1, 2, 2, 4}}

	chromosome1 = append(make([]map[string]int, 0), map[string]int{"width": 4, "height": 5}, map[string]int{"width": 4, "height": 5}, map[string]int{"width": 4, "height": 5}, map[string]int{"width": 4, "height": 5})
	chromosome2 = append(make([][]int, 0), []int{3, 2}, []int{3, 2}, []int{3, 2}, []int{3, 2})

	board1 = map[string]int{"width": 4, "height": 5}
	board2 = []int{1, 7}
)

func main() {
	algorithm := ga.Default()
	if err := algorithm.Init(matrix, board1, chromosome1); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := algorithm.Run(); err != nil {

	}
}
