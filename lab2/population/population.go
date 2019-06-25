package population

import (
	"math/rand"
	"time"
)

type Person []float64

func GeneratePopulation(personsNumber int, restrictions []float64, measures int) []Person {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	result := make([]Person, personsNumber)
	min := restrictions[0]
	max := restrictions[1]

	for i := 0; i < personsNumber; i++ {
		person := make(Person, measures)
		for j := 0; j < measures; j++ {
			person[j] = min + r.Float64()*(max-min)
		}
		result[i] = person
	}

	return result
}
