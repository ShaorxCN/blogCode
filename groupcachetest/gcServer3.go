package main

import (
	"errors"
	"github.com/golang/groupcache"
	"log"
	"net/http"
)

var peer_addr = []string{"http://127.0.0.1:8001", "http://127.0.0.1:8002", "http://127.0.0.1:8003"}

func main() {
	peer := groupcache.NewHTTPPool("http://127.0.0.1:8003")
	peer.Set(peer_addr...)
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

	log.Fatal(http.ListenAndServe(":8003", nil))
}
