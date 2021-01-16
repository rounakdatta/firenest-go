package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractClass(t *testing.T) {
	hdfcAccount := HDFCAccount{}
	Process(&hdfcAccount)

	axisAccount := AxisAccount{}
	Process(&axisAccount)

	assert.Equal(t, hdfcAccount.Name, "HDFC")
	assert.Equal(t, axisAccount.Name, "Axis")
}
