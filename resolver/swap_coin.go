package resolver

import (
	coinModel "backendProject/db"
	"math"
)

func ChangeCoin(input *coinModel.InputCoin) *coinModel.OutputCoin {

	output := &coinModel.OutputCoin{}
	switch input.Source {
	case "TWD":
		switch input.Target {
		// TWD -> TWD
		case "TWD":
			output.Amount = input.Amount
		// TWD -> JPY
		case "JPY":
			output.Amount = input.Amount * 3.669
		// TWD -> USD
		case "USD":
			output.Amount = input.Amount * 0.03281
		}
	case "JPY":
		switch input.Target {
		// JPY -> TWD
		case "TWD":
			output.Amount = input.Amount * 0.26956
		// JPY -> JPY
		case "JPY":
			output.Amount = input.Amount
		// JPY -> USD
		case "USD":
			output.Amount = input.Amount * 0.00885
		}
	case "USD":
		switch input.Target {
		// USD -> TWD
		case "TWD":
			output.Amount = input.Amount * 30.444
		// USD -> JPY
		case "JPY":
			output.Amount = input.Amount * 111.801
		// USD -> USD
		case "USD":
			output.Amount = input.Amount
		}
	}

	output.Amount = math.Round(output.Amount*100) / 100

	return output
}
