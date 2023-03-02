package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DamirYoqubov/import/v2/pkg/import/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 3,
		},
		{
			Amount: 4,
		},
		{
			Amount: 5,
		},
	}
	avg := Avg(payments)
	fmt.Println(avg)
	//Output: 4
}

func ExampleTotalInCategory()  {
	category := types.Category("Auto")
	category2 := types.Category("car")
	payments := []types.Payment{
		{
			Category: category,
			Amount: 50_00,
		},
		{
			Category: category,
			Amount: 10_00,
		},
		{
			Category: category2,
			Amount: 20_00,
		},
	}
	result := TotalInCategory(payments, category)

	fmt.Println(result)
	//Output: 6000
}

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) !=0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFIlterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid expexcted: %v; result: %v", expected, result)
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 2_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 1, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_00,
		"food": 2_000_00,
		"fun": 5_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid expexcted: %v; result: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid expexcted: %v; result: %v", expected, result)
	}
}