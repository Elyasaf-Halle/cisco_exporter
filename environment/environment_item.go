package environment

type EnvironmentItem struct {
	Name         string
	Status       string
	OK           bool
	IsPowerUsage bool
	Power        float64
	IsTemp       bool
	Temperature  float64
}
