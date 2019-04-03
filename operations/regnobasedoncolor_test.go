package operations

import (
	"reflect"
	"testing"
)

func TestNewRegistrationNoBasedOnColorCMD(t *testing.T) {
	tests := []struct {
		name string
		want *RegistrationNoBasedOnColor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegistrationNoBasedOnColorCMD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegistrationNoBasedOnColorCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistrationNoBasedOnColor_GetName(t *testing.T) {
	type fields struct {
		opName string
		color  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &RegistrationNoBasedOnColor{
				opName: tt.fields.opName,
				color:  tt.fields.color,
			}
			if got := rbc.GetName(); got != tt.want {
				t.Errorf("RegistrationNoBasedOnColor.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistrationNoBasedOnColor_Parse(t *testing.T) {
	type fields struct {
		opName string
		color  string
	}
	type args struct {
		argVal string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &RegistrationNoBasedOnColor{
				opName: tt.fields.opName,
				color:  tt.fields.color,
			}
			if err := rbc.Parse(tt.args.argVal); (err != nil) != tt.wantErr {
				t.Errorf("RegistrationNoBasedOnColor.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegistrationNoBasedOnColor_Execute(t *testing.T) {
	type fields struct {
		opName string
		color  string
	}
	type args struct {
		argVal string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbc := &RegistrationNoBasedOnColor{
				opName: tt.fields.opName,
				color:  tt.fields.color,
			}
			if got := rbc.Execute(tt.args.argVal); got != tt.want {
				t.Errorf("RegistrationNoBasedOnColor.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
