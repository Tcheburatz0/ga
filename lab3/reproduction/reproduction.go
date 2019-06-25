package reproduction

import (
	"github.com/pityara/labs/ga/lab3/mathFn"
	"github.com/pityara/labs/ga/lab3/population"
	"github.com/pityara/labs/ga/lab3/roulette"
)

/* Функция репродукции */
func MakeReproduction(weights [][]string, popul []population.Person) []population.Person {
	solutions := mathFn.GetSolutions(weights, popul)
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
