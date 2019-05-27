package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

// var textprotoReaderPool sync.Pool

// func newTextprotoReader(br *bufio.Reader) *textproto.Reader {
// 	if v := textprotoReaderPool.Get(); v != nil {
// 		tr := v.(*textproto.Reader)
// 		tr.R = br
// 		return tr
// 	}
// 	return textproto.NewReader(br)
// }

// func putTextprotoReader(r *textproto.Reader) {
// 	r.R = nil
// 	textprotoReaderPool.Put(r)
// }

func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
	s1 := strings.Index(line, " ")
	s2 := strings.Index(line[s1+1:], " ")
	if s1 < 0 || s2 < 0 {
		return
	}
	s2 += s1 + 1
	return line[:s1], line[s1+1 : s2], line[s2+1:], true
}

var ress = []string{`<html><title>Chunk</title>`, res2, res3, res4, res5}
var res2 = `<p style="color:red">message 0</p><br>`
var res3 = `<p style="color:black">message 1</p><br>`
var res4 = `<p style="color:blue">message 2</p><br>`
var res5 = `<p>end</p></html>`

func handle(c net.Conn) {
	defer c.Close()
	rwc := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	lineHead, _, err := rwc.Reader.ReadLine()
	if err != nil {
		log.Println("readline err:", err)
		return
	}

	log.Println(string(lineHead))

	// 读取header 这边先不管
	for {
		line, _, err := rwc.Reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		fmt.Println("read header:", string(line))

	}

	_, _, proto, _ := parseRequestLine(string(lineHead))
	rwc.Writer.Write([]byte(fmt.Sprintf("%s %d %s\r\n", proto, 200, "ok")))
	rwc.Writer.Write([]byte("Content-type: text/html\r\n"))
	rwc.Writer.Write([]byte("Transfer-Encoding: chunked\r\n"))
	rwc.Writer.Write([]byte("\r\n"))

	rwc.Writer.Flush()

	for _, v := range ress {
		rwc.Writer.Write([]byte(fmt.Sprintf("%0x\r\n%s\r\n", len(v), v)))
		rwc.Writer.Flush()
		time.Sleep(1 * time.Second)
	}

	rwc.Writer.Write([]byte(fmt.Sprintf("%0x\r\n\r\n", 0)))
	rwc.Writer.Flush()
}

func main() {

	log.SetFlags(log.Lshortfile)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("listen errr: ", err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("accept connection err :%v\n", err)
			continue
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					buf := make([]byte, 64*1024)
					buf = buf[:runtime.Stack(buf, false)]
					log.Printf("panic in connection  handle, err: %v, stack: %s", r, buf)
				}
			}()

			handle(c)
		}()
	}

}
