package  dynamic

import(
	"fmt"
)




func  getStepNum(n int) int{
	if n == 1 || n == 2{
		return n
	}

	return getStepNum(n-1) + getStepNum(n-2)
}


func GetStepNum(n int) (int,error){
	if n <= 0 {
		return 0 ,fmt.Errorf("invalid arg n:%d", n)
	}

	return getStepNum(n),nil
}


func getStepFunc() func(i int) int{
	var d int
	return func(i int) int{
		defer func(){
			d = i
			
		}()
		return d+i
	}
}

func GetStepNumWithClosure(n int)(int ,error){
	if n <= 0 {
		return 0 ,fmt.Errorf("invalid arg n:%d", n)
	}

	num := 1

	f := getStepFunc()

	for i:=0;i<n;i++{
		num = f(num)

	}

	return num,nil
}

