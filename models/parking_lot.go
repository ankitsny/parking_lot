package models

// ParkingLot :
type ParkingLot struct {
	lotID         string
	address       string
	parkingSpots  []*ParkingSpot
	size          int
	nextSpot      int
	hasEmptySpace bool
}
