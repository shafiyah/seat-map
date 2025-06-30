package main

import (
	"fmt"
	"log"

	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/configs"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/flight"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/passenger"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/selection"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/router"
)

func main() {

	db, err := configs.ConnectDB()
	if err != nil {
		log.Fatal("Error coonect to database")
	}
	fmt.Printf("Success connect to database")

	db.AutoMigrate(
		&flight.Flight{},
		&flight.Cabin{},
		&flight.SeatRow{},
		&seat.Seat{},
		&selection.SeatSelection{},
		&passenger.Passenger{},
	)

	r := router.NewRouter(db)
	r.Run(":8080")

}
