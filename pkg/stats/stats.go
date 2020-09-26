package stats

import (
	"github.com/Qadriddin-dev/bank/v2/pkg/types"
)


func Avg(payments []types.Payment) types.Money {
	var mid types.Money
	for _, payment := range payments {
		if payment.Status == types.StatusOk{
			mid += payment.Amount
		}
	}
	length := len(payments)
	mid /= types.Money(length)
	return mid
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}