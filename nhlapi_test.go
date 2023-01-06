package main

import (
	"testing"

	"github.com/tidwall/gjson"
)

func Test_standardRequest(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Schedule",
			args{"/schedule"},
		}, {
			"Standings",
			args{"/standings"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := standardRequest(tt.args.endpoint)
			if gjson.GetBytes(resp, "copyright").Exists() {
				t.Error("Did not find copyright in response")
			}
		})
	}
}
