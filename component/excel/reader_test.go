package excel

import (
	"fmt"
	"testing"
)

func TestReader(t *testing.T) {
	r := Reader(filename, 10)
	go r.Read(0)

	for i := range r.OutPut() {
		fmt.Println(i)
	}
}
