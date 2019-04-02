package operations

import (
	"fmt"
	"strconv"

	"github.com/anks333/parking_lot/store"
)

// CreateParkingLot :
// this struct should implement ICommand interface
type CreateParkingLot struct {
	opName   string
	capacity int
}

// NewCreateParkingLot :
func NewCreateParkingLot() *CreateParkingLot {
	return &CreateParkingLot{
		opName: "create_parking_lot",
	}
}

// GetName :
func (cpl *CreateParkingLot) GetName() string {
	return cpl.opName
}

// Parse :
func (cpl *CreateParkingLot) Parse(argVal string) error {
	num, err := strconv.Atoi(argVal)

	cpl.capacity = num

	return err
	// return fmt.Sprintf("Created a parking lot with %v slots", cpl.capacity)
}

// Execute :
func (cpl *CreateParkingLot) Execute(argVal string) string {
	if err := cpl.Parse(argVal); err != nil {
		return "Invalid arguments"
	}
	store.NewStore(cpl.capacity)

	return fmt.Sprintf("Created a parking lot with %v slots", cpl.capacity)
}
