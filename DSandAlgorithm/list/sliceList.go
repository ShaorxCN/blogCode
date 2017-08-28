package main

import "log"

func main() {

	listA := []int{2, 3, 3, 5, 1, 6, 77, 1}
	listB := listA[1:5]
	log.Println(listA, listB)
	//删除一个元素，但是基于同一个底层数组的切片会受影响
	//listA = append(listA[:3],listA[4:]...)
	//log.Println(listA,listB)

	//insert 将22插入到第四个数字后面
	log.Println(listA, listB)
	insertElement := 22
	listTemp := append([]int{}, listA[4:]...)

	listA = append(listA[:4], insertElement)
	log.Println(len(listA), cap(listA))
	listA = append(listA, listTemp...)
	//此处底层数组不够用，重新开辟并将值复制,所以在下面修改listA的时候没有影响到listB
	log.Println(listA, listB, cap(listA))

	listA[0] = 1
	log.Println(listA, listB)

	log.Println(33, 1)

}
