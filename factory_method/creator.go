package squad

type Member interface {
	Type() string
	Name() string
	Health() int
}

type Creator interface {
	CreateSquad(size, averageHealth int) []Member
}