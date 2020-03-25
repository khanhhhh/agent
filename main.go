package main

import (
	"fmt"
	"math/rand"
	"model"
)

var valueArr = [10]float64{3, 3, 3, 2, 0, 1, 4, 7, 8, 3}

var actionArr = []int{-2, -1, 0, +1, +2}

var transition = func(current model.State, action model.Action) (model.State, model.Reward) {
	next := model.State(int(current) + int(action))
	if int(next) < 0 {
		next = 0
	}
	if int(next) >= len(valueArr) {
		next = model.State(len(valueArr) - 1)
	}
	return next, valueArr[next] - valueArr[current]
}

func main() {
	fmt.Println("Hello, World!")

	policy := model.NewDiscretePolicy()
	var currentState model.State = 0
	for i := 0; i < 1000; i++ {
		action := model.Action(actionArr[rand.Intn(len(actionArr))])
		nextState, reward := transition(currentState, action)
		policy = model.Iterate(policy, currentState, action, nextState, reward, 0.99)
		currentState = nextState
	}

	// Optimal Policy
	var state model.State = 0
	fmt.Println("Initial State:", state)
	for i := 0; i < 10; i++ {
		action := model.OptimalAction(policy, state)
		state, _ = transition(state, action)
		fmt.Println("Next State:", state)
	}

}
