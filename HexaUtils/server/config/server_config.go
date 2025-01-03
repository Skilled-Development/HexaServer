package config

import (
	"HexaUtils/entities/player"
	"HexaUtils/server/data"
)

var ServerConfigInstance *ServerConfig

type ServerConfig struct {
	MOTD               MOTD
	MaxPlayers         int
	SimulationDistance int
	ViewDistance       int
	DefaultGamemode    player.GameMode
	InvalidTpIdMessage string
}

func NewServerConfig(motd MOTD, viewDistance int) *ServerConfig {
	config := &ServerConfig{
		MOTD:               motd,
		MaxPlayers:         1000,
		ViewDistance:       viewDistance,
		SimulationDistance: 10,
		DefaultGamemode:    player.Creative,
		InvalidTpIdMessage: "&cInvalid teleport ID",
	}
	ServerConfigInstance = config
	data.NewServerData()
	return config
}

func (s *ServerConfig) GetMaxPlayers() int {
	return s.MaxPlayers
}

func (s *ServerConfig) GetViewDistance() int {
	return s.ViewDistance
}

func (s *ServerConfig) GetMOTD() MOTD {
	return s.MOTD
}

func (s *ServerConfig) GetSimulationDistance() int {
	return s.SimulationDistance
}

func (s *ServerConfig) GetInvalidTpIdMessage() string {
	return s.InvalidTpIdMessage
}

func (s *ServerConfig) GetDefaultGamemode() player.GameMode {
	return s.DefaultGamemode
}
