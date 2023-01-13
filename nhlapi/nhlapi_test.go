package nhlapi

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
		}, {
			"Franchises",
			args{"/franchises"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := standardRequest(tt.args.endpoint)
			if gjson.GetBytes(resp, "Copyright").Exists() {
				t.Error("Did not find copyright in response")
			}
		})
	}
}

func Test_today(t *testing.T) {
	tests := []struct {
		name string
		// wantErr bool
	}{
		{
			"Standard",
			// true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule := Today()
			if schedule.Copyright == "" {
				t.Errorf("Did not find copyright in response")
			}
		})
	}
}
