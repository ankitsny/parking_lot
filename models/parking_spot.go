package models

import "time"

// ParkingSpot :
type ParkingSpot struct {
	// TODO: change the size of parking spot to enum/constant
	size        string     // size of the parking spot
	spotNo      int        // Parking spot no
	vehicle     *Vehicle   // Vehicle itself
	parkingTime *time.Time // parking time
	level       string     // TODO: make it enum
}

// NewEmptyParkingSpot : function to create empty parking spot
func NewEmptyParkingSpot(size string, spotNo int, level string) *ParkingSpot {
	return &ParkingSpot{
		size:   size,
		spotNo: spotNo,
		level:  level,
	}
}

// GetVehicle : return vehicle
func (ps *ParkingSpot) GetVehicle() *Vehicle {
	return ps.vehicle
}

// ParkVehicle : parks a vehice the spot
func (ps *ParkingSpot) ParkVehicle(v *Vehicle) {
	ps.vehicle = v
}

// VackateVehicle : parks a vehice the spot
func (ps *ParkingSpot) VackateVehicle() {
	// TODO: Before vacating make a backup some where
	ps.vehicle = nil
	ps.parkingTime = nil
}
