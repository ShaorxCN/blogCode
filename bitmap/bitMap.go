package bitmap

import (
	"fmt"
	"sync"
)

const (
	//表示的最大个数是2^32,可以配合crc32一类的hash法使用，占用空间大约512M 可以修改。
	// 但如果超过uint64最大值,得改下结构,不过内存估计也不允许
	MaxSize = 0x01 << 32
)

type bitMap struct {
	//存储数据
	value []byte
	//bitmap最大容量
	maxSize uint64
	//已经置位的最大值,方便后面的输出
	max uint64

	lock sync.RWMutex
}

//越界错误
type OutOfRange struct {
	message string
}

func (o *OutOfRange) Error() string {
	return o.message
}

func New() *bitMap {
	return NewWithMaxSize(MaxSize)
}

//根据size 返回生成的bitMap
func NewWithMaxSize(maxSize int) *bitMap {
	if maxSize == 0 || maxSize > MaxSize {
		maxSize = MaxSize
	} else {
		//补足为8的倍数以申请byte
		reminder := maxSize % 8
		if reminder != 0 {
			maxSize += 8 - reminder
		}
	}
	return &bitMap{value: make([]byte, maxSize>>3), maxSize: uint64(maxSize)}
}

//非0置1,暂不支持链式
func (b *bitMap) SetValue(offset uint64, value int) error {
	if b.maxSize < offset {
		return &OutOfRange{message: "index out of range"}
	}
	//获取需要set的值的下标以及byte中的第几位
	index, pos := offset/8, offset%8

	//置零
	b.lock.Lock()
	defer b.lock.Unlock()
	if value == 0 {
		b.value[index] &^= 0x01 << pos
	} else {
		b.value[index] |= 0x01 << pos
	}

	if b.max < offset {
		b.max = offset
	}

	return nil
}

func (b *bitMap) GetBit(offset uint64) (uint8, error) {
	if b.maxSize < offset {
		return 0, &OutOfRange{message: "index out of range"}
	}

	//获取需要get的值的下标以及byte中的第几位
	index, pos := offset/8, offset%8
	b.lock.RLock()
	defer b.lock.RUnlock()
	return (b.value[index] >> pos) & 0x01, nil
}

//不过数据量大改为输出最大值以及位置
//todo
func (b *bitMap) String() string {
	return fmt.Sprintf("max : %d;", b.max)
}
