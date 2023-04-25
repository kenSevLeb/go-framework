package number

import (
	"math"
	"math/rand"
	"time"
)

// return min number
func Min(a, b int) int {
	if a < b {
		return a
	}
	if b < a {
		return b
	}
	return a
}

// return max number
func Max(a, b int) int {
	if a > b {
		return a
	}
	if b > a {
		return b
	}
	return a
}

// return a/b
func Div(a, b int) int {
	if b == 0 || a == 0 {
		return 0
	}

	return a / b
}

// 向上取整
func Ceil(a, b int) int {
	if b == 0 || a == 0 {
		return 0
	}

	return int(math.Ceil(float64(a) / float64(b)))
}

// rounding-off method
func Round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// 取随机数
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// DefaultInt
func DefaultInt(val, defaultVal int) int {
	if val == 0 {
		return defaultVal
	}

	return val
}

// Decimal 浮点数保留n位小数
func Decimal(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc(f*n10) / n10
}
