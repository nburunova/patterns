package config

import "fmt"

type ConfigObserver struct {
	id int
}

func NewConfigObserver(id int) *ConfigObserver {
	return &ConfigObserver{
		id: id,
	}
}

func (c *ConfigObserver) ID() int {
	return c.id
}

func (c *ConfigObserver) Update(caFile string) {
	fmt.Printf("observer %v received new caFile %v\n", c.id, caFile)
}
