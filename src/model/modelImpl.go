package model

// Iterate :
func Iterate(policy Policy, currentState State, action Action, nextState State, reward Reward, discount float64) (nextPolicy Policy) {
	// update policy
	_, utility := maxChoose(policy.Action(nextState))
	newQuality := reward + discount*utility
	policy.Update(currentState, action, newQuality)

	// return
	return policy
}
