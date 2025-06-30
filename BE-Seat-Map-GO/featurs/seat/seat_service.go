package seat

type SeatServiceInf interface {
	GetSeatByFlightID(flightId uint) ([]Seat, error)
	GetSeatById(seatId uint) (Seat, error)
	UpdateSeat(seat Seat) (Seat, error)
}

type seatService struct {
	seatRepo SeatRepositoryInf
}

func SeatService(seatRepo SeatRepositoryInf) *seatService {
	return &seatService{seatRepo}
}

func (s *seatService) GetSeatByFlightID(flightId uint) ([]Seat, error) {
	return s.seatRepo.GetSeatByFlightID(flightId)
}

func (s *seatService) GetSeatById(seatId uint) (Seat, error) {
	return s.seatRepo.GetSaetById(seatId)
}

func (s *seatService) UpdateSeat(seat Seat) (Seat, error) {
	return s.seatRepo.UpdateSeat(seat)
}
