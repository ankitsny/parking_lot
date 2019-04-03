package models

import "errors"

// ParkingLot :
type ParkingLot struct {
	lotID         string
	address       string
	parkingSpots  []*ParkingSpot
	capacity      int
	nextSpot      int
	hasEmptySpace bool
}

// CreateParkingLot : this function reurns a parking lot with the provided size of parking spots
// by default paking spot size will be "SM" (ASSUMPTION)
// by default level will be "0" (ASSUMPTION)
func CreateParkingLot(lotID, address string, capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, errors.New("Invalid Parking Spot Aize")
	}

	var parkingSpots []*ParkingSpot

	for i := 0; i < capacity; i++ {
		parkingSpots = append(parkingSpots, NewEmptyParkingSpot("SM", i+1, "0"))
	}
	pl := &ParkingLot{
		address:       "Bangalore HSR",
		hasEmptySpace: true,
		lotID:         "LOT_1",
		nextSpot:      1,
		capacity:      capacity,
		parkingSpots:  parkingSpots,
	}
	return pl, nil
}

// GetNearestParkingSpot :
// TODO: we can pass vehicle size in params to incorporate the vehicle size feature
func (pl *ParkingLot) GetNearestParkingSpot() (int, error) {
	for i, spot := range pl.parkingSpots {
		if spot.GetVehicle() == nil {
			return i + 1, nil
		}
	}
	return 0, errors.New("There is no parking spot available")
}

// GetParkingSpotByVehicleNo : this method returns the parking spot object based on vehicle no
// if the vehicle is not parked in the parking lot then returns error
func (pl *ParkingLot) GetParkingSpotByVehicleNo(vNo string) (*ParkingSpot, error) {
	for _, spot := range pl.parkingSpots {
		if spot.GetVehicle() != nil && spot.GetVehicle().GetVehicleNo() == vNo {
			return spot, nil
		}
	}
	return nil, errors.New("Vehicle is not parked in this parking lot :(")
}

// GetParkingSpotByColor : this method returns the parking spot object based on vehicle's color
// if the vehicle is not parked with given color in the parking lot then returns error
func (pl *ParkingLot) GetParkingSpotByColor(color string) ([]*ParkingSpot, error) {
	var spots []*ParkingSpot
	for _, spot := range pl.parkingSpots {
		if spot.GetVehicle() != nil && spot.GetVehicle().GetColor() == color {
			spots = append(spots, spot)
		}
	}
	if len(spots) == 0 {
		return nil, errors.New(color + " Color Vehicle is not parked in this parking lot :(")
	}
	return spots, nil
}

// Park :
func (pl *ParkingLot) Park(vehicleNo string, color string) (*ParkingSpot, error) {
	i, err := pl.GetNearestParkingSpot()
	if err != nil {
		return nil, err
	}
	pl.parkingSpots[i-1].ParkVehicle(CreateVehicle(color, vehicleNo, "SM"))
	return pl.parkingSpots[i-1], nil
}
