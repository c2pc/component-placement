package algorithms

import "errors"

var (
	ErrChromosomeSizeNegative = errors.New("chromosome size is negative")
	ErrChromosomeSizeEmpty = errors.New("chromosome size is empty")
	ErrChromosomeInvalidType = errors.New("chromosome type is invalid")

	ErrBoardInvalidType = errors.New("board type is invalid")
	ErrBoardSizeEmpty = errors.New("board size is empty")
	ErrBoardSizeNegative = errors.New("board size is negative")

	ErrMatrixElementNegative = errors.New("matrix element is negative")
	ErrMatrixSmallSize = errors.New("matrix size is less than chromosomes count")
	ErrMatrixEmpty = errors.New("matrix is empty")
)
