package roulette

import (
	"math/rand"
	"time"

	"github.com/pityara/labs/ga/lab2/mathFn"
)

type Sector struct {
	index int
	start float64
	end   float64
}

type Roulete []Sector

/* Сумма всех значений массива */
func getSum(a []mathFn.Solution) float64 {
	sum := .0
	for _, val := range a {
		sum += 1 / val.Y
	}

	return sum
}

/* Генерация рулетки */
func GenerateRoulette(solutions []mathFn.Solution) Roulete {
	result := make([]Sector, len(solutions))
	sum := getSum(solutions)
	start := .0
	for index, val := range solutions {
		l := (1 / val.Y) / sum
		result[index] = Sector{index, start, start + l}
		start += l
	}

	return result
}

/* Вращение рулетки */
func (roulette Roulete) Spin() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	rand := r.Float64()

	for index, val := range roulette {
		if rand > val.start && rand <= val.end {
			return index
		}
	}

	return -1
}
