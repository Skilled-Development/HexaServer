package player

// Definir un "enum" para los estados del jugador
type ClientState int

const (
	Handshake     ClientState = iota // 0
	Status                           // 1
	Login                            // 2
	Configuration                    // 3
	Play                             // 4
)

// Función para imprimir el estado como cadena
func (s ClientState) String() string {
	switch s {
	case Handshake:
		return "Handshake"
	case Status:
		return "Status"
	case Login:
		return "Login"
	case Configuration:
		return "Configuration"
	case Play:
		return "Play"
	default:
		return "Unknown"
	}
}

func (s ClientState) Value() int {
	return int(s)
}
