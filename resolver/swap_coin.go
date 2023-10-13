package resolver

import (
	dbtest "backendProject/db"
	"fmt"
)

func ChangeCoin(input *dbtest.InputCoin) any {
	fmt.Println("resevoler", input.Source, input.Target, input.Amount)
	return nil
}
