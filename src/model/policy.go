package model

// Quality :
type Quality = float64

// Policy :
type Policy struct {
	mapping map[State]map[Action]Quality
}
