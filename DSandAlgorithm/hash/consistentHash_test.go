package hash

import (
	"testing"
	"strconv"
)

func TestHash(t *testing.T){
	h := NewHashRing(50)
	m := make(map[string]int)
	m["127.0.0.1"] = 2
	m["127.0.0.2"] = 4
	m["127.0.0.3"] = 6
	m["127.0.0.4"] = 4
	m["127.0.0.5"] = 8

	h.AddNodes(m)
	m = make(map[string]int)
	for i:=0;i<100000;i++{
		l := h.GetNode("key"+strconv.Itoa(i))
		if _,ok :=  m[l];ok{
			m[l]++
		}else{
			m[l] = 1
		}
	}
	for k,v := range m{
		t.Logf("%s : %d ",k,v)
	}
}