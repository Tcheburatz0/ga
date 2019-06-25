package slicing

import (
	"fmt"
	"math"
	"strconv"
)

type Restr struct {
	Floor     float64 `json:"floor"`
	Ceil      float64 `json:"ceil"`
	Precision int     `json:"precision"`
	Pow       int8    `json:"pow"`
}

/* Вычисление диапозона целых значений для представления промежутка */
func calcRange(floor float64, ceil float64, precision int8) float64 {
	multiplier := math.Pow10(int(precision))
	return math.Abs(ceil-floor) * multiplier
}

/* Вычисление степени 2 */
func CalcPow(floor float64, ceil float64, precision int8) int8 {
	cRange := calcRange(floor, ceil, precision)

	pow, val := .0, .0
	for val <= cRange {
		val = math.Pow(2, pow)
		pow++
	}

	return int8(pow) - 1
}

/* Конвертация бинарного значения в float */
func (r Restr) BinaryToValue(binary string) (float64, error) {

	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return .0, err
	}
	result := r.Floor + float64(i)*((r.Ceil-r.Floor)/(math.Pow(2, float64(r.Pow))-1))
	return strconv.ParseFloat(fmt.Sprintf("%.3f", result), 3)
}
