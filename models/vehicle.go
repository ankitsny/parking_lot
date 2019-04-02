package models

// Vehicle : Vehicle type to create and store a vehicle in parking spot
type Vehicle struct {
	color     string // Color of the vehicle
	vehicleNo string // Registration no of vehicle

	// TODO: change type of size, string to enum/constant
	size string // stores the vehicle size
}

// CreateVehicle : create a new vehicle
func CreateVehicle(color, vehicleNo, size string) *Vehicle {
	return &Vehicle{
		color:     color,
		vehicleNo: vehicleNo,
		size:      size,
	}
}

// GetColor : returns color of the vehicle
func (v *Vehicle) GetColor() string {
	return v.color
}

// GetVehicleNo : returns vehicle no
func (v *Vehicle) GetVehicleNo() string {
	return v.vehicleNo
}

// GetSize : returns size of the vehicle
func (v *Vehicle) GetSize() string {
	return v.size
}
