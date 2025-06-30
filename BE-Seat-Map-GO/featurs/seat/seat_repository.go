package seat

import "gorm.io/gorm"

type SeatRepositoryInf interface {
	GetSeatByFlightID(flightId uint) ([]Seat, error)
	GetSaetById(seatId uint) (Seat, error)
	UpdateSeat(seat Seat) (Seat, error)
}

type seatRepository struct {
	db *gorm.DB
}

func SeatRepository(db *gorm.DB) SeatRepositoryInf {
	return &seatRepository{db: db}
}

func (s *seatRepository) GetSeatByFlightID(flightId uint) ([]Seat, error) {
	var seats []Seat

	err := s.db.
		Joins("JOIN seat_rows ON seats.seat_row_id = seat_rows.id").
		Joins("JOIN cabins ON seat_rows.cabin_id = cabins.id").
		Where("cabins.flight_id = ?", flightId).
		Order("seats.row,seats.column").Find(&seats).Error

	if err != nil {
		return nil, err
	}
	return seats, nil

}

func (s *seatRepository) GetSaetById(seatId uint) (Seat, error) {
	var seat Seat
	err := s.db.Find(&seat, seatId).Error
	return seat, err
}

func (s *seatRepository) UpdateSeat(seat Seat) (Seat, error) {
	err := s.db.Save(&seat).Error
	return seat, err
}
