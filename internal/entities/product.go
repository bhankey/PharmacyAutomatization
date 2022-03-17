package entities

import "time"

type Product struct {
	ID             int
	Name           string
	Price          string
	ExpirationDate time.Duration
	InstructionURL string
	ImgURL         string
	Comment        string
	RecipeOnly     bool
}

type ProductItem struct {
	ID               int
	ProductID        int
	PharmacyID       int
	ReceiptID        int
	Position         string
	ManufacturedTime string
	ReservationUUID  string
	IsSold           bool
	IsExpired        bool
	Priority         int
}

type PharmacyProductItem struct {
	Name           string
	Price          string
	InstructionURL string
	ImgURL         string
	Comment        string
	RecipeOnly     bool
	Position       string
	Count          int
}

type PurchaseProductItem struct {
	Name  string
	Count int
	Price int
}
