package operations

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anks333/parking_lot/store"
)

// registration_numbers_for_cars_with_colour

// RegistrationNoBasedOnColor :
type RegistrationNoBasedOnColor struct {
	opName string
	color  string
}

// NewRegistrationNoBasedOnColorCMD :
func NewRegistrationNoBasedOnColorCMD() *RegistrationNoBasedOnColor {
	return &RegistrationNoBasedOnColor{
		opName: "registration_numbers_for_cars_with_colour",
	}
}

// GetName :
func (rbc *RegistrationNoBasedOnColor) GetName() string {
	return rbc.opName
}

// Parse :
func (rbc *RegistrationNoBasedOnColor) Parse(argVal string) error {
	args := strings.Split(argVal, " ")
	if args[0] == "" {
		return errors.New("Invalid args for regitration no based on color")
	}
	rbc.color = args[0]
	return nil
	// return fmt.Sprintf("Created a parking lot with %v slots", cpl.capacity)
}

// Execute :
func (rbc *RegistrationNoBasedOnColor) Execute(argVal string) string {
	if err := rbc.Parse(argVal); err != nil {
		return fmt.Sprintf("Invalid args for command")
	}
	slots, err := store.GetStorage().GetParkingSpotByColor(rbc.color)
	if err != nil {
		return "Not Found" // TODO: pass proper error
	}
	var out []string
	for _, slot := range slots {
		out = append(out, string(slot.GetParkingSpotNo()))
	}
	return strings.Join(out, ", ")
}
