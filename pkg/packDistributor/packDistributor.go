package packDistributor

import "sort"

type PackDistributor struct {
	PackSizes []int
}

func NewDistributor(packSizes []int) *PackDistributor {
	sort.Ints(packSizes)

	return &PackDistributor{
		PackSizes: packSizes,
	}
}

func (pd *PackDistributor) Distribute(order int) (map[int]int, error) {
	if order <= 0 {
		return map[int]int{}, nil
	}

	returnMap := make(map[int]int, 0)

	for i := len(pd.PackSizes) - 1; i >= 0; i-- {
		// if the order size is between two pack sizes, return the larger one.
		if pd.PackSizes[i]/order == 1 {
			returnMap[pd.PackSizes[i]] = 1
			order = 0
			break
		}

		// fit the most packs of the remaining ones.
		size := order / pd.PackSizes[i]
		if size != 0 {
			returnMap[pd.PackSizes[i]] = size
			order = order - (pd.PackSizes[i] * size)
		}
	}

	// if we have a remainder, it is less than the smallest pack, so send 1 of it.
	if order != 0 {
		returnMap[pd.PackSizes[0]] = 1
	}

	return returnMap, nil
}
