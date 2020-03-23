package model

// NewPolicy :
func NewPolicy() Policy {
	return Policy{
		mapping: make(map[State]map[Action]Quality),
	}
}

// Iterate :
func (policy Policy) Iterate(model Model, currentState State, discount float64) Policy {
	if _, exist := policy.mapping[currentState]; exist == false {
		// state not exist -> initialize
		policy.mapping[currentState] = make(map[Action]float64)
		allActions := model.action(currentState)
		for action := range allActions {
			policy.mapping[currentState][action] = 1.0
		}
	}
	// random choose action
	actionProb := softmax(policy.mapping[currentState])
	action, _ := randomChoose(actionProb)
	// transition
	nextState, reward := model.transition(currentState, action)
	// update policy
	_, utility := maxChoose(policy.mapping[nextState])
	policy.mapping[currentState][action] = reward + discount*utility

	// return
	return policy
}

// OptimalAction :
func (policy Policy) OptimalAction(currentState State) Action {
	action, _ := maxChoose(policy.mapping[currentState])
	return action
}
