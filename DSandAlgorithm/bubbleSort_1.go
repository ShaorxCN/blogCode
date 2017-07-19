package main

import "log"

func selectMinKey(a []int, start int) int {
	for i := start; i < len(a); i++ {
		if a[start] > a[i] {
			start = i
		}
	}

	return start
}

func main() {
	//冒泡
	s := []int{3, 4, 1, 1, 6, 71, 11, 31, 4}
	log.Println(s)
	change := true
	for i := 0; i < len(s)-1 && change; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			change = false
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				change = true
			}
		}
	}

	log.Println(s)
	//插入
	s1 := []int{3, 4, 1, 1, 6, 71, 11, 31, 4}
	log.Println(s1)
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s1[j] < s1[j-1]; j-- {
			s1[j], s1[j-1] = s1[j-1], s1[j]
		}
	}
	log.Println(s1)

	//选择
	s2 := []int{3, 4, 1, 1, 6, 71, 11, 31, 4}
	log.Println(s2)
	for i := 0; i < len(s2); i++ {
		min := selectMinKey(s2, i)
		if min != i {
			s2[i], s2[min] = s2[min], s2[i]
		}
	}
	log.Println(s2)
}
