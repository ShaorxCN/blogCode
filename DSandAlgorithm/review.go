package main

//这里都是基础版本，可以选择整理优化比如选择可以二分，快速可以在low == high的时候再交换基准值和low位置的值等

import "log"


func bubbleSort(a []int){
	log.Printf("before bubble sort:%v\n",a)
	
	todo := make([]int,len(a))
	copy(todo,a)
	
	change := true
	for i := 0;i<len(todo)-1&&change;i++{
		change = false
		for  j:=0;j<len(todo)-i-1;j++{
			if todo[j] > todo[j+1]{
				change = true
				todo[j],todo[j+1] = todo[j+1],todo[j]
			}
		}
	
	}
	
	log.Printf("after bubble sort:%v\n",todo)

}


func selectSort(a []int){
	log.Printf("before select sort:%v\n",a)
	
	todo := make([]int,len(a))
	copy(todo,a)
	
	for i:=0;i<len(todo)-1;i++{
		for j:=i+1;j<len(todo);j++{
			if todo[j]<todo[i]{
				todo[j],todo[i] = todo[i],todo[j]
			}
		}
	}
	
	log.Printf("after select sort:%v\n",todo)
}


func insertSort(a []int){
	log.Printf("before insert sort:%v\n",a)
	
	todo := make([]int,len(a))
	copy(todo,a)
	
	for i:=1;i<len(todo);i++{
		for j:=i;j>0&&todo[j]<todo[j-1];j--{
			todo[j],todo[j-1] = todo[j-1],todo[j]
		}
	}
	
	log.Printf("after insert sort:%v\n",todo)
}


func getStp(todo []int,low,high int)int{
	stv := todo[low]
	
	for low <high{
		for low < high&&todo[high]>=stv{
			high--
		}
		
		
		todo[high],todo[low] = todo[low],todo[high]
		
		
		for low < high&&todo[low]<=stv{
			low++
		}
		
		
		todo[low],todo[high] = todo[high],todo[low]
		
	}
	
	return low
}

func quicksort(todo []int,low,high int){
	if low<high{
		stp := getStp(todo,low,high)
		quicksort(todo,low,stp-1)
		quicksort(todo,stp+1,high)
	
	}
}

func quickSort(a []int){
	log.Printf("before quick sort:%v\n",a)
	
	todo := make([]int,len(a))
	copy(todo,a)
	
	low,high := 0,len(todo)-1
	
	quicksort(todo,low,high)
	
	log.Printf("after quick sort:%v\n",todo)
}


func main(){
	s := []int{2,3,41,3,66,7,11,8,11,4,55}
	
	
	
	bubbleSort(s)
	selectSort(s)
	insertSort(s)
	quickSort(s)
	
}