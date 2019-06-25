package helpers

import (
	"log"

	"github.com/pityara/labs/ga/lab3/population"
)

func Check(person population.Person) bool {
	current := person[0]
	for person[current] != 0 {
		current = person[current]
		log.Println(1)
	}

	return true
}
