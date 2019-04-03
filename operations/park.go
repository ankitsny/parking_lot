package operations

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anks333/parking_lot/store"
)

// Park :
// this struct should implement ICommand interface
type Park struct {
	opName    string
	vehicleNo string
	color     string
	// Add size
}

// NewParkCMD :
func NewParkCMD() *Park {
	return &Park{
		opName: "park",
	}
}

// GetName :
func (p *Park) GetName() string {
	return p.opName
}

// Parse :
func (p *Park) Parse(argVal string) error {
	p.color = ""
	p.vehicleNo = ""
	args := strings.Split(argVal, " ")
	if len(args) != 2 {
		return errors.New("Invalid args for park")
	}
	p.vehicleNo = args[0]
	p.color = args[1]
	return nil
	// return fmt.Sprintf("Created a parking lot with %v slots", cpl.capacity)
}

// Execute :
func (p *Park) Execute(argVal string) string {
	fmt.Println("Park", argVal)
	if err := p.Parse(argVal); err != nil {
		return fmt.Sprintf("Invalid args for park command")
	}
	slot, err := store.GetStorage().Park(p.vehicleNo, p.color)
	if err != nil {
		return "Sorry, parking lot is full" // TODO: pass proper error
	}
	return fmt.Sprintf("Allocated slot number: %v", slot.GetParkingSpotNo())
}
