package mutation

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeMutation(population []string) []string {
	/* Шан мутации 0.1% */
	chance := 0.001

	/* Инициализация генератора случайных чисел */
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	personLength := len(population[0])

	for index, val := range population {
		go func() {
			if r.Float64() <= chance {
				position := r.Intn(personLength - 1)
				person := []rune(val)
				/* Инверсия случайного бита */
				person[position] = inv(person[position])
				fmt.Println("Make mutation for person ", val, " on index ", position)
				population[index] = string(person)
			}
		}()
	}

	return population
}

/* Инверсия */
func inv(val rune) rune {
	if val == '0' {
		return '1'
	}

	return '0'
}
