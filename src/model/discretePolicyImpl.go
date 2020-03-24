package model

type discretePolicy struct {
	mapping map[State]map[Action]Quality
}

// NewDiscretePolicy :
func NewDiscretePolicy() Policy {
	return discretePolicy{
		mapping: make(map[State]map[Action]Quality),
	}
}

// Action :
func (policy discretePolicy) Action(state State) map[Action]Quality {
	if _, exist := policy.mapping[state]; exist == false {
		// update new state
		policy.mapping[state] = make(map[Action]float64)
	}
	return policy.mapping[state]
}

// Update :
func (policy discretePolicy) Update(state State, action Action, quality Quality) {
	if _, exist := policy.mapping[state]; exist == false {
		// update new state
		policy.mapping[state] = make(map[Action]float64)
	}
	if _, exist := policy.mapping[state][action]; exist == false {
		// update new action
		policy.mapping[state][action] = 0.0
	}
	policy.mapping[state][action] = quality
}
