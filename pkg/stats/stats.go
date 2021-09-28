package stats

import (
	"github.com/NekruziKabirzoda/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) (types.Money) {

	sum := types.Money(0)

	for i:=0; i<len(payments); i++{
		if payments[i].Status==types.StatusFail{
			continue
		}
			 
		
    sum+=payments[i].Amount
	}
	
	
	sum = sum/types.Money(len(payments)-1)
return sum
}

func TotalInCategory (payments []types.Payment, category types.Category) types.Money{
	sum := types.Money(0)

	for i:=0; i<len(payments); i++{
		if payments[i].Category == category{
			if   payments[i].Status== types.StatusFail{
				continue
			}
    sum+=payments[i].Amount
	}}
  return sum
} 