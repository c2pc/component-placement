package algorithms

type Chromosome interface {
}

type ChromosomeComponent struct {
	Sizes  Size
	Center Center
}

func NewChromosomes(chroms interface{}) ([]ChromosomeComponent, error) {
	var chromosomes []ChromosomeComponent

	switch chroms.(type) {
	case [][]int:
		for _, chromosome := range chroms.([][]int) {
			chrome, err := NewChromosome(Size{Height: chromosome[0], Width: chromosome[1]})
			if err != nil {
				return nil, err
			}

			chromosomes = append(chromosomes, chrome)
		}
	case []map[string]int:
		for _, chromosome := range chroms.([]map[string]int) {
			chrome, err := NewChromosome(Size{Height: chromosome["height"], Width: chromosome["width"]})
			if err != nil {
				return nil, err
			}

			chromosomes = append(chromosomes, chrome)
		}
	default:
		return nil, ErrChromosomeInvalidType
	}

	return chromosomes, nil
}

func NewChromosome(size Size) (ChromosomeComponent, error) {
	if size.Width == 0 || size.Height == 0 {
		return ChromosomeComponent{}, ErrChromosomeSizeEmpty
	}
	if size.Width < 0 || size.Height < 0 {
		return ChromosomeComponent{}, ErrChromosomeSizeNegative
	}

	return ChromosomeComponent{
		Sizes: size,
		Center: Center{
			Y: (size.Height + size.Height%2) / 2,
			X: (size.Width + size.Width%2) / 2,
		},
	}, nil
}
