package models

import "errors"

// ParkingLot :
type ParkingLot struct {
	lotID         string
	address       string
	parkingSpots  []*ParkingSpot
	size          int
	nextSpot      int
	hasEmptySpace bool
}

// CreateParkingLot : this function reurns a parking lot with the provided size of parking spots
// by default paking spot size will be "SM" (ASSUMPTION)
// by default level will be "0" (ASSUMPTION)
func CreateParkingLot(lotID, address string, size int) (*ParkingLot, error) {
	if size <= 0 {
		return nil, errors.New("Invalid Parking Spot Aize")
	}

	var parkingSpots []*ParkingSpot

	for i := 0; i < size; i++ {
		parkingSpots = append(parkingSpots, NewEmptyParkingSpot("SM", i+1, "0"))
	}
	pl := &ParkingLot{
		address:       "Bangalore HSR",
		hasEmptySpace: true,
		lotID:         "LOT_1",
		nextSpot:      0,
		size:          size,
		parkingSpots:  parkingSpots,
	}
	return pl, nil
}
