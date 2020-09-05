package geoip

import (
	"os"
	"testing"

	"github.com/blushft/strana/platform/config"
)

func Test_new(t *testing.T) {
	license := os.Getenv("STRANA_MAX_MIND_LICENSE")
	if license == "" {
		return
	}

	type args struct {
		conf config.Processor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_auto_dl",
			args: args{
				conf: config.Processor{
					Name: "geoip",
					Type: "geoip",
					Options: map[string]interface{}{
						"database_path":    "../../.fixtures/geoip",
						"max_mind_license": license,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := new(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("new() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
