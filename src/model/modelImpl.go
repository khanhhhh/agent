package model

// NewModel :
func NewModel(action func(State) map[Action]struct{}, transition func(State, Action) (State, Reward)) Model {
	return Model{
		action:     action,
		transition: transition,
	}
}
