package flight

import "github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"

type Flight struct {
	ID                uint `gorm:"primaryKey"`
	Aircraft          string
	FlightNumber      string
	Origin            string
	Destination       string
	DepartureTime     string
	ArrivalTime       string
	Duration          string
	AirlineCode       string
	DepartureTerminal string
	ArrivalTerminal   string
	Cabins            []Cabin
}

type Cabin struct {
	ID          uint `gorm:"primaryKey"`
	FlightID    uint
	Deck        string
	SeatColumns string
	Class       string
	SeatRows    []SeatRow
}

type SeatRow struct {
	ID        uint `gorm:"primaryKey"`
	CabinID   uint
	RowNumber int
	Seats     []seat.Seat
}
