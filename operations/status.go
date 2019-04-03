package operations

import (
	"fmt"
	"strings"

	"github.com/anks333/parking_lot/models"

	"github.com/anks333/parking_lot/store"
)

// status

// Status :
type Status struct {
	opName  string
	storage *models.ParkingLot
}

// NewStatusCMD :
func NewStatusCMD() *Status {
	return &Status{
		opName: "status",
	}
}

// GetName :
func (s *Status) GetName() string {
	return s.opName
}

// Parse :
func (s *Status) Parse(argVal string) error {
	return nil
	// return fmt.Sprintf("Created a parking lot with %v slots", cpl.capacity)
}

// Execute :
func (s *Status) Execute(argVal string) string {
	storage := store.GetStorage()
	if storage == nil {
		return "Parking lot is not created"
	}
	slots, err := storage.GetParkedSpots()
	outStr := []string{
		fmt.Sprintf("%s\t%s\t%s",
			"Slot No.",
			"Registration No",
			"Colour",
		),
	}
	if err != nil {
		outStr = append(outStr, "No Data Found")
		return strings.Join(outStr, "\n")
	}

	for _, slot := range slots {
		outStr = append(outStr,
			fmt.Sprintf(
				"%v\t%v\t%v",
				slot.GetParkingSpotNo(),
				slot.GetVehicle().GetVehicleNo(),
				slot.GetVehicle().GetColor(),
			),
		)
	}
	return strings.Join(outStr, "\n")
}
