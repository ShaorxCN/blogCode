package main
//简单选择排序
import "log"


func selectMinKey(a []int,start int) int{

	for i := start+1;i<len(a);i++{
		if a[start] > a[i]{
			start = i
		}
	}


	return start
}




func selectSort(a []int){

	for i :=0;i<len(a)-1;i++{
		min := selectMinKey(a,i)


		if min != i{
			a[i] ,a[min]= a[min],a[i]
		}


	}

	log.Println(a)

}

func main(){
	s := []int{3, 7, 1, 9, 11, 0, 3, 13}
	log.Println(s)

	selectSort(s)


}
