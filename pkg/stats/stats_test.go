package stats

import (
	"reflect"
	"testing"

	"github.com/Qadriddin-dev/bank/v2/pkg/types"
)


func TestCategoriesAvg_multiple(t *testing.T) {
	payments := []types.Payment{
	  {
		ID:       1,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       2,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       3,
		Category: "restaurant",
		Amount:   10_00,
	  },
	  {
		ID:       2,
		Category: "restaurant",
		Amount:   100_00,
	  },
	}
	  expected := map[types.Category]types.Money{
	  "cafe":100_00,
	  "restaurant":55_00,
	}
  
  
	res := CategoriesAvg(payments)
  
  
	if !reflect.DeepEqual(expected, res) {
	  t.Errorf(" got > %v want > nil", res)
	} 
  
  }