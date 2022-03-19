package entities

type Receipt struct {
	ID         int
	UserID     int
	PharmacyID int
	Sum        float64
	Discount   int
}
