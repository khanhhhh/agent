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
type Policy struct {
	mapping map[State]map[Action]Quality
}
