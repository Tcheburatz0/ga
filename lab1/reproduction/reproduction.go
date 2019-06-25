package reproduction

import (
	"github.com/pityara/labs/ga/lab1/mathFn"
	"github.com/pityara/labs/ga/lab1/roulette"
	"github.com/pityara/labs/ga/lab1/slicing"
)
/* Функция репродукции */
func MakeReproduction(population []string, r slicing.Restr) []string {
	solutions := mathFn.MakePopulationSolutions(population, r)
	populationLength := len(population)
	/* Инициализация рулетки */
	roulette := roulette.GenerateRoulette(solutions)

	result := make([]string, populationLength)

	/* Подбор особей для репродукции */
	for i := 0; i < populationLength; i++ {
		index := roulette.Spin()
		result[i] = population[index]
	}
	return result
}
