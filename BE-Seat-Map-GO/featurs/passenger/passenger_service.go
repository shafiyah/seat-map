package passenger

type PassengerServiceInf interface {
	GetPassangerById(psngrId uint) (Passenger, error)
}

type passengerService struct {
	pasengerRepo PassengerRepositoryInf
}

func PassengerService(passengerRepo PassengerRepositoryInf) *passengerService {
	return &passengerService{passengerRepo}
}

func (s *passengerService) GetPassangerById(psngrId uint) (Passenger, error) {
	return s.pasengerRepo.GetPassangerById(psngrId)
}
