package population

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

/* Генерация особи */
func makePerson(pow int8) string {
	/* Инициализация генератора случайных чисел */
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	rand := r.Int63n(int64(math.Pow(2, float64(pow))))
	return strconv.FormatUint(uint64(rand), 2)
}

/* Генерация популяции */
func GeneratePopulation(quantity int, pow int8) []string {
	result := make([]string, quantity)

	for i := 0; i < quantity; i++ {
		result[i] = makePerson(pow)
	}

	return result
}
