package selection

import (
	"github.com/shafiyah/seat-map/BE-Seat-Map-Go/featurs/seat"
)

type SelecetionServiceInfc interface {
	SeatSelection(selection SeatSelection) error
}

type selectionService struct {
	selectionRepo SelectionRepositoryInfc
	seatService   seat.SeatServiceInf
}

func SelectionService(selectionRepo SelectionRepositoryInfc, seatService seat.SeatServiceInf) *selectionService {
	return &selectionService{selectionRepo: selectionRepo, seatService: seatService}
}

func (s *selectionService) SeatSelection(selection SeatSelection) error {

	if _, err := s.selectionRepo.InsertSelection(selection); err != nil {
		return err
	}

	seat, err := s.seatService.GetSeatById(selection.SeatID)
	if err != nil {
		return err
	}

	seat.Available = false

	_, err = s.seatService.UpdateSeat(seat)
	if err != nil {
		return err
	}

	return nil

}
