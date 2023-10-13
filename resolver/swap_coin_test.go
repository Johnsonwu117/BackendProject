package resolver

import (
	coinModel "backendProject/db"
	"testing"
)

func TestChangeCoin(t *testing.T) {
	tests := []struct {
		input    *coinModel.InputCoin
		expected float64
	}{
		{
			input: &coinModel.InputCoin{
				Source: "TWD",
				Target: "TWD",
				Amount: 100,
			},
			expected: 100.00,
		},
		{
			input: &coinModel.InputCoin{
				Source: "TWD",
				Target: "JPY",
				Amount: 100,
			},
			expected: 366.9,
		},
		{
			input: &coinModel.InputCoin{
				Source: "USD",
				Target: "JPY",
				Amount: 100,
			},
			expected: 11180.1,
		},
		{
			input: &coinModel.InputCoin{
				Source: "USD",
				Target: "JPY",
				Amount: 1525,
			},
			expected: 170496.53,
		},
	}

	output := &coinModel.OutputCoin{}
	for _, test := range tests {
		output = ChangeCoin(test.input)
		if output.Amount != test.expected {
			t.Errorf("For source %s to target %s,  expected %.2f but got %.2f", test.input.Source, test.input.Target, test.expected, output.Amount)
		}
	}
}
