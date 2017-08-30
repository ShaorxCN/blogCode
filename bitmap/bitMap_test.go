package bitmap

import (
	"log"
	"math/rand"
	"sort"
	"testing"
)

//测试去重
func TestBitMap(t *testing.T) {
	b := NewWithMaxSize(100)

	//这里仅仅是测试，这些数据可以有其他的方法去重
	origin := make([]int, 100)
	answer := make([]int, 0)
	result := make([]int, 0)

	for i := 0; i < 100; i++ {
		origin[i] = rand.Intn(100)
	}
	log.Println(origin)
	sort.Ints(origin)
	log.Println(origin)
	for i := 0; i < 100; i++ {
		if i > 0 && origin[i-1] == origin[i] {
			continue
		}

		answer = append(answer, origin[i])
	}

	log.Println(answer)

	for i := 1; i < 100; i++ {
		v, err := b.GetBit(uint64(origin[i]))
		if err != nil {
			log.Fatal(err)
		}
		if v == 0 {
			b.SetValue(uint64(origin[i]), 1)
			result = append(result, origin[i])
		}
		continue
	}

	log.Println(result)

	if len(result) != len(answer) {
		t.Errorf("answer is %v,result is %v ", answer, result)
	}

	if (len(answer) != len(result)) || ((answer == nil) != (result == nil)) {
		t.Errorf("answer is %v,result is %v ", answer, result)
	}

	for k, v := range result {
		if v != answer[k] {
			t.Errorf("answer is %v,result is %v ", answer, result)
			return
		}
	}
}
