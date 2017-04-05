package spider

import (
	"container/list"
	"log"
	"os"
	"os/signal"

	"sync"
	"time"
)

var (
	spiderPool chan Spider = make(chan Spider, 50)
	stopFalg   chan os.Signal
	initUrl    string
	urls       *queue
)

//spider 方法定义
type Spider interface {
	//设置要爬取的url
	SetUrl(url string)
	//获取html内容
	GetHtml()
	//获取下一层的节点
	GetUrls() []string
	//分析数据
	Analy()
	//保存数据
	SaveData()
}

//简单实现读写线程安全的list,去重交给数据库
type queue struct {
	l   *list.List
	mux sync.RWMutex
}

func (q *queue) PushBack(url string) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.l.PushBack(url)
}

func (q *queue) Get() string {
	q.mux.Lock()
	defer q.mux.Unlock()
	e := q.l.Front()
	return q.l.Remove(e).(string)
}

func (q *queue) Len() int {
	q.mux.RLock()
	defer q.mux.RUnlock()
	return q.l.Len()
}

func AddToPool(s Spider) {

	spiderPool <- s

}

func Start(url string) {

	if spiderPool == nil || len(spiderPool) == 0 {
		log.Fatal("no spider registed !")
	}

	log.Println("start ...")
	urls = &queue{l: list.New()}
	initUrl = url
	log.Println(url)

	//等待接受终止信号
	stopFalg = make(chan os.Signal, 1)
	signal.Notify(stopFalg, os.Interrupt, os.Kill)

	urls.PushBack(initUrl)
rowlooper:
	for {
		//防止知乎防御攻击禁止访问，在公司的网络比较卡，1s一次没有被禁止，回家后1s的间隔就会forbidden.具体几秒有待测试
		time.Sleep(time.Second * 2)
		select {
		case <-stopFalg:
			log.Println("end")
			break rowlooper

		default:

			if urls.Len() != 0 {
				s := <-spiderPool
				url := urls.Get()
				go spiderProcess(s, url)
			} else {
				if len(spiderPool) != 50 {
					log.Println("暂时没有待爬取的url,请等待运行中的worker结束")
				} else {
					log.Println("无待爬取的节点,运行结束")
					break rowlooper
				}
			}

		}
	}

	Stop()

}

func spiderProcess(s Spider, url string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			spiderPool <- s
			return
		}

	}()

	s.SetUrl(url)
	s.GetHtml()
	s.Analy()
	for _, v := range s.GetUrls() {
		urls.PushBack(v)
	}
	s.SaveData()

	spiderPool <- s
}

func Stop() {
	for {
		time.Sleep(time.Second * 1)
		done := len(spiderPool)
		if done != 50 {
			log.Printf("等待剩余goroutine运行结束,共%d,运行中:%d", 50, 50-done)
			continue
		} else {

			break
		}
	}

	log.Println("运行结束~")
	os.Exit(0)
}
