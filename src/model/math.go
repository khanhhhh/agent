package model

import (
	"math/rand"
)

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
