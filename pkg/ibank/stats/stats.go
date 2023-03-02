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

// CategoriesTotal возвращает сумму платежей по каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// CategoriesAvg находит среднее значение расходов по каждый категори
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	
	categories := map[types.Category]types.Money{}

	avg := map[types.Category]int{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		avg[payment.Category]++ 
	}

	for key, value := range categories {
		categories[key] = types.Money( int(value) / avg[key])
 	}

	return categories
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money,) map[types.Category] types.Money {
	
	difference := map[types.Category]types.Money{}

	for key, value := range second {
		if _, ok := first[key]; !ok {
			difference[key] = 0
		}

		difference[key] = value - first[key]
	}

	for key, value := range first {
		if _, ok := second[key]; !ok {
			difference[key] = 0 - value
		}
	}
	
	return difference
}