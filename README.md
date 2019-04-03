# GOJEK Parking Lot Problem

## Steps to run (Assuming go env is already configured)
    1. Clone the Repo
    2. CWD to `prking_log`
    3. Then run `go main.go` for interactive mode.
       run `go main.go ./input.txt` for file mode.
       or run `go build` then ./parking_lot [filePath]? 


## Available Commands: 
    1. create_parking_lot SIZE
    2. park VEHICLE_NO COLOR
    3. leave VEHICLE_NO
    4. registration_numbers_for_cars_with_colour COLOR
    5. slot_numbers_for_cars_with_colour COLOR
    6. slot_number_for_registration_number VEHICLE_NO


## To run test
    `go test ./... --cover`

