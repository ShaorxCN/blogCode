package main

//bubble_sort 正序排序
import (
	"log"
)

func main() {

	s := []int{3, 7, 1, 9, 11, 0, 3, 13}
	log.Println(s)
	change := true //如果某一轮没有进行交换说明已经排序完成，则终止循环
	for i := len(s) - 1; i >= 1 && change; i-- {
		change = false
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				change = true
			}
		}
	}

	log.Println(s)
}
