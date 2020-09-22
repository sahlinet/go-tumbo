package project

import (
	"reflect"
	"testing"
)

func TestNewProject(t *testing.T) {
	tests := []struct {
		name string
		want Project
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := NewProject(); !reflect.DeepEqual(got, tt.want) {
				if err != nil {
					t.Errorf("err %s not expected", err.Error())
				}
				t.Errorf("NewProject() = %v, want %v", got, tt.want)
			}
		})
	}
}
