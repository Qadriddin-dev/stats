package stats

import (
	"github.com/Qadriddin-dev/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
		  ID:2,
		  Amount:50_00,
		  Category: "Cat",
		  Status: types.StatusOk,
		},
		{
		  ID:1,
		  Amount:52_00,
		  Category: "Cat",
		  Status: types.StatusOk,
		},
		{
		  ID:3,
		  Amount:52_00,
		  Category: "Cat",
		  Status: types.StatusFail,
		},
	  }
	
	  fmt.Println(Avg(payments))
	
	  //Output: 3400
	
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
	  {
		ID:2,
		Amount:53_00,
		Category: "Cafe",
		Status: types.StatusOk,
	  },
	  {
		ID:1,
		Amount:51_00,
		Category: "Cafe",
		Status: types.StatusFail,
	  },
	  {
		ID:3,
		Amount:52_00,
		Category: "Restaurant",
		Status: types.StatusOk,
	  },
	}
  
	fmt.Println(TotalInCategory(payments, "Cafe"))
  
	//Output: 5300
  }