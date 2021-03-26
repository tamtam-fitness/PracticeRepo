package practice

import "fmt"

func CheckNumber() {
	for {
		val := 0
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		if val == 0 {
			fmt.Println("val is 0")
		} else if val < 0 {
			panic(val)
		} else {
			fmt.Println("You entered:", val)
		}
	}

}
