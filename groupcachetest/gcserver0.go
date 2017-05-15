package main

import (
	"errors"
	"github.com/golang/groupcache"
	"log"
	"net/http"
)

/**
*   此处不设置peers.单机使用groupcache.启动后访问localhost:8080/gc?key=keyValue即可
 */
func main() {
	group := groupcache.NewGroup("test", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			log.Printf("get %s from slowdb", key)
			if key != "bad" {
				dest.SetString(key)
				return nil
			}
			return errors.New("illegal key")
		}))

	http.HandleFunc("/gc", func(w http.ResponseWriter, r *http.Request) {
		var data []byte
		key := r.URL.Query().Get("key")
		log.Println(key)
		err := group.Get(nil, key, groupcache.AllocatingByteSliceSink(&data))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
