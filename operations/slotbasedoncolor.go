package operations

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anks333/parking_lot/store"
)

// slot_numbers_for_cars_with_colour

// SlotNoBasedOnColor :
type SlotNoBasedOnColor struct {
	opName string
	color  string
}

// NewSlotNoBasedOnColorCMD :
func NewSlotNoBasedOnColorCMD() *SlotNoBasedOnColor {
	return &SlotNoBasedOnColor{
		opName: "slot_numbers_for_cars_with_colour",
	}
}

// GetName :
func (rbc *SlotNoBasedOnColor) GetName() string {
	return rbc.opName
}

// Parse :
func (rbc *SlotNoBasedOnColor) Parse(argVal string) error {
	args := strings.Split(argVal, " ")
	if args[0] == "" {
		return errors.New("Invalid args for slot no based on color")
	}
	rbc.color = args[0]
	return nil
	// return fmt.Sprintf("Created a parking lot with %v slots", cpl.capacity)
}

// Execute :
func (rbc *SlotNoBasedOnColor) Execute(argVal string) string {
	if err := rbc.Parse(argVal); err != nil {
		return fmt.Sprintf("Invalid args")
	}
	slots, err := store.GetStorage().GetParkingSpotByColor(rbc.color)
	if err != nil {
		return "Not Found" // TODO: pass proper error
	}
	var out []string
	for _, slot := range slots {
		out = append(out, slot.GetVehicle().GetVehicleNo())
	}
	return strings.Join(out, ", ")
}
