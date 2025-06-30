package passenger

type Passenger struct {
	ID          uint `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Email       string
	DateOfBirth string
	Gender      string
}
