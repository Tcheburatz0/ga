package crossover

import (
	"math/rand"
	"strings"
	"time"
)

func MakeCrossover(population []string) []string {
	chance := .5                               // Шанс кроссовера 50%
	s := rand.NewSource(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	r := rand.New(s)
	personLen := len(population[0])
	populationLen := len(population)
	result := make([]string, populationLen)
	for i := 0; i < populationLen; i += 2 {
		/* Выбор случайных отца и матери */
		dad := population[r.Intn(populationLen)]
		mom := population[r.Intn(populationLen)]
		if r.Float64() < chance {
			result[i] = dad
			if populationLen == i+1 {
				continue
			}
			result[i+1] = mom
			continue
		}
		runeDad := []rune(dad)
		runeMom := []rune(mom)
		position := r.Intn(personLen - 1)
		/* Операция кроссовера на случайной позиции */
		result[i] = strings.Join([]string{string(runeMom[:position]), string(runeDad[position:])}, "")

		if populationLen == i+1 {
			continue
		}
		result[i+1] = strings.Join([]string{string(runeDad[:position]), string(runeMom[position:])}, "")
	}
	return result
}
