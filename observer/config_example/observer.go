package config_example

type Observer interface {
	ID() int
	Update(caFile string)
}