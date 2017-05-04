package factorial
//Factorial of a number implemented in GoLang - number can be as large as 100 and result should be in exact number

import(
	"strconv"	
)

func multiply(res []int, x int) []int {
	carry := 0
	for i:=0; i<len(res); i++{
		prod := res[i]*x + carry
		res[i] = prod%10
		carry = prod/10
	}
	for carry > 0 {
		res = append(res,carry%10)
		carry = carry/10
	}
	return res
}

func Factorial(n int) (number string) {
	arr := make([]int, 0)
	arr = append(arr, 1)
	for i:=2; i<=n; i++ {
		arr = multiply(arr, i)
	}
	for i := range arr {
		number += strconv.Itoa(arr[len(arr)-i-1])
	}
	return
}
