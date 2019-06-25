package reproduction

import (
	"github.com/pityara/labs/ga/lab2/mathFn"
	"github.com/pityara/labs/ga/lab2/population"
	"github.com/pityara/labs/ga/lab2/roulette"
)

/* Функция репродукции */
func MakeReproduction(popul []population.Person) []population.Person {
	solutions := mathFn.MakePopulationSolutions(popul)
	populationLength := len(popul)
	/* Инициализация рулетки */
	roulette := roulette.GenerateRoulette(solutions)

	result := make([]population.Person, populationLength)

	/* Подбор особей для репродукции */
	for i := 0; i < populationLength; i++ {
		index := roulette.Spin()
		result[i] = popul[index]
	}
	return result
}
