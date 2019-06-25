package lab2

import (
	"log"

	"E:/8_семестр/Эволюционные_методы_моделирования/lab2lab2/crossover"
	"github.com/pityara/labs/ga/lab2/mathFn"
	"github.com/pityara/labs/ga/lab2/mutation"
	"github.com/pityara/labs/ga/lab2/population"
	"github.com/pityara/labs/ga/lab2/reproduction"
)

//f2(x)=sum(100·(x(i+1)-x(i)^2)^2+(1-x(i))^2), i=1:n-1; x при [-2.048, 2.048]/
//

/* Структура для входных данных */
type In struct {
	Floor      float64 `json:"floor"`      // нижняя граница
	Ceil       float64 `json:"ceil"`       // верхняя граница
	Measures   int     `json:"measures"`   // количество измерений
	Quatity    int     `json:"quantity"`   // количество особей
	Iterations int     `json:"iterations"` // количество поколений
}

func Resolve(in *In) []mathFn.Solution {
	restrictions := []float64{in.Floor, in.Ceil}
	popul := population.GeneratePopulation(in.Quatity, restrictions, in.Measures)
	reproducedPopulation := make([]population.Person, in.Quatity)
	crossoverPopulation := make([]population.Person, in.Quatity)
	mutatedPopulation := mutation.MakeMutation(popul, restrictions)
	for i := 0; i < in.Iterations; i++ {
		reproducedPopulation = reproduction.MakeReproduction(mutatedPopulation)
		crossoverPopulation = crossover.MakeCrossover(reproducedPopulation)
		mutatedPopulation = mutation.MakeMutation(crossoverPopulation, restrictions)
	}

	solutions := mathFn.MakePopulationSolutions(mutatedPopulation)

	return solutions
}

func main() {
	startPopulation := population.GeneratePopulation(10, []float64{-2.385, 2.345}, 2)

	first := reproduction.MakeReproduction(startPopulation)
	log.Print(crossover.MakeCrossover(first))
}
