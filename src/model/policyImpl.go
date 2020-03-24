package model

// NewPolicy :
func NewPolicy() Policy {
	return Policy{
		mapping: make(map[State]map[Action]Quality),
	}
}

// Iterate :
func (policy Policy) Iterate(model func(State, Action) (State, Reward), currentState State, action Action, discount float64) (nextPolicy Policy, nextState State) {
	if _, exist := policy.mapping[currentState]; exist == false {
		// update new state
		policy.mapping[currentState] = make(map[Action]float64)
	}
	if _, exist := policy.mapping[currentState][action]; exist == false {
		// update new action
		policy.mapping[currentState][action] = 0.0
	}
	// transition
	nextState, reward := model(currentState, action)
	// update policy
	_, utility := maxChoose(policy.mapping[nextState])
	policy.mapping[currentState][action] = reward + discount*utility

	// return
	return policy, nextState
}

// OptimalAction :
func (policy Policy) OptimalAction(currentState State) Action {
	action, _ := maxChoose(policy.mapping[currentState])
	return action
}
