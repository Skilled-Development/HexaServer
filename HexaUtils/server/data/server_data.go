package data

import "HexaUtils/regionreader"

var ServerDataInstance *ServerData

type ServerData struct {
	RegionsLoadedList []regionreader.Region
}

func NewServerData() *ServerData {
	data := &ServerData{
		RegionsLoadedList: []regionreader.Region{},
	}
	ServerDataInstance = data
	return data
}

func (s *ServerData) GetRegionsLoadedList() []regionreader.Region {
	return s.RegionsLoadedList
}

func GetRegionsLoadedList() []regionreader.Region {
	return ServerDataInstance.GetRegionsLoadedList()
}

func AddRegion(region regionreader.Region) {
	ServerDataInstance.AddRegion(region)
}

func (s *ServerData) AddRegion(region regionreader.Region) {
	s.RegionsLoadedList = append(s.RegionsLoadedList, region)
}
