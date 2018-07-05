package dynamic

import(
	"testing"
)

func TestGetStepNum(t *testing.T){
	n := 10
	t.Log(GetStepNum(n))
}



func TestGetStepNumWithClosure(t *testing.T){
	n := 10
	t.Log(GetStepNumWithClosure(n))
}