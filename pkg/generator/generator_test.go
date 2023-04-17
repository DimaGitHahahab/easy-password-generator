package generator

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name           string
		length         int
		excludeUpper   bool
		excludeLower   bool
		excludeNumbers bool
		excludeSpecial bool
		expectedError  bool
	}{
		{
			name:           "length too short",
			length:         3,
			excludeUpper:   false,
			excludeLower:   false,
			excludeNumbers: false,
			excludeSpecial: false,
			expectedError:  true,
		},
		{
			name:           "exclude all character types",
			length:         10,
			excludeUpper:   true,
			excludeLower:   true,
			excludeNumbers: true,
			excludeSpecial: true,
			expectedError:  true,
		},
		{
			name:           "valid password",
			length:         10,
			excludeUpper:   false,
			excludeLower:   false,
			excludeNumbers: false,
			excludeSpecial: false,
			expectedError:  false,
		}, {
			name:           "valid password with excluded Upper",
			length:         10,
			excludeUpper:   true,
			excludeLower:   false,
			excludeNumbers: false,
			excludeSpecial: false,
			expectedError:  false,
		},
		{
			name:           "valid password with excluded Lower",
			length:         10,
			excludeUpper:   true,
			excludeLower:   false,
			excludeNumbers: false,
			excludeSpecial: false,
			expectedError:  false,
		},
		{
			name:           "valid password with excluded Numbers",
			length:         10,
			excludeUpper:   true,
			excludeLower:   false,
			excludeNumbers: false,
			excludeSpecial: false,
			expectedError:  false,
		},
		{
			name:           "valid password with excluded Special",
			length:         10,
			excludeUpper:   true,
			excludeLower:   false,
			excludeNumbers: false,
			excludeSpecial: false,
			expectedError:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := GeneratePassword(test.length, test.excludeUpper, test.excludeLower, test.excludeNumbers, test.excludeSpecial)
			if test.expectedError && err == nil {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}
