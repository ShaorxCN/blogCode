package main

import (
	"errors"
	"github.com/golang/groupcache"
	"log"
	"net/http"
	// "os"
)

var peer_addr = []string{"http://127.0.0.1:8001", "http://127.0.0.1:8002", "http://127.0.0.1:8003"}

func main() {
	//这里注册了groupcache的http服务,供peers之间通信
	peer := groupcache.NewHTTPPool("http://127.0.0.1:8002")
	peer.Set(peer_addr...)
	gc := groupcache.NewGroup("testgroup", 64<<20, groupcache.GetterFunc(
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

	//这里在handler中调用groupcache的Get方法
	http.HandleFunc("/gc", func(w http.ResponseWriter, r *http.Request) {
		var data []byte
		k := r.URL.Query().Get("key")
		err := gc.Get(nil, k, groupcache.AllocatingByteSliceSink(&data))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8002", nil))
}
