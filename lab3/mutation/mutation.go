package mutation

import (
	"math/rand"
	"time"

	"github.com/pityara/labs/ga/lab3/population"
)

func MakeMutation(popul []population.Person) []population.Person {
	/* Шанc мутации 0.1% */
	chance := .1

	/* Инициализация генератора случайных чисел */
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	personLength := len(popul[0])

	for key, val := range popul {
		if r.Float64() <= chance {
			position := r.Intn(personLength - 1)
			ch := make(chan population.Person)
			go makeMutationForPersonFromPosition(val, uint32(position), ch)
			popul[key] = <-ch
		}
	}

	return popul
}

func makeMutationForPersonFromPosition(person population.Person, position uint32, out chan population.Person) {
	initialSlice := generateEmptyChild(len(person))
	for i := uint32(0); i != position; i = person[i] {
		initialSlice[i] = person[i]
	}
	verticesNumber := uint32(len(person))

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	currentPosition := position
	for hasEmptyValues(initialSlice) {
		next := nextValue(initialSlice, verticesNumber, r, currentPosition)
		initialSlice[currentPosition] = next
		currentPosition = next
	}
	initialSlice[currentPosition] = 0

	out <- initialSlice
}

func nextValue(person population.Person, verticesNumber uint32, r *rand.Rand, prev uint32) uint32 {
	next := uint32(r.Int63n(int64(verticesNumber)))
	for person[next] != verticesNumber+2 || prev == next {
		next = uint32(r.Int63n(int64(verticesNumber)))
	}

	return uint32(next)
}

func generateEmptyChild(length int) population.Person {
	result := make(population.Person, length)

	for key := range result {
		result[key] = uint32(length + 2)
	}

	return result
}

func hasEmptyValues(person population.Person) bool {
	emptyValue := uint32(len(person) + 2)
	result := 0
	for _, val := range person {
		if val == emptyValue {
			result++
		}
	}

	return result > 1
}
