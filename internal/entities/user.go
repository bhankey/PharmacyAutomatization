package entities

type User struct {
	ID                int
	Name              string
	Surname           string
	Email             string
	Password          string
	PasswordHash      string
	DefaultPharmacyID int
}

// TODO use this for all auth operations.
type UserIdentifyData struct {
	IP          string
	UserAgent   string
	FingerPrint string
}
