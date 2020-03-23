package main

import (
	"fmt"
	"model"
)

var rewardArr = [10]float64{3, 3, 3, 2, 0, 1, 4, 7, 8, 3}

var actionArr = [3]int{-1, 0, +1}

var transition = func(state model.State, action model.Action) (model.State, model.Reward) {
	next := model.State(int(state) + int(action))
	if int(next) < 0 {
		next = 0
		return next, 0
	}
	if int(next) >= len(rewardArr) {
		next = model.State(len(rewardArr) - 1)
		return next, 0
	}
	return next, rewardArr[next] - rewardArr[state]
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

	m := model.NewModel(action, transition)
	p := model.NewPolicy()
	var state model.State = 0
	for i := 0; i < 100000; i++ {
		fmt.Printf("iter: %v, state: %v\n", i, state)
		p = p.Iterate(m, state, 0.99999)
		state = model.State((int(state) + 1) % len(rewardArr))
	}

	state = 0
	fmt.Println("Initial State:", state)
	for i := 0; i < 10; i++ {
		action := p.OptimalAction(state)
		state, _ = transition(state, action)
		fmt.Println("Next State:", state)
	}

}
