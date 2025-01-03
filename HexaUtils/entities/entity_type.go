package entities

type EntityType int

const (
	Player EntityType = iota
	Cow
	Ghast
	Blaze
	// Add more entity types as needed
)

func (e EntityType) String() string {
	return [...]string{"Player", "Cow", "Ghast", "Blaze"}[e]
}
