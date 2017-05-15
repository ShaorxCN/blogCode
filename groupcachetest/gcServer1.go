package main

import (
	"errors"
	"github.com/golang/groupcache"
	"log"
	"net/http"
)

//配合gcServer2,3使用,如果启动的节点和Set的不一样会无法在peers之间通信
var peer_addr = []string{"http://127.0.0.1:8001", "http://127.0.0.1:8002", "http://127.0.0.1:8003"}

func main() {
	//设置HTTPPool，服务注册到DefaultServeMux中
	peer := groupcache.NewHTTPPool(peer_addr[0])
	//创建group,参数分别为groupName,缓存大小,无缓存时获取数据的getter方法
	groupcache.NewGroup("testgroup", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			log.Printf("get data of %s from slowDB\n", key)
			if "1" == key {
				dest.SetString("one")
			} else if "2" == key {
				dest.SetString("two")
			} else {
				return errors.New("illegal key :" + key)
			}

			return nil
		}))

	peer.Set(peer_addr...)
	//这里NewHTTPPool默认注册到defaultServeMux的路由中
	//默认地址是_groupcache/groupName/key,同时这里也是服务的入口
	log.Fatal(http.ListenAndServe(":8001", nil))
}
