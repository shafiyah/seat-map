package passenger

import "gorm.io/gorm"

type PassengerRepositoryInf interface {
	GetPassangerById(psngrId uint) (Passenger, error)
}

type passangerRepository struct {
	db *gorm.DB
}

func PassangerRepository(db *gorm.DB) PassengerRepositoryInf {
	return &passangerRepository{db: db}
}

func (r *passangerRepository) GetPassangerById(psngrId uint) (Passenger, error) {
	var passenger Passenger
	err := r.db.First(&passenger, psngrId).Error
	return passenger, err
}
