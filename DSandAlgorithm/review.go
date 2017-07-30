package main

import "log"

func selectMin(a []int, min int) int {
	for i := min; i < len(a); i++ {
		if a[min] > a[i] {
			min = i
		}
	}
	return min
}

func getStv(a []int, low, high int) int {
	//默认第一个元素是基准元素
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
		st := getStv(a, low, high)

		quickSort(a, low, st-1)
		quickSort(a, st+1, high)
	}
}

func main() {
	//bubble
	s := []int{3, 1, 54, 6, 12, 6, 1, 88, 32}
	change := true
	for i := 0; i < len(s)-1 && change; i++ {
		change = false
		for j := 0; j < len(s)-i-1; j++ {

			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				change = true
			}
		}
	}

	log.Println(s)

	//insert
	s1 := []int{3, 1, 54, 6, 12, 6, 1, 88, 32}
	for i := 1; i < len(s1); i++ {
		for j := i; j > 0 && s1[j] < s1[j-1]; j-- {
			s1[j], s1[j-1] = s1[j-1], s1[j]
		}
	}
	log.Println(s1)

	//select
	s2 := []int{3, 1, 54, 6, 12, 6, 1, 88, 32}

	for i := 0; i < len(s2)-1; i++ {
		min := selectMin(s2, i)

		if min != i {
			s2[min], s2[i] = s2[i], s2[min]
		}
	}

	log.Println(s2)

	//quick
	s3 := []int{3, 1, 54, 6, 12, 6, 1, 88, 32}

	low, high := 0, len(s3)-1

	quickSort(s3, low, high)

	log.Println(s3)


}
