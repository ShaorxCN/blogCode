package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"speedDownload/configutil"
	"speedDownload/download"
	"strconv"
	"time"
)

var downloadFile, num string
var goroutines int64

func checkAndPrepare() (err error) {
	c := new(configutil.Config)
	c.Anly("config" + string(os.PathSeparator) + "app.config")

	if c == nil || len(c.Secs) == 0 {
		return errors.New("initial fail, checkout the config and try again")
	}

	downloadFile = c.Secs["properties"]["downloadFile"]
	num = c.Secs["properties"]["goroutines"]

	if downloadFile == "" || num == "" {
		return errors.New("initial fail, checkout the config and try again")
	}

	goroutines, err = strconv.ParseInt(num, 10, 64)
	if err != nil {
		return errors.New("initial fail[goroutines must be digital], checkout the config and try again")
	}

	_, err = os.Stat(downloadFile)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		fmt.Println("directory is not exist,create it")

	}

	err = os.Mkdir(downloadFile, os.ModePerm)

	if err != nil {
		return fmt.Errorf("Create directory fail: %v", err)
	}

	return nil

}

func main() {
	fmt.Println("check the config...")
	err := checkAndPrepare()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("initial success ,donload file:%s,goroutines:%d\n\n", downloadFile, goroutines)

	bio := bufio.NewReader(os.Stdin)
	//var input string
	for {
		fmt.Printf("input the source url(press q/Q to exit):")
		urlbytes, _, err := bio.ReadLine()
		if err != nil {
			fmt.Println("input error and try again")
			continue
		}

		urlstring := string(urlbytes)

		if urlstring == "q" || urlstring == "Q" {
			fmt.Println("bye...")
			time.Sleep(2 * time.Second)
			os.Exit(1)
		}

		fmt.Printf("url is %s\n\nshall we start?(y/Y to start and you can choose another url with anytings other input )", urlstring)

		surebytes, _, err := bio.ReadLine()
		if err != nil {
			fmt.Println("input error and try again from the beginning\n")
			continue
		}

		sure := string(surebytes)

		if sure == "y" || sure == "Y" {
			fmt.Printf("start get from %s", urlstring)
			err := download.Handle(urlstring, downloadFile, goroutines)

			if err != nil {
				fmt.Printf("get fail from url[%s]", urlstring)
				continue
			}

			fmt.Printf("get from url[%s] success,and check the file in[%s]", urlstring, downloadFile)

		}

		continue

	}
}
