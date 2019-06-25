package mathFn

import (
	"math"

	"github.com/pityara/labs/ga/lab1/slicing"
)

type Solution struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

//Cos(2x)/x^2 функция
func calcFunction(x float64) float64 {
	return math.Cos(2.0*x) / math.Pow(x, 2.0)
}

/* Расчёт функции для единичной особи */
func getSolution(x float64) Solution {
	return Solution{x, calcFunction(x)}
}

/* Расчёт функции для всей популяции */
func MakePopulationSolutions(population []string, r slicing.Restr) []Solution {
	solutions := make([]Solution, len(population))
	for index, person := range population {
		x, _ := r.BinaryToValue(person)

		solutions[index] = getSolution(x)
	}

	return solutions
}
