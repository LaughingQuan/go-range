package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseConfig(t *testing.T) {
	copyFile("./config.yml", "./../config.yml")
	defer os.Remove("./config.yml")
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := parseConfig(); (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func copyFile(dst, src string) error {
	srcContent, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, srcContent, 0644)
	return err
}
