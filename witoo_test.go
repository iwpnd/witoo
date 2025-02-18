package witoo_test

import (
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
		},
		{
			name:           "ValidName",
			input:          "John",
			expectedOutput: "Hi, John",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := witoo.SayHiTo(tc.input)
			assert.Equal(t, tc.expectedOutput, output, "output did not match")
		})
	}
}
