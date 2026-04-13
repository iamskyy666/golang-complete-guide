package main

import (
	"fmt"
	"strings"
)

// conditionals-project

var productPrices = map[string]float64{
	"TSHIRT":20.00,
	"JEANS":39.99,
	"BOOKS":15.55,
	"CHAIR":22.50,
}

func calculateItemPrice(itemCode string)(float64,bool){
	basePrice,found:=productPrices[itemCode]

	if !found{
		if before, ok :=strings.CutSuffix(itemCode, "_SALE"); ok {
			originalItemCode:=before
			basePrice,found=productPrices[originalItemCode]
			if found{
				salesPrice:=basePrice * 0.90
				fmt.Printf(" - Item %s (Sale! Original: %.2f , Sale Price: %.2f)\n",
			originalItemCode, basePrice, salesPrice)
			return salesPrice,true
			}
		}

		fmt.Printf("- Item: %s (Product not found!)\n",itemCode)
		return  0.0,false
	}
	return basePrice,true
}

func main() {
	fmt.Println("-------- SIMPLE SALES ORDER PROCESSOR ---------")
	orderItems:=[]string{
		"TSHIRT", "JEANS","BOOKS_SALE","CHAIR_SALE",
	}

	var subTotal float64
	fmt.Println("----------- :Processing Order Items:----------")

	for _,item:=range orderItems{
		price,found:=calculateItemPrice(item)
		if found{
			subTotal+=price
		}
	}

	fmt.Printf("Subtotal Price: %.2f\n",subTotal)
}

// O/P:
// $ go run main.go
// -------- SIMPLE SALES ORDER PROCESSOR ---------
// ----------- :Processing Order Items:----------
//  - Item BOOKS (Sale! Original: 15.55 , Sale Price: 14.00)
//  - Item CHAIR (Sale! Original: 22.50 , Sale Price: 20.25)
// Subtotal Price: 94.23
