package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
	"os"
	"path"
	"io/ioutil"
	"bytes"
)

func main() {
	dir, _ := os.Getwd()
	filename := path.Join(dir, "book/project/tail/my.log")

	// tail 不支持 windows 文件，需要将其中的 \r\n 替换为 \n
	f, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	buf = bytes.ReplaceAll(buf, []byte("\r\n"), []byte{'\n'})
	f.WriteAt(buf, 0)

	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <- tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
