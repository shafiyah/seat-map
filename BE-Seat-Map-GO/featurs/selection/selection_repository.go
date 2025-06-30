package selection

import (
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
	"gorm.io/gorm"
)

type SelectionRepositoryInfc interface {
	InsertSelection(selection SeatSelection) (SeatSelection, error)
}

type selectionRepository struct {
	db *gorm.DB
}

func SelectionRepository(db *gorm.DB) SelectionRepositoryInfc {
	return &selectionRepository{db: db}
}

func (r *selectionRepository) InsertSelection(selection SeatSelection) (SeatSelection, error) {

	var existingSelection SeatSelection
	err := r.db.Where("flight_id = ? AND passenger_id = ?", selection.FlightID, selection.PassengerID).First(&existingSelection).Error
	if err == nil {
		var seat seat.Seat
		r.db.Model(&seat).Where("id = ?", existingSelection.SeatID).Update("available", true)

		r.db.Delete(&existingSelection)
	}

	err = r.db.Create(&selection).Error
	return selection, err
}
