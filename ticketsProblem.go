package main

import "fmt"

func main(){
	cases := 0
	bills25 := 0
	bills50 := 0
	bills100 := 0
	fmt.Println("Enter number of cases")
	fmt.Scanln(&cases)
	fmt.Println("Enter bills")
	billsCount := make([]int, cases)
	for i := 0; i < cases; i++ {
		if billsCount[i] == 25 {
			bills25 := bills25 + 1 
		}
		else if billsCount[i] == 50 {
			if bills25 > 0 {
				bills25 := bills25 - 1
				bills50 := bills50 + 1
			}
			else {
				return "NO"
			}
		}
		else if billsCount[i] == 100 {
			if bills25 > 0 && bills50 > 0 {
				bills25 := bills25 - 1
				bills50 := bills50 - 1
				bills100 := bills100 + 1
			}
			else if(bills25 > 2) {
				bills25 := bills25 - 3
				bills100 := bills100 + 1
			}
			else {
				return "NO"
			}
		}
	}

	return "YES"
}