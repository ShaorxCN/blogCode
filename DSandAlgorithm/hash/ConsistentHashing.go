package hash

//简单实现一致性hash,看资料上还有增加权重选项
import (
	"errors"
	"fmt"
	"hash/crc32"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	//节点副本数量比例
	DefaultVirtualNum = 30
)

type HashRing struct {
	nodes      nodeRing //这个实际是待映射的环
	virtualNum int
	weights    map[string]int
	RWMux      sync.RWMutex
}

type nodeRing []node

type node struct {
	//节点key,这里可以存放ip一类
	nodeKey string
	//权重
	weight int
	//环上对应的位置
	location uint32
}

func (c nodeRing) Len() int {
	return len(c)
}

func (c nodeRing) Less(i, j int) bool {
	return c[i].location < c[j].location
}

func (c nodeRing) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c nodeRing) Sort() {
	sort.Sort(c)
}

func NewHashRing(v int) *HashRing {
	if v <= 0 {
		v = DefaultVirtualNum
	}

	return &HashRing{
		virtualNum: v,
		weights:    make(map[string]int),
	}

}

func (h *HashRing) generate() {
	totalNum := h.virtualNum * len(h.weights)
	var totalweights int
	for _, v := range h.weights {
		totalweights += v
	}

	for key, v := range h.weights {
		num := int(math.Floor(float64(v) / float64(totalweights) * float64(totalNum)))
		for i := 0; i < num; i++ {
			hash := crc32.ChecksumIEEE([]byte(key + strconv.Itoa(i)))
			n := node{nodeKey: key, weight: v, location: hash}
			h.nodes = append(h.nodes, n)
		}

	}
	h.nodes.Sort()
}

func (h *HashRing) AddNodes(weights map[string]int) error {
	h.RWMux.Lock()
	defer h.RWMux.Unlock()
	for k, v := range weights {
		if _, ok := h.weights[k]; !ok {
			return errors.New(fmt.Sprintf("nodeKey : %s is existed\n", k))
		}

		h.weights[k] = v
	}
	h.generate()
	return nil
}

func (h *HashRing) AddNode(key string, weight int) error {
	h.RWMux.Lock()
	defer h.RWMux.Unlock()
	if _, ok := h.weights[key]; ok {
		return errors.New(fmt.Sprintf("nodeKey : %s is existed\n", key))
	}

	h.weights[key] = weight
	h.generate()
	return nil
}

func (h *HashRing) RemoveNode(key string) {
	h.RWMux.Lock()
	defer h.RWMux.Unlock()

	delete(h.weights, key)
	h.generate()
}

func (h *HashRing) UpdateNode(key string, weight int) error {
	h.RWMux.Lock()
	defer h.RWMux.Unlock()
	if _, ok := h.weights[key]; ok {
		return errors.New(fmt.Sprintf("nodeKey : %s does not  exist\n", key))
	}

	h.weights[key] = weight
	h.generate()
	return nil
}

//这是根据缓存key获取其应归属的node
func (h *HashRing) GetNode(key string) string {
	if len(h.nodes) == 0 {
		return ""
	}

	keyHash := crc32.ChecksumIEEE([]byte(key))

	//这里是二分法查找
	i := sort.Search(len(h.nodes), func(n int) bool { return h.nodes[n].location > keyHash })

	if i == len(h.nodes) {
		i = 0
	}
	return h.nodes[i].nodeKey
}
