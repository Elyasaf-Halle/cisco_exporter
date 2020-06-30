package environment

type EnvironmentItem struct {
	Name         string
	Status       string
	OK           bool
	IsPowerUsage bool
	IsTemp       bool
	Value        float64
}
