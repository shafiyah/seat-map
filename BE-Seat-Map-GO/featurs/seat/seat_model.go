package seat

type Seat struct {
	ID                  uint `gorm:"primaryKey"`
	SeatRowID           uint
	Code                string
	Available           bool
	StorefrontSlotCode  string
	Price               float64
	Taxes               float64
	Total               float64
	CurrencyPrice       string
	CurrencyTaxes       string
	CurrencyTotal       string
	SeatCharacteristics string
	Designations        string
	Limitations         string
	Location            string
	Row                 uint
	Column              string
}
