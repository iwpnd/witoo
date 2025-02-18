package witoo_test

import (
	"fmt"
	"testing"

	"github.com/iwpnd/witoo"
	"github.com/stretchr/testify/assert"
)

func TestSayHiTo(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "EmptyName",
			input:          "",
			expectedOutput: "",
			expectedError:  fmt.Errorf("name cannot be empty"),
		},
		{
			name:           "ValidName",
			input:          "John",
			expectedOutput: "Hi, John",
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := witoo.SayHiTo(tc.input)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error(), "expected error did not match")
			} else {
				assert.NoError(t, err, "unexpected error")
				assert.Equal(t, tc.expectedOutput, output, "output did not match")
			}
		})
	}
}

