package lab1

import (
	"github.com/pityara/labs/ga/lab1/crossover"
	"github.com/pityara/labs/ga/lab1/mathFn"
	"github.com/pityara/labs/ga/lab1/mutation"
	"github.com/pityara/labs/ga/lab1/population"
	"github.com/pityara/labs/ga/lab1/reproduction"
	"github.com/pityara/labs/ga/lab1/slicing"
)

//Cos(2x)/x2 x при. [-20,-2.3]/
//

/* Структура для входных данных */
type In struct {
	Floor      float64 `json:"floor"`      // нижняя граница
	Ceil       float64 `json:"ceil"`       // верхняя граница
	Precision  int     `json:"precision"`  // точность
	Quatity    int     `json:"quantity"`   // количество особей
	Iterations int     `json:"iterations"` // количество поколений
}

// Основная функция программы
func Resolve(in *In) []mathFn.Solution {
	pow := slicing.CalcPow(in.Floor, in.Ceil, int8(in.Precision))
	r := slicing.Restr{Floor: in.Floor, Ceil: in.Ceil, Precision: in.Precision, Pow: pow}
	population := generateStartPopulation(in, pow)
	reproducedPopulation := make([]string, in.Quatity)
	crossoverPopulation := make([]string, in.Quatity)
	mutatedPopulation := mutation.MakeMutation(population)
	for i := 0; i < in.Iterations; i++ {
		reproducedPopulation = reproduction.MakeReproduction(mutatedPopulation, r)
		crossoverPopulation = crossover.MakeCrossover(reproducedPopulation)
		mutatedPopulation = mutation.MakeMutation(crossoverPopulation)
	}

	solutions := mathFn.MakePopulationSolutions(mutatedPopulation, r)

	return solutions
}

func generateStartPopulation(in *In, pow int8) []string {

	return population.GeneratePopulation(in.Quatity, pow)
}
