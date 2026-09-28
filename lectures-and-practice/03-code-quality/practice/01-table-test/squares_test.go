package tabletest

import (
	"errors"
	"testing"
)

func TestSumSquares(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
		wantErr error
	}{
		{name: "regular", numbers: []int{1, 2, 3}, want: 14},
		{name: "empty", numbers: []int{}, want: 0},
		{name: "negative", numbers: []int{1, -2, 3}, wantErr: ErrNegativeNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumSquares(tt.numbers)

			if got != tt.want {
				t.Errorf("SumSquares() = %d, want %d", got, tt.want)
			}
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("SumSquares() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
