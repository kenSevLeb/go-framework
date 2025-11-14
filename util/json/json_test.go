package json

import (
	"fmt"
	"github.com/json-iterator/go/extra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeForceFromString(t *testing.T) {
	type Product struct {
		Name   string
		Price  float64 `json:"price"`
		Number int     `json:"number"`
	}
	s := `{"name":"Galaxy Nexus", "price":"3460.00", "number":"1"}`
	var pro Product
	extra.RegisterFuzzyDecoders()
	err := DecodeForceFromString(s, &pro)
	assert.Nil(t, err)
	fmt.Println(pro)

	arr := `[{"name":"Galaxy Nexus", "price":"3460.00", "number":"2"},{"name":"Galaxy Nexus", "price":"3460.00", "number":"1"}]`
	var items []Product
	assert.Nil(t, DecodeForce([]byte(arr), &items))
	fmt.Println(items)
}

func TestFastEncode(t *testing.T) {
	type Product struct {
		Name   string
		Price  float64 `json:"price"`
		Number int     `json:"number"`
	}
	p := Product{
		Name:   "Galaxy Nexus",
		Price:  3460.00,
		Number: 1,
	}

	s, err := Encode(p)
	assert.Nil(t, err)
	var r Product
	assert.Nil(t, Decode([]byte(s), &r))
	assert.Equal(t, p.Name, r.Name)
	assert.Equal(t, p.Price, r.Price)
	assert.Equal(t, p.Number, r.Number)
}

func TestEncode(t *testing.T) {
	items := map[string]interface{}{"name": "test"}
	str, err := Encode(items)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"test"}`, str)
}

func TestMustEncode(t *testing.T) {
	items := map[string]interface{}{"name": "test"}
	str := MustEncode(items)
	assert.Equal(t, `{"name":"test"}`, str)
}

func TestDecode(t *testing.T) {
	str := `{"name":"test"}`
	var items map[string]interface{}
	err := Decode([]byte(str), &items)
	assert.Nil(t, err)
	assert.Equal(t, "test", items["name"])
}
