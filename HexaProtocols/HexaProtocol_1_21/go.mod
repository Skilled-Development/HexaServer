module HexaProtocol_1_21

go 1.23.4

replace HexaUtils => ../../HexaUtils

replace HexaServer => ../../HexaServer

require HexaUtils v0.0.0-00010101000000-000000000000

require HexaServer v0.0.0-00010101000000-000000000000

require (
	github.com/google/uuid v1.6.0
	github.com/shirou/gopsutil v3.21.11+incompatible
)

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/sys v0.28.0 // indirect
)