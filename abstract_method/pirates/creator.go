package pirates

import (
	"fmt"
	"math/rand"
	"patterns/abstract_method"
)

const (
	PirateType string = "Pirate"
)

type Pirate struct {
	id     int
	health int
}

func (p *Pirate) Name() string {
	return fmt.Sprintf("%v %v", PirateType, p.id)
}

func (p *Pirate) Type() string {
	return PirateType
}

func (p *Pirate) Health() int {
	return p.health
}

type PiratesGang struct {} 

func (pg *PiratesGang) CreateSquad(size, averageHealth int) []squad.Member {
	maxDiff := averageHealth / 3
	members := make([]squad.Member, 0, size)
	var diff, i int
	for {
		if i >= size {
			break
		}
		diff = rand.Intn(maxDiff)
		members = append(members,
			&Pirate{
				id:     i + 1,
				health: averageHealth - diff,
			},
			&Pirate{
				id:     i + 2,
				health: averageHealth + diff,
			})
		i += 2
	}
	return members
}
