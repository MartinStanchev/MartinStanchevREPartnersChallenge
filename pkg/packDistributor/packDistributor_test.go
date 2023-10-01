package packDistributor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistribute(t *testing.T) {
	tests := []struct {
		input     int
		want      map[int]int
		expectErr bool
	}{
		{
			input:     1,
			want:      map[int]int{250: 1},
			expectErr: false,
		},
		{
			input:     250,
			want:      map[int]int{250: 1},
			expectErr: false,
		},
		{
			input:     251,
			want:      map[int]int{500: 1},
			expectErr: false,
		},
		{
			input:     501,
			want:      map[int]int{1000: 1},
			expectErr: false,
		},
		{
			input:     12001,
			want:      map[int]int{5000: 2, 2000: 1, 250: 1},
			expectErr: false,
		},
		{
			input:     -1,
			want:      map[int]int{},
			expectErr: true,
		},
		{
			input:     0,
			want:      map[int]int{},
			expectErr: true,
		},
	}

	packSizes := []int{
		250,
		500,
		1000,
		2000,
		5000,
	}

	distributor := NewDistributor(packSizes)

	for _, tc := range tests {
		got, err := distributor.Distribute(tc.input)
		if err != nil {
			if tc.expectErr == false {
				assert.Failf(t, "Got error: %v, but was not expecting one", err.Error())
			}
		}

		assert.Equal(t, got, tc.want)
	}
}
