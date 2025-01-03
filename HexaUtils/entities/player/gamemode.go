package player

type GameMode int

const (
	Survival GameMode = iota
	Creative
	Adventure
	Spectator
	Hardcore
)
