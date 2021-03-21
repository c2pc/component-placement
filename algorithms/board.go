package algorithms

type Board interface {
}

type BoardComponent struct {
	Sizes Size
	Board [][]int
}

func NewBoard(sizes interface{}) (BoardComponent, error) {
	size, err := parseSize(sizes)
	if err != nil {
		return BoardComponent{}, err
	}
	return BoardComponent{
		Sizes: size,
		Board: initBoard(size),
	}, nil
}

func initBoard(sizes Size) [][]int {
	board := make([][]int, sizes.Height)
	for i := range board {
		board[i] = make([]int, sizes.Width)
		for j := range board[i] {
			board[i][j] = -1
		}
	}
	return board
}

func parseSize(sizes interface{}) (Size, error) {
	var height, width int
	switch sizes.(type) {
	case []int:
		height, width = sizes.([]int)[0], sizes.([]int)[1]
	case map[string]int:
		height, width = sizes.(map[string]int)["height"], sizes.(map[string]int)["width"]
	default:
		return Size{}, ErrBoardInvalidType
	}

	if height == 0 || width == 0 {
		return Size{}, ErrBoardSizeEmpty
	}
	if height < 0 || width < 0 {
		return Size{}, ErrBoardSizeNegative
	}

	return Size{Height: height, Width: width}, nil

}
