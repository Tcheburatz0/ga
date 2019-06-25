package population

import (
	"math/rand"
	"time"
)

type Person []uint32

func GeneratePopulation(personsNumber uint32, verticesNumber uint32) []Person {
	result := make([]Person, personsNumber)
	out := make(chan Person)
	for i := uint32(0); i < personsNumber; i++ {
		go generatePerson(out, verticesNumber)
		result[i] = <-out
	}
	return result
}

func generatePerson(out chan<- Person, verticesNumber uint32) {
	person := initArray(verticesNumber)
	currentPosition := uint32(0)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := uint32(0); i < verticesNumber-1; i++ {
		val := nextValue(person, verticesNumber, r, currentPosition)
		person[currentPosition] = val

		currentPosition = val
	}
	person[currentPosition] = 0

	out <- person
}

func nextValue(person Person, verticesNumber uint32, r *rand.Rand, prev uint32) uint32 {
	next := uint32(r.Int63n(int64(verticesNumber)))
	for person[next] != verticesNumber+2 || next == person[next] || prev == next || next == 0 {
		next = uint32(r.Int63n(int64(verticesNumber)))
	}

	return uint32(next)
}

func initArray(length uint32) Person {
	result := make(Person, length)

	for key := range result {
		result[key] = length + 2
	}

	return result
}
