package model

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	mapIn := map[Action]float64{
		0: 0,
		1: 1,
		2: 2,
	}
	prob := softmax(mapIn)
	fmt.Println(prob)
	fmt.Println(maxChoose(mapIn))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
	fmt.Println(randomChoose(prob))
}
