package main

//小的放到前面
import (
	"log"
)

func main() {
	s := []int{0, 1, 3, 4, 2, 6, 4, 88, 11, 13}
	change := true
	for i := 0; i < len(s)-1 && change; i++ {
		change = false

		// for j := 0; j < len(s)-i-1; j++ {
		// 	if s[j] > s[j+1] {
		// 		s[j], s[j+1] = s[j+1], s[j]
		// 		change = true
		// 	}
		// }

		for j := len(s) - 1; j > i; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				change = true
			}
		}
	}

	log.Println(s)

}
