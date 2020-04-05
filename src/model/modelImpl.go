package model

// Iterate :
func Iterate(policy Policy, currentState State, action Action, nextState State, reward Reward, discount float64) (nextPolicy Policy) {
	// update policy
	_, nextBestQuality := maxChoose(policy.Action(nextState))
	newQuality := reward + discount*nextBestQuality
	policy.Update(currentState, action, newQuality)

	// return
	return policy
}
