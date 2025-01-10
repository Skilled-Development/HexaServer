// server.go
package server

import (
	tick_1_21 "HexaProtocol_1_21/entities/tick"
	"HexaProtocol_1_21/packets"
	hexaProtocol_1_21 "HexaProtocol_1_21/protocol"
	entities_manager "HexaServer/entities/manager"
	"HexaServer/entities/player"
	"HexaServer/packet"
	hexapackets "HexaUtils/packets/utils"
	region "HexaUtils/regionreader"
	"HexaUtils/server/config"
	data "HexaUtils/server/data"
	"HexaUtils/utils"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" // Importa el paquete pprof
	"runtime"
	"sync"
	"time"

	"HexaProtocol_1_21/packets/clientbound"

	"github.com/shirou/gopsutil/cpu"
)

// Estructura para el servidor
type Server struct {
	listener         net.Listener
	clients          map[string]net.Conn
	EntitiesManager  *entities_manager.EntityManager
	RegistriesManger *config.RegistriesManager
	packet_reader    *packet.PlayerPacketReader
	server_config    *config.ServerConfig
	tickRate         time.Duration
}

func NewServer(motd *config.ServerConfig) *Server {
	return &Server{
		clients:       make(map[string]net.Conn),
		server_config: motd,
		tickRate:      time.Second / 20, // 20 ticks por segundo
	}
}

// Inicia el servidor en el puerto 25565
func (s *Server) Start() {
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)
	runtime.ReadMemStats(&m)
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)
	utils.NewDebugger()
	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)
	utils.PrintLog("Starting server...")

	startTime := utils.GetTimeMillis()
	hexaProtocol_1_21.NewHexaProtocol_1_21()
	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)
	utils.SetDebugTest(false)
	regionFile, err := region.OpenRegion("worlds/overworld/region/r.0.0.mca")
	if err != nil {
		log.Fatalf("Error opening region file: %v", err)
	}
	defer regionFile.Close()
	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)

	data.AddRegion(*regionFile)
	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)

	fmt.Println("Region file processed.")

	s.RegistriesManger = config.NewRegistriesManager()
	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)
	s.EntitiesManager = entities_manager.NewEntityManager()
	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)
	s.packet_reader = packet.NewPlayerPacketReader()
	runtime.ReadMemStats(&m)
	fmt.Printf("Objetos en el heap: %v\n", m.HeapObjects)

	var err2 error
	s.listener, err2 = net.Listen("tcp", ":25565")
	if err2 != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err2)
	}
	defer s.listener.Close()

	s.clients = make(map[string]net.Conn)

	fmt.Println("Servidor escuchando en el puerto 25565...")
	fmt.Printf("Tiempo de inicio: %d ms\n", utils.GetTimeMillis()-startTime)

	// Inicia el goroutine del tick del servidor
	go s.runTick()

	for {
		conn, err2 := s.listener.Accept()
		if err2 != nil {
			log.Printf("Error al aceptar conexión: %v", err2)
			continue
		}

		// Agregar cliente
		s.clients[conn.RemoteAddr().String()] = conn
		println("Nuevo cliente conectado: %s\n", conn.RemoteAddr())

		p := s.EntitiesManager.CreatePlayer(conn)
		player, ok := p.(*player.Player)
		if !ok {
			log.Printf("Error al convertir el jugador: %v\n", conn.RemoteAddr())
			conn.Close()
			continue
		}
		go s.handleClient(player)
	}
}

func (s *Server) handleClient(Player *player.Player) {
	conn := *Player.GetConn()
	defer conn.Close()
	buffer := make([]byte, 2048)
	packetReader := hexapackets.NewPacketReader(nil) // Crea una unica instancia aquí

	for {
		// Lee los paquetes del cliente
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error al leer del cliente %s: %v\n", conn.RemoteAddr(), err)
			conn.Close()
			s.EntitiesManager.RemovePlayer(Player)
			return
		}
		packetReader.SetBuffer(buffer[:n]) // Set buffer del packetReader
		s.packet_reader.ReadPlayerPacket(packetReader, Player, s.server_config)
		//clear buffer
		buffer = make([]byte, 2048)
	}
}

var m runtime.MemStats
var maxRamUsage float64
var minRamUsage float64
var maxHeapObjects uint64
var minHeapObjects uint64
var maxCPUUsage float64

func (s *Server) runTick() {
	minRamUsage = 10000000
	minHeapObjects = 10000000
	ticker := time.NewTicker(s.tickRate)
	defer ticker.Stop()

	tickCount := 0
	var msptList []float64

	for range ticker.C {
		tickStart := time.Now()
		s.tick(tickCount)
		elapsedTime := time.Since(tickStart)
		mspt := float64(elapsedTime.Microseconds()) / 1000.0000000
		msptList = append(msptList, mspt)
		tickCount++

		if tickCount == 20 {
			//runtime.GC()
			tickCount = 0
			/*var sum float64
			for _, val := range msptList {
				sum += val
			}
			medianMspt := sum / float64(len(msptList))
			tps := 1000.0000000 / medianMspt

			//log.Printf("MSPT: %.4fms TPS: %.2f\n", medianMspt, tps)
			msptList = msptList[:0] // Limpiamos la lista
			runtime.ReadMemStats(&m)

			heapAllocMB := bytesToMB(m.HeapAlloc)
			stackInUseMB := bytesToMB(m.StackInuse)
			sysMB := bytesToMB(m.Sys)
			cpuUsage := GetCPUUsage()

			if heapAllocMB > maxRamUsage {
				maxRamUsage = heapAllocMB
			}

			if heapAllocMB < minRamUsage {
				minRamUsage = heapAllocMB
			}

			if m.HeapObjects > maxHeapObjects {
				maxHeapObjects = m.HeapObjects
			}

			if m.HeapObjects < minHeapObjects {
				minHeapObjects = m.HeapObjects
			}

			if cpuUsage > maxCPUUsage {
				maxCPUUsage = cpuUsage
			}

			entities := s.EntitiesManager.GetAllEntities()
			for _, entity := range entities {
				if entity.GetEntityType().String() == "Player" {
					player, ok := entity.(*player.Player)
					if !ok {
						continue
					}
					if player.GetClientState().String() == "Play" {
						for i := 0; i < 60; i++ {
							sendMessage(player, "          ")
						}
						sendMessage(player, "MSPT: "+fmt.Sprintf("%.6fms TPS: %.6f", medianMspt, tps))
						sendMessage(player, "Memoria asignada en el heap: "+fmt.Sprintf("%.2f MB", heapAllocMB))
						sendMessage(player, "Memoria del stack: "+fmt.Sprintf("%.2f MB", stackInUseMB))
						sendMessage(player, "Memoria reservada por Go en el sistema: "+fmt.Sprintf("%.2f MB", sysMB))
						sendMessage(player, "Memoria total usada por Go: "+fmt.Sprintf("%.2f MB", heapAllocMB+stackInUseMB))
						sendMessage(player, "Objetos en el heap: "+fmt.Sprintf("%v", m.HeapObjects))
						sendMessage(player, "Goroutines: "+fmt.Sprintf("%v", runtime.NumGoroutine()))
						sendMessage(player, "Uso de CPU: "+fmt.Sprintf("%.2f%%", cpuUsage))
						sendMessage(player, "Uso máximo de RAM: "+fmt.Sprintf("%.2f MB", maxRamUsage))
						sendMessage(player, "Uso mínimo de RAM: "+fmt.Sprintf("%.2f MB", minRamUsage))
						sendMessage(player, "Máximo de objetos en el heap: "+fmt.Sprintf("%v", maxHeapObjects))
						sendMessage(player, "Mínimo de objetos en el heap: "+fmt.Sprintf("%v", minHeapObjects))
						sendMessage(player, "Máximo uso de CPU: "+fmt.Sprintf("%.2f%%", maxCPUUsage))

					}

				}
			}*/
			//runtime.GC()
		}

		// Calcula el tiempo restante para alcanzar el tickRate
		sleepTime := s.tickRate - elapsedTime
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}

var packetPool = sync.Pool{
	New: func() interface{} {
		return clientbound.NewSystemChatMessagePacket_1_21("", false)
	},
}

func sendMessage(player *player.Player, message string) {
	packet := packetPool.Get().(*clientbound.SystemChatMessagePacket_1_21)
	packet.SetContent(message)
	packet.GetPacket(player).Send(player)
	packetPool.Put(packet)
}
func (s *Server) tick(tickNumber int) {
	// Lógica del tick del servidor aquí
	// por ejemplo, actualizar la posición de los objetos, manejar la IA, enviar paquetes a los clientes, etc.
	packets.RunPlayTick(s.EntitiesManager)
	tick_1_21.TickEntities(tickNumber)
}
func bytesToMB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024)
}
func GetCPUUsage() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
		return 0
	}
	if len(percent) > 0 {
		return percent[0]
	}
	return 0
}
