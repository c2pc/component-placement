package ga

import "ComponentPlacement/algorithms"

type Ga algorithms.AlgorithmComponent

func Default() algorithms.Algorithm {
	return &Ga{
		Matrix:      algorithms.MatrixComponent{},
		Board:       algorithms.BoardComponent{},
		Chromosomes: make([]algorithms.ChromosomeComponent, 0),
		Populations: make([]algorithms.PopulationComponent, 0),
	}
}

func (m *Ga) Init(matrix [][]int, board interface{}, chromosomes interface{}) error {
	var err error
	m.Chromosomes, err = algorithms.InitChromosomes(chromosomes)
	if err != nil {
		return err
	}

	m.Board, err = algorithms.InitBoards(board)
	if err != nil {
		return err
	}

	m.Matrix, err = algorithms.InitMatrix(matrix, len(m.Chromosomes))
	if err != nil {
		return err
	}

	m.Populations = algorithms.InitPopulations(m.Chromosomes, m.Board)

	return nil
}

func (m *Ga) Run() error {

	return nil
}
