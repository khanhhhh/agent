package model

// State :
type State int

// Action :
type Action int

// Reward :
type Reward = float64

// Quality :
type Quality = float64

// Probability :
type Probability = float64

// Policy :
type Policy interface {
	Action(State) map[Action]Quality
	Update(State, Action, Quality)
}

// OptimalAction :
func OptimalAction(policy Policy, state State) Action {
	action, _ := maxChoose(policy.Action(state))
	return action
}
