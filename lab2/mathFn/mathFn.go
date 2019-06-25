package mathFn

import (
	"math"

	"github.com/pityara/labs/ga/lab2/population"
)

type Solution struct {
	X []float64 `json:"x"`
	Y float64   `json:"y"`
}

type sumFn func([]float64) float64

/* f2(x)=sum(100·(x(i+1)-x(i)^2)^2+(1-x(i))^2) */

func getSolution(x []float64) Solution {
	fnVals := make([]float64, len(x))

	for i := 0; i < len(x); i += 2 {
		fnVals[i] = 100*math.Pow((x[i+1]-math.Pow(x[i], 2)), 2) + math.Pow((1-x[i]), 2)
	}

	return Solution{x, sum(fnVals)}
}

func sum(values []float64) float64 {
	result := .0

	for _, value := range values {
		result += value
	}

	return result
}

/* Расчёт функции для всей популяции */
func MakePopulationSolutions(population []population.Person) []Solution {
	solutions := make([]Solution, len(population))
	for index, person := range population {
		solutions[index] = getSolution(person)
	}

	return solutions
}
