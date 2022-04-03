package knights

import (
	"fmt"
	"math/rand"
	"patterns/abstract_method"
)

const (
	KnightType string = "Knight"
)

type Knight struct {
	id     int
	health int
}

func (p *Knight) Name() string {
	return fmt.Sprintf("%v %v", KnightType, p.id)
}

func (p *Knight) Type() string {
	return KnightType
}

func (p *Knight) Health() int {
	return p.health
}

type KnightsPlatoon struct {}

func (kp *KnightsPlatoon) CreateSquad(size, averageHealth int) []squad.Member {
	maxDiff := averageHealth / 3
	members := make([]squad.Member, 0, size)
	var diff, i int
	for {
		if i >= size {
			break
		}
		diff = rand.Intn(maxDiff)
		members = append(members,
			&Knight{
				id:     i + 1,
				health: averageHealth - diff,
			},
			&Knight{
				id:     i + 2,
				health: averageHealth + diff,
			})
		i += 2
	}
	return members
}