package main

import "testing"

func TestCreateYaml(t *testing.T) {
	type args struct {
		data interface{}
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				data: map[string]interface{}{
					"name": "test",
					"age":  18,
				},
				path: "./test.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateYaml(tt.args.data, tt.args.path)
		})
	}
}
