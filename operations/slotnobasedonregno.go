package operations

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/anks333/parking_lot/store"
)

// slot_number_for_registration_number

// SlotNoBasedOnRegNo :
type SlotNoBasedOnRegNo struct {
	opName    string
	vehicleNo string
}

// NewSlotNoBasedOnRegNoCMD :
func NewSlotNoBasedOnRegNoCMD() *SlotNoBasedOnRegNo {
	return &SlotNoBasedOnRegNo{
		opName: "slot_number_for_registration_number",
	}
}

// GetName :
func (rbc *SlotNoBasedOnRegNo) GetName() string {
	return rbc.opName
}

// Parse :
func (rbc *SlotNoBasedOnRegNo) Parse(argVal string) error {
	args := strings.Split(argVal, " ")
	if args[0] == "" {
		return errors.New("Invalid args for slot no based on reg no")
	}
	rbc.vehicleNo = args[0]
	return nil
	// return fmt.Sprintf("Created a parking lot with %v slots", cpl.capacity)
}

// Execute :
func (rbc *SlotNoBasedOnRegNo) Execute(argVal string) string {
	if err := rbc.Parse(argVal); err != nil {
		return fmt.Sprintf("Invalid args")
	}
	slot, err := store.GetStorage().GetParkingSpotByVehicleNo(rbc.vehicleNo)
	if err != nil {
		return "Not Found" // TODO: pass proper error
	}
	r := strconv.Itoa(slot.GetParkingSpotNo())
	return r
}
