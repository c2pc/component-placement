package algorithms

import (
	"math/rand"
	"time"
)

type Algorithm interface {
	Init(matrix [][]int, board interface{}, chromosomes interface{}) error
	Run() error
}

type AlgorithmComponent struct {
	Matrix      MatrixComponent
	Populations []PopulationComponent
	Chromosomes []ChromosomeComponent
	Board       BoardComponent
}

type Size struct {
	Width  int
	Height int
}

type Center struct {
	X int
	Y int
}

type PopulationComponent struct {
	chromosomes []int
	board       Board
}

func Default() Algorithm {
	return &AlgorithmComponent{
		Matrix:      MatrixComponent{},
		Board:       BoardComponent{},
		Chromosomes: make([]ChromosomeComponent, 0),
		Populations: make([]PopulationComponent, 0),
	}
}

func (a *AlgorithmComponent) Init(matrix [][]int, board interface{}, chromosomes interface{}) error {
	panic("implements Init")
}

func (a *AlgorithmComponent) Run() error {
	panic("implements Run")
}

func InitChromosomes(chromosomes interface{}) ([]ChromosomeComponent, error) {
	return NewChromosomes(chromosomes)
}

func InitBoards(size interface{}) (BoardComponent, error) {
	return NewBoard(size)
}

func InitMatrix(matrix [][]int, length int) (MatrixComponent, error) {
	return NewMatrix(matrix, length)
}

func InitPopulations(chromosomes []ChromosomeComponent, board Board) []PopulationComponent {
	chroms := make([]int, len(chromosomes))
	for i := range chroms {
		chroms[i] = i
	}

	return []PopulationComponent{{chromosomes: Random(chroms), board: board}}
}

func Random(chroms []int) []int{

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(chroms),
		func(i, j int) { chroms[i], chroms[j] = chroms[j], chroms[i] })

	return chroms
}