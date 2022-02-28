package entities

type User struct {
	ID                int
	Name              string
	Email             string
	PasswordHash      string
	DefaultPharmacyID int
}

type UserIdentifyData struct {
}
