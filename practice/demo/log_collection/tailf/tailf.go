package tailf

import (
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/conf"
	"github.com/hpcloud/tail"
	"fmt"
	"os"
	"time"
	"github.com/astaxie/beego/logs"
)

var (
	tailMgr *TailObjMgr
)

type TailObj struct {
	tail *tail.Tail
	conf conf.CollectConf
}

type TailObjMgr struct {
	tails []*TailObj
	msgChan chan *TextMsg
}

type TextMsg struct {
	Msg string
	Topic string
}

func GetLine() (*TextMsg, bool) {
	msg, ok := <- tailMgr.msgChan
	return msg, ok
}

func InitTail(conf []conf.CollectConf, chanSize int) error {
	if len(conf) == 0 {
		panic(fmt.Sprintf("invalid collect conf"))
	}

	tailMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	for _, v := range conf {
		obj := &TailObj{
			conf: v,
		}
		tails, err := getTails(v.LogPath)
		if err != nil {
			return err
		}
		obj.tail = tails
		tailMgr.tails = append(tailMgr.tails, obj)

		go readFromTail(obj)
	}

	logs.Debug("init tail success")
	return nil
}

func getTails(filename string) (*tail.Tail, error) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Sprintf("Open File err: %s", err))
	}
	defer f.Close()
	//if runtime.GOOS == "windows" { // tail 不支持 windows 文件，需要将其中的 \r\n 替换为 \n
	//	buf, err := ioutil.ReadAll(f)
	//	if err != nil {
	//		panic(err)
	//	}
	//	buf = bytes.ReplaceAll(buf, []byte("\r\n"), []byte{'\n'})
	//	f.WriteAt(buf, 0)
	//}

	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		return nil, fmt.Errorf("tail file err: %s", err)
	}
	return tails, nil
}

func readFromTail(obj *TailObj) {
	for {
		line, ok := <- obj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename: %s", obj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textMsg := TextMsg{
			Msg: line.Text,
			Topic: obj.conf.Topic,
		}

		tailMgr.msgChan <- &textMsg
	}
}