package model

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("Valid sha256 hashsum"))
	validHashSum := fmt.Sprintf("%x", h.Sum(nil))

	type args struct {
		name   string
		date   string
		number int
		sum    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid document",
			args: args{
				name:   "Valid name",
				date:   "20201102",
				number: 1,
				sum:    validHashSum,
			},
			wantErr: false,
		},
		{
			name: "Invalid checksum",
			args: args{
				name:   "Valid name",
				date:   "20201102",
				number: 1,
				sum:    strings.Repeat("a", 63),
			},
			wantErr: true,
		},
		{
			name: "Invalid document(empty name)",
			args: args{
				name:   "",
				date:   "20201102",
				number: 1,
				sum:    strings.Repeat("a", 64),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.name, tt.args.date, tt.args.number, tt.args.sum)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
