// package main

// import "testing"

// func testCalculate(t *testing.T){
// 	if calculate(2) != 4 {
// 		t.Error("Expected 2+2 is to 4")
// 	}

// }

//testing using Testify

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
	assert.Equal(t, Calculate(5), 7)
	assert.Equal(t, Calculate(8), 10)
}
