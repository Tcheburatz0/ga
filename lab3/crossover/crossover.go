package crossover

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/pityara/labs/ga/lab3/population"
)

type Crossover struct {
	currentMomPosition   int
	currentDadPosition   int
	currentChildPosition int
	dad                  population.Person
	mom                  population.Person
	child                population.Person
}

func MakeCrossover(weights [][]string, popul []population.Person) []population.Person {
	chance := .5
	s := rand.NewSource(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	r := rand.New(s)
	result := make([]population.Person, len(popul))
	var wg sync.WaitGroup

	for i := 0; i < len(popul); i += 2 {
		if r.Float64() > chance {
			wg.Add(1)
			go makeCrossover(popul[i], popul[i+1], weights, &result[i], &result[i+1], &wg)
			continue

		}
		result[i] = popul[i]
		result[i+1] = popul[i+1]
	}

	wg.Wait()

	return result
}

func makeCrossover(mom population.Person, dad population.Person, weights [][]string, firstChild *population.Person, secondChild *population.Person, wg *sync.WaitGroup) {
	*firstChild = makeChild(mom, dad, weights)
	*secondChild = makeChild(dad, mom, weights)

	wg.Done()
}

func makeChild(mom population.Person, dad population.Person, weights [][]string) population.Person {
	crossover := Crossover{0, 0, 0, dad, mom, generateEmptyChild(len(dad))}
	num := 0
	for num < len(dad)-1 {
		crossover.setNextPosition(weights)
		num++
	}

	crossover.child[crossover.currentChildPosition] = 0

	return crossover.child
}

func (c *Crossover) setNextPosition(weights [][]string) {
	momPosition := c.getPositionFromMom()
	dadPosition := c.getPositionFromDad()
	momWeight, _ := strconv.ParseInt(weights[c.currentChildPosition][momPosition], 10, 64)
	dadWeight, _ := strconv.ParseInt(weights[c.currentChildPosition][dadPosition], 10, 64)
	position := uint32(0)
	if momWeight < dadWeight {
		position = momPosition
	} else {
		position = dadPosition
	}
	c.child[c.currentChildPosition] = position
	c.currentChildPosition = int(position)
}

func (c *Crossover) getPositionFromDad() uint32 {
	pos := c.currentChildPosition
	position := c.dad[pos]
	for c.child[position] != uint32(len(c.child)+2) || c.child[position] == 0 || c.currentChildPosition == int(position) {
		if pos == (len(c.dad) - 1) {
			pos = 0
		} else {
			pos++
		}
		position = uint32(pos)
	}

	c.currentDadPosition = pos

	return position
}

func (c *Crossover) getPositionFromMom() uint32 {
	pos := c.currentChildPosition
	position := c.mom[pos]
	for c.child[position] != uint32(len(c.child)+2) || c.child[position] == 0 || c.currentChildPosition == int(position) {
		if pos == (len(c.mom) - 1) {
			pos = 0
		} else {
			pos++
		}
		position = uint32(pos)
	}

	c.currentMomPosition = pos

	return position
}

func generateEmptyChild(length int) population.Person {
	result := make(population.Person, length)

	for key := range result {
		result[key] = uint32(length + 2)
	}

	return result
}
