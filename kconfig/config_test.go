package kconfig

import (
	"testing"
)

type TestStruct struct {
	Data string
	Name string
}

func TestGetConfig(t *testing.T) {
	type args struct {
		name []string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "test",
			args: args{
				name: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := GetConfig[TestStruct](tt.args.name...)
			t.Log(data.Data, data.Name)
		})
	}
}
