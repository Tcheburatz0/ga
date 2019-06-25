package mutation

import (
	"math/rand"
	"time"

	"github.com/pityara/labs/ga/lab2/population"
)

func MakeMutation(population []population.Person, restrictions []float64) []population.Person {
	/* Шанc мутации 0.1% */
	chance := 0.001

	/* Инициализация генератора случайных чисел */
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	personLength := len(population[0])

	min := restrictions[0]
	max := restrictions[1]

	for _, val := range population {
		if r.Float64() <= chance {
			position := r.Intn(personLength - 1)
			val[position] = min + rand.Float64()*(max-min)
		}
	}

	return population
}
