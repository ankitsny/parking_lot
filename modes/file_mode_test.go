package modes

import (
	"reflect"
	"testing"
)

func TestNewFileMode(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *FileMode
	}{
		{
			name: "1",
			args: args{
				path: "./input.txt",
			},
			want: &FileMode{filePath: "./input.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileMode(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileMode_Start(t *testing.T) {
	type fields struct {
		filePath string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "1",
			fields: fields{
				filePath: "../input.txt",
			},
		},
		{
			name: "2. invalid file path",
			fields: fields{
				filePath: "../input.txtt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FileMode{
				filePath: tt.fields.filePath,
			}
			fm.Start()
		})
	}
}
