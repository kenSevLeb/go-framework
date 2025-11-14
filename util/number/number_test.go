package number

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		if n := RandInt(1, 3); n != 1 {
			fmt.Println(n)
		}
	}
}

func TestDiv(t *testing.T) {
	assert.Equal(t, 1, Div(1, 1))
	assert.Equal(t, 0, Div(1, 2))
	assert.Equal(t, 1, Div(3, 2))
}

func TestCeil(t *testing.T) {
	assert.Equal(t, 1, Ceil(2, 2))
	assert.Equal(t, 1, Ceil(3, 4))
	assert.Equal(t, 2, Ceil(5, 4))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, Max(2, 1))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 3))
}

func TestRound(t *testing.T) {
	fmt.Println(Round(1.6))
}

func TestDefaultInt(t *testing.T) {
	assert.Equal(t, 1, DefaultInt(0, 1))
}

func TestDecimal(t *testing.T) {
	f := 10.123456789
	assert.Equal(t, Decimal(f, 0), 10.0)
	assert.Equal(t, Decimal(f, 1), 10.1)
	assert.Equal(t, Decimal(f, 2), 10.12)
	assert.Equal(t, Decimal(f, 3), 10.123)
	assert.Equal(t, Decimal(f, 4), 10.1234)
	assert.Equal(t, Decimal(f, 5), 10.12345)
	assert.Equal(t, Decimal(f, 6), 10.123456)
	assert.Equal(t, Decimal(f, 7), 10.1234567)
	assert.Equal(t, Decimal(f, 8), 10.12345678)
}
