package algorithms


type Matrix interface {
}

type MatrixComponent struct {
	Matrix [][]int
}

func NewMatrix(matrix [][]int, length int) (MatrixComponent, error) {
	array, err := parseMatrix(matrix, length)
	if err != nil {
		return MatrixComponent{}, err
	}

	return MatrixComponent{
		Matrix: array,
	}, nil
}

func parseMatrix(matrix [][]int, length int) ([][]int, error) {

	if len(matrix) == 0 {
		return nil, ErrMatrixEmpty
	}
	if len(matrix) < length {
		return nil, ErrMatrixSmallSize
	}

	var array [][]int
	for i := range matrix {
		if len(matrix[i]) < length {
			return nil, ErrMatrixSmallSize
		}
		for j := range matrix[i]{
			if matrix[i][j] < 0 {
				return nil, ErrMatrixElementNegative
			}
		}
		array = append(array, matrix[i][:length])
	}

	return array, nil
}
