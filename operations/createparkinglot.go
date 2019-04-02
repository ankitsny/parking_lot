package operations

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
