package lab3

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/pityara/labs/ga/lab3/crossover"
	"github.com/pityara/labs/ga/lab3/mathFn"
	"github.com/pityara/labs/ga/lab3/mutation"
	"github.com/pityara/labs/ga/lab3/population"
	"github.com/pityara/labs/ga/lab3/reproduction"
)

type Out struct {
	result population.Person
}

func Resolve(weights [][]string, quantity uint32) []uint32 {
	startPopulation := population.GeneratePopulation(quantity, uint32(len(weights)))

	mutatedPopulation := mutation.MakeMutation(startPopulation)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	reproducedPopulation := reproduction.MakeReproduction(weights, mutatedPopulation)
	crossedOverPopulation := crossover.MakeCrossover(weights, reproducedPopulation)
	log.Println(crossedOverPopulation)
	index := 0
	solution := uint32(2200)
	iteration := 0
	for solution >= 2200 {
		for i := 0; i < 29*len(reproducedPopulation); i++ {
			a := r.Float64()
			b := r.Float64()
			math.Sqrt((a*a + b*b))
		}
		mutatedPopulation = mutation.MakeMutation(crossedOverPopulation)
		reproducedPopulation = reproduction.MakeReproduction(weights, mutatedPopulation)
		crossedOverPopulation = crossover.MakeCrossover(weights, reproducedPopulation)
		solutions := mathFn.GetSolutions(weights, crossedOverPopulation)
		index = mathFn.GetMinIndex(solutions)
		solution = solutions[index]
		iteration++
		log.Println("iteration:", iteration)
		log.Println("solution:", solution)
	}

	return crossedOverPopulation[index]
}
