package store

import (
	"github.com/anks333/parking_lot/models"
)

var storage *models.ParkingLot

// NewStore :
func NewStore(capacity int) error {
	var err error
	storage, err = models.CreateParkingLot("LOT_1", "Bangalore", capacity)
	return err
}

// GetStorage :
func GetStorage() *models.ParkingLot {
	return storage
}
