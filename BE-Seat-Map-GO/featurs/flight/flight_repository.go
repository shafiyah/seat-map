package flight

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/passenger"
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
	"gorm.io/gorm"
)

type FlightRepositoryInf interface {
	ImportAndMappingFile(jasonData map[string]interface{}) error
	GetFlightById(flightId uint) (Flight, error)
}

type flightRepository struct {
	db *gorm.DB
}

func FlightRepository(db *gorm.DB) FlightRepositoryInf {
	return &flightRepository{db: db}
}

func (r *flightRepository) ImportAndMappingFile(jsonData map[string]interface{}) error {

	seatsParts := jsonData["seatsItineraryParts"].([]interface{})
	segmentSeatMap := seatsParts[0].(map[string]interface{})["segmentSeatMaps"].([]interface{})[0].(map[string]interface{})

	segment := segmentSeatMap["segment"].(map[string]interface{})
	passengerSeatMaps := segmentSeatMap["passengerSeatMaps"].([]interface{})
	seatMap := passengerSeatMaps[0].(map[string]interface{})["seatMap"].(map[string]interface{})
	passengerMap := passengerSeatMaps[0].(map[string]interface{})["passenger"].(map[string]interface{})

	flight := Flight{

		Aircraft:          fmt.Sprint(seatMap["aircraft"]),
		FlightNumber:      fmt.Sprint(segment["flight"].(map[string]interface{})["flightNumber"]),
		Origin:            fmt.Sprint(segment["origin"]),
		Destination:       fmt.Sprint(segment["destination"]),
		DepartureTime:     fmt.Sprint(segment["departure"]),
		ArrivalTime:       fmt.Sprint(segment["arrival"]),
		Duration:          fmt.Sprint(segment["duration"]),
		AirlineCode:       fmt.Sprint(segment["flight"].(map[string]interface{})["airlineCode"]),
		DepartureTerminal: fmt.Sprint(segment["flight"].(map[string]interface{})["departureTerminal"]),
		ArrivalTerminal:   fmt.Sprint(segment["flight"].(map[string]interface{})["arrivalTerminal"]),
	}

	fmt.Println(flight)

	if err := r.db.Create(&flight).Error; err != nil {
		return err
	}

	cabins := seatMap["cabins"].([]interface{})
	for _, c := range cabins {
		cabinObj := c.(map[string]interface{})
		seatColumns := extractStringArray(cabinObj["seatColumns"])
		cabin := Cabin{
			FlightID:    flight.ID,
			Deck:        fmt.Sprint(cabinObj["deck"]),
			SeatColumns: strings.Join(seatColumns, ","),
			Class:       fmt.Sprint(segment["cabinClass"]),
		}
		if err := r.db.Create(&cabin).Error; err != nil {
			return err
		}

		seatRows := cabinObj["seatRows"].([]interface{})
		for _, robj := range seatRows {
			sr := robj.(map[string]interface{})
			row := SeatRow{
				CabinID:   cabin.ID,
				RowNumber: int(sr["rowNumber"].(float64)),
			}

			if err := r.db.Create(&row).Error; err != nil {
				return err
			}

			seats := sr["seats"].([]interface{})

			for _, s := range seats {
				seatObj := s.(map[string]interface{})
				if fmt.Sprint(seatObj["storefrontSlotCode"]) == "SEAT" {
					price := seatObj["prices"].(map[string]interface{})
					alt := price["alternatives"].([]interface{})[0].([]interface{})[0].(map[string]interface{})

					total := seatObj["total"].(map[string]interface{})
					altTotal := total["alternatives"].([]interface{})[0].([]interface{})[0].(map[string]interface{})

					tax := 0.0
					currency := ""
					if taxes, ok := seatObj["taxes"].(map[string]interface{}); ok {
						if alts, ok := taxes["alternatives"].([]interface{}); ok && len(alts) > 0 {
							altTaxs := taxes["alternatives"].([]interface{})[0].([]interface{})[0].(map[string]interface{})
							tax = altTaxs["amount"].(float64)
							currency = fmt.Sprint(altTaxs["currency"])
						}

					}

					rowNumStr, column := splitSeatCode(fmt.Sprint(seatObj["code"]))
					rowNum, _ := strconv.Atoi(rowNumStr)

					newSeat := seat.Seat{
						SeatRowID:           row.ID,
						Code:                fmt.Sprint(seatObj["code"]),
						Available:           seatObj["available"].(bool),
						StorefrontSlotCode:  fmt.Sprint(seatObj["storefrontSlotCode"]),
						Price:               alt["amount"].(float64),
						CurrencyPrice:       fmt.Sprint(alt["currency"]),
						Total:               altTotal["amount"].(float64),
						CurrencyTotal:       fmt.Sprint(alt["currency"]),
						Taxes:               tax,
						CurrencyTaxes:       currency,
						SeatCharacteristics: joinStrings(seatObj["seatCharacteristics"]),
						Designations:        joinStrings(seatObj["designations"]),
						Limitations:         joinStrings(seatObj["limitations"]),
						Location:            locationSeat(seatObj["seatCharacteristics"]),
						Row:                 uint(rowNum),
						Column:              column,
					}
					if err := r.db.Create(&newSeat).Error; err != nil {
						return err
					}
				}
			}
		}

		passengerDetail := passengerMap["passengerDetails"].(map[string]interface{})
		passengerInfo := passengerMap["passengerInfo"].(map[string]interface{})
		passengerEmail := passengerInfo["emails"].([]interface{})

		passenger := passenger.Passenger{

			FirstName:   fmt.Sprint(passengerDetail["firstName"]),
			LastName:    fmt.Sprint(passengerDetail["lastName"]),
			DateOfBirth: fmt.Sprint(passengerInfo["dateOfBirth"]),
			Gender:      fmt.Sprint(passengerInfo["gender"]),
			Email:       fmt.Sprint(passengerEmail[0]),
		}
		if err := r.db.Create(&passenger).Error; err != nil {
			return err
		}

	}

	return nil
}

func extractStringArray(data interface{}) []string {
	arr := data.([]interface{})
	result := make([]string, 0, len(arr))
	for _, item := range arr {
		result = append(result, fmt.Sprint(item))
	}
	return result
}

func joinStrings(field interface{}) string {

	if field == nil {
		return ""
	}

	var seatCharMeaning = map[string]string{
		"W":  "Window",
		"A":  "Aisle",
		"M":  "Middle",
		"CH": "Chargeable",
		"FC": "Front Cabin",
		"OW": "Over Wing",
		"EX": "Extra Legroom",
		"RS": "Right Side",
		"LS": "Left Side",
		"B":  "Bulkhead",
		"X":  "Exit Row",
	}

	list := field.([]interface{})
	out := []string{}

	for _, v := range list {
		out = append(out, seatCharMeaning[fmt.Sprint(v)])
	}
	return strings.Join(out, ",")

}

func locationSeat(field interface{}) string {

	seatCharacteristics := field.([]interface{})
	has := func(key string) bool {
		for _, ch := range seatCharacteristics {
			if ch == fmt.Sprint(key) {
				return true
			}
		}
		return false
	}

	switch {
	case has("W"):
		return "Window"
	case has("A"):
		return "Aisle"
	default:
		return "Middle"
	}
}

func splitSeatCode(code string) (row string, column string) {
	for _, c := range code {
		if c >= '0' && c <= '9' {
			row += string(c)
		} else {
			column += string(c)
		}
	}
	return
}

func (r *flightRepository) GetFlightById(flightId uint) (Flight, error) {
	var flight Flight
	err := r.db.First(&flight, flightId).Error
	return flight, err
}
