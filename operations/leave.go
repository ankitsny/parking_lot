package operations

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anks333/parking_lot/store"
)

// Leave :
// this struct should implement ICommand interface
type Leave struct {
	opName    string
	vehicleNo string
}

// NewLeaveCMD :
func NewLeaveCMD() *Leave {
	return &Leave{
		opName: "leave",
	}
}

// GetName :
func (l *Leave) GetName() string {
	return l.opName
}

// Parse :
func (l *Leave) Parse(argVal string) error {
	args := strings.Split(argVal, " ")
	if args[0] == "" {
		return errors.New("Invalid args for leave")
	}
	l.vehicleNo = args[0]
	return nil
}

// Execute :
func (l *Leave) Execute(argVal string) string {
	if err := l.Parse(argVal); err != nil {
		return fmt.Sprintf("Invalid args for leave command")
	}
	slot, err := store.GetStorage().GetParkingSpotByVehicleNo(l.vehicleNo)
	if err != nil {
		return "Not Found" // TODO: pass proper error
	}
	slot.VackateVehicle()
	return fmt.Sprintf("Slot number %v is free", slot.GetParkingSpotNo())
}
