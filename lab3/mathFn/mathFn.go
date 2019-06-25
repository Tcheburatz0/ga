package mathFn

import (
	"strconv"

	"github.com/pityara/labs/ga/lab3/population"
)

func GetSolutions(weights [][]string, routes []population.Person) []uint32 {
	result := make([]uint32, len(routes))

	for key, value := range routes {
		firstCity := true
		sum := uint32(0)
		for i := 0; i != 0 || firstCity; i = int(value[i]) {
			firstCity = false
			res, _ := strconv.ParseInt(weights[i][int(value[i])], 10, 32)
			sum += uint32(res)
		}

		result[key] = sum
	}

	return result
}

func GetMinIndex(v []uint32) int {
	min := v[0]
	result := 0
	for key, value := range v {
		if value < min {
			result = key
		}
	}

	return result
}
