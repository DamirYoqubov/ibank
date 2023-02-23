package stats

import "github.com/DamirYoqubov/import/v2/pkg/import/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	
	sum := types.Money(0)
	
	count := int32(0)

	for _, payment := range payments {
		if payment.Status != types.StatusFail{
		sum += payment.Amount
		count++
		}
	}
	return sum/types.Money(count)
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail{
			sum += payment.Amount
		}
	}
	return sum
}

// FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	
	return filtered
}