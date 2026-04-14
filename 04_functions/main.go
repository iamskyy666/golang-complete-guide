package main

import (
	"fmt"
	"strings"
)

// 📂 04_functions
// Proj: Save Math Lib.

type MathError struct{
	Operation string
	InputA int
	InputB int
	Message string

}

const(
division = "Division"
divisionErrMsg = "Division by 0 is not allowed!"
) 

func (e *MathError) Error()string{
	var inputs []string
	if e.Operation == "Division"{
		inputs = append(inputs, fmt.Sprintf("a= %d",e.InputA))
		inputs = append(inputs, fmt.Sprintf("b= %d",e.InputB))
	}
	return fmt.Sprintf("Math Error in %s (%s): %s",e.Operation,strings.Join(inputs,","),e.Message)
}

func Sum(nums ...int)int{
	defer fmt.Println("Sum finished ✅")
	total:=0
	for _,num:= range nums{
		total+=num
	}
	return total
}

func SafeDeivision(a,b int)(int,error){
	defer fmt.Println("Division finished ✅")
	if b==0{
		return 0,&MathError{
			Operation: division,
			InputA: a,
			InputB: b,
			Message: divisionErrMsg,
		}
	}
	return a/b,nil
	
}


func main() {	
	fmt.Println(Sum(1,2,3))
	fmt.Println(SafeDeivision(10,5))
	fmt.Println(SafeDeivision(35,7))
	fmt.Println(SafeDeivision(22,0))

}

// $ go run main.go
// Sum finished ✅
// 6
// Division finished ✅
// 2 <nil>
// Division finished ✅
// 5 <nil>
// Division finished ✅
// 0 Math Error in Division (a= 22,b= 0): Division by 0 is not allowed!


