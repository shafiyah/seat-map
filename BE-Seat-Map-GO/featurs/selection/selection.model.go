package selection

import (
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/flight"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/passenger"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
)

type SeatSelection struct {
	ID          uint                `gorm:"primaryKey"`
	PassengerID uint                `json:"passengerId"`
	Passenger   passenger.Passenger `gorm:"foreignKey:PassengerID"`
	FlightID    uint                `json:"flightId"`
	Flight      flight.Flight       `gorm:"foreignKey:FlightID"`
	SeatID      uint                `json:"seatId"`
	Seat        seat.Seat           `gorm:"foreignKey:SeatID"`
}
