package stats

import "github.com/DamirYoqubov/import/pkg/import/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	
	sum := types.Money(0)
	
	count := int32(0)

	for _, payment := range payments {
		sum += payment.Amount
		count++
	}
	return sum/types.Money(count)
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	 
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}