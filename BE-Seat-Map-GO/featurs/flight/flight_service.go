package flight

type FlightServiceInf interface {
	ImportAndMappingFile(jasonData map[string]interface{}) error
	GetFlightById(flightId uint) (Flight, error)
}

type flightService struct {
	flightRepo FlightRepositoryInf
}

func FlightService(flightRepo FlightRepositoryInf) *flightService {
	return &flightService{flightRepo}
}

func (s *flightService) ImportAndMappingFile(jsonData map[string]interface{}) error {
	return s.flightRepo.ImportAndMappingFile(jsonData)
}

func (s *flightService) GetFlightById(flightId uint) (Flight, error) {
	return s.flightRepo.GetFlightById(flightId)
}
