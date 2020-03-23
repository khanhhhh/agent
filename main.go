package main

import (
	"fmt"
	"model"
)

var valueArr = [10]float64{3, 3, 3, 2, 0, 1, 4, 7, 8, 3}

var actionArr = [3]int{-1, 0, +1}

var transition = func(current model.State, action model.Action) (model.State, model.Reward) {
	next := model.State(int(current) + int(action))
	if int(next) < 0 {
		next = 0
		return next, 0
	}
	if int(next) >= len(valueArr) {
		next = model.State(len(valueArr) - 1)
		return next, 0
	}
	return next, valueArr[next] - valueArr[current]
}

var action = func(model.State) map[model.Action]struct{} {
	return map[model.Action]struct{}{
		-1: struct{}{},
		0:  struct{}{},
		+1: struct{}{},
	}
}

func main() {
	fmt.Println("Hello, World!")

	p := model.NewPolicy()
	var currentState model.State = 0
	var nextState model.State
	for i := 0; i < 100000; i++ {
		fmt.Printf("iter: %v, state: %v\n", i, currentState)
		p, nextState = p.Iterate(transition, currentState, action(currentState), 0.99)
		currentState = nextState
	}

	// Optimal Policy
	var state model.State = 0
	fmt.Println("Initial State:", state)
	for i := 0; i < 10; i++ {
		action := p.OptimalAction(state)
		state, _ = transition(state, action)
		fmt.Println("Next State:", state)
	}

}
