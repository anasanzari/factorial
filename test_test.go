package factorial

import (
	"testing"
	"strings"
	"fmt"
)

func TestSmall(t *testing.T){
	answer := Factorial(5)
	fmt.Println("Factorial calculated for 5 :", answer);
	if strings.Compare(answer, "120") != 0{
		t.Fail()
	}
}

func TestBig(t *testing.T){
	answer := Factorial(100)
	fmt.Println("Factorial calculated for 100 :", answer);
	if strings.Compare(answer, "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000") != 0{
		t.Fail()
	}
}