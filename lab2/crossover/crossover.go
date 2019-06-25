package crossover

import (
	"math/rand"
	"time"

	"github.com/pityara/labs/ga/lab2/population"
)

func MakeCrossover(popul []population.Person) []population.Person {
	chance := .5                               // Шанс кроссовера 50%
	w := .370                                  // Коэффицент
	s := rand.NewSource(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	r := rand.New(s)
	populationLen := len(popul)
	result := make([]population.Person, populationLen)
	for i := 0; i < populationLen; i += 2 {
		/* Выбор случайных отца и матери */
		dad := popul[r.Intn(populationLen)]
		mom := popul[r.Intn(populationLen)]
		if r.Float64() < chance {
			result[i] = dad
			if populationLen == i+1 {
				continue
			}
			result[i+1] = mom
			continue
		}

		childs := getChilds(mom, dad, w)

		result[i] = childs[0]

		if populationLen == i+1 {
			continue
		}
		result[i+1] = childs[1]
	}

	return result
}

func getChilds(mom population.Person, dad population.Person, w float64) [2]population.Person {
	firstChild := make([]float64, len(mom))
	secondChild := make([]float64, len(mom))
	for key, val := range mom {
		firstChild[key] = val*w + dad[key]*(1-w)
		secondChild[key] = dad[key]*w + val*(1-w)
	}
	return [2]population.Person{firstChild, secondChild}
}
