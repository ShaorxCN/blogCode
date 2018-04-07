package download

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// 创建并写入文件分块数据
func createAndAppendFile(path string, value string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		return errors.New("create fileList fail")
	}

	_, err = file.WriteString(value)

	return nil

}

// Handle处理下载
func Handle(urlstring, path string, num int64) error {
	res, err := http.Get(urlstring)
	length := res.ContentLength

	if length <= 0 {
		return errors.New("please ensure the url?for download?")
	}
	part := length / num
	var str string
	var start, end int64
	var i int64
	for i = 0; i < num-int64(1); i++ {
		end = start + part
		str += fmt.Sprintf("file%d %d-%d\r\n", i, start, end)
		start = end + 1
	}

	str += fmt.Sprintf("file%d %d-%d(%d)\r\n", num-1, start, length, length)

	err = createAndAppendFile(path+string(os.PathSeparator)+"fileList.txt", str)

	return err
}
