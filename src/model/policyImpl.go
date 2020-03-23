package model

// NewPolicy :
func NewPolicy() Policy {
	return Policy{
		mapping: make(map[State]map[Action]Quality),
	}
}

func match(subAction map[Action]struct{}, allAction map[Action]Quality) map[Action]Quality {
	out := make(map[Action]Quality)
	for action := range subAction {
		out[action] = allAction[action]
	}
	return out
}

// Iterate :
func (policy Policy) Iterate(model func(State, Action) (State, Reward), currentState State, chosenActions map[Action]struct{}, discount float64) (nextPolicy Policy, nextState State) {
	if _, exist := policy.mapping[currentState]; exist == false {
		// update new state
		policy.mapping[currentState] = make(map[Action]float64)
	}
	for action := range chosenActions {
		// update new action
		if _, exist := policy.mapping[currentState][action]; exist == false {
			policy.mapping[currentState][action] = 0.0
		}
	}
	// random choose action
	actionProb := softmax(match(chosenActions, policy.mapping[currentState]))
	action, _ := randomChoose(actionProb)
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
