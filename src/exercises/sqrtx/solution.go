package solution

import (
	"math"
)

func mySqrt(s int) int {
	if s == 0 {
		return 0
	}

	S := float64(s)
	if S < 0 {
		return 0 // Возвращаем NaN для отрицательных чисел
	}
	epsilon := 0.1

	x := S                          // Начальное приближение
	for math.Abs(x*x-S) > epsilon { // Пока разница больше заданной точности
		x = 0.5 * (x + S/x) // Формула Ньютона
	}
	return int(x / 1.0)
}
