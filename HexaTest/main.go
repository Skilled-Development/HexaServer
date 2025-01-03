package main

import (
	"HexaServer/server"
	"HexaUtils/server/config"
	"fmt"
)

func main() {
	motd := config.NewMOTD()
	server_config := config.NewServerConfig(*motd, 10)
	myServer := server.NewServer(server_config)
	fmt.Println("Iniciando el servidor...")
	go myServer.Start()

	// Mantener el programa en ejecución
	select {}
}
