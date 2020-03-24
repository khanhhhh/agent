package model

// Iterate :
func Iterate(policy Policy, model func(State, Action) (State, Reward), currentState State, action Action, discount float64) (nextPolicy Policy, nextState State) {
	// transition
	nextState, reward := model(currentState, action)
	// update policy
	_, utility := maxChoose(policy.Action(nextState))
	newQuality := reward + discount*utility
	policy.Update(currentState, action, newQuality)

	// return
	return policy, nextState
}
