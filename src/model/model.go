package model

// State :
type State int

// Action :
type Action int

// Reward :
type Reward = float64

// Model :
type Model struct {
	action     func(State) map[Action]struct{}     // stocahstic
	transition func(State, Action) (State, Reward) // stochastic
}
