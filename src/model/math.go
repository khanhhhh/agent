package model

import (
	"math"
	"math/rand"
)

func softmax(mapIn map[Action]Quality) (mapOut map[Action]Probability) {
	mapOut = make(map[Action]Probability)
	var sum float64 = 0
	for key, value := range mapIn {
		expvalue := math.Exp(value)
		sum += expvalue
		mapOut[key] = expvalue
	}
	for key, expvalue := range mapOut {
		mapOut[key] = expvalue / sum
	}
	return mapOut
}

func maxChoose(mapIn map[Action]Quality) (maxAction Action, maxValue Quality) {
	maxValue = 0.0
	for action, value := range mapIn {
		if value > maxValue {
			maxValue = value
			maxAction = action
		}
	}
	return maxAction, maxValue
}
func randomChoose(mapIn map[Action]Probability) (action Action, value Probability) {
	random := rand.Float64()
	cumulative := 0.0
	for key, prob := range mapIn {
		cumulative += prob
		if random < cumulative {
			return key, prob
		}
	}
	panic("randomChoose failed")
}
