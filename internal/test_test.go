package internal

import (
	"log"
	"testing"
)

func Test_timeMst(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := timeMst()
			log.Println(got)
		})
	}
}
