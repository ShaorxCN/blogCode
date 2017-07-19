package main

//quick sort.
import (
	"log"
	//"sort"
)

//默认第一个值是基准元素
func getSt(a []int, low, high int) int {
	stV := a[low]

	for low < high {
		for low < high && a[high] >= stV {
			high--
		}

		a[low], a[high] = a[high], a[low]

		for low < high && a[low] <= stV {
			low++
		}

		a[low], a[high] = a[high], a[low]
	}

	return low
}

func quickSort(a []int, low, high int) {
	if low < high {
		//获得每次的基准元素的位置
		st := getSt(a, low, high)
		quickSort(a, low, st-1)
		quickSort(a, st+1, high)
	}
}

func main() {
	s := []int{3, 7, 1, 9, 11, 0, 3, 13}
	log.Println(s)
	quickSort(s, 0, len(s)-1)
	//sort.Ints(s)
	log.Println(s)
}
