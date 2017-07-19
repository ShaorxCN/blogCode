package main
//直接插入排序
//基本思路选取一个元素作为一个有序的序列，然后将下一个元素作为有序序列插入其中。相同则插在后面.

import "log"

func main(){
	s := []int{3, 7, 1, 9, 11, 0, 3, 13}
	log.Println(s)
	//此处取s[0]作为待插入的序列

	for i:=1;i<len(s);i++{
		change := true
		//此处就直接默认先插到末尾再排序
		for j := i;j>0&&s[j] < s[j-1];j--{
			s[j],s[j-1] = s[j-1],s[j]
		}
	}

	log.Println(s)
}