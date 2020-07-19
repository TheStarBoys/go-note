package tailf

import (
	"container/list"
	"fmt"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/conf"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"os"
	"sync"
	"time"
)

var (
	tailMgr *TailObjMgr
)

type TailObj struct {
	tail *tail.Tail
	conf conf.CollectConf
	exitChan chan int
}

type TailObjMgr struct {
	objs    list.List
	msgChan chan *TextMsg
	lock 	sync.Mutex
}

func (mgr *TailObjMgr) AppendObj(obj *TailObj) {
	mgr.objs.PushBack(obj)
}

func (mgr *TailObjMgr) DelObj(obj *TailObj) *TailObj {
	for e := mgr.objs.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*TailObj)
		if tmp.conf.LogPath == obj.conf.LogPath && tmp.conf.Topic == obj.conf.Topic {
			mgr.objs.Remove(e)
			return tmp
		}
	}

	return nil
}

func (mgr *TailObjMgr) RangeObjs() (objs []*TailObj) {
	for e := mgr.objs.Front(); e != nil; e = e.Next() {
		objs = append(objs, e.Value.(*TailObj))
	}

	return objs
}

type TextMsg struct {
	Msg string
	Topic string
}

func GetLine() (*TextMsg, bool) {
	msg, ok := <- tailMgr.msgChan
	return msg, ok
}

func UpdateConfig(confs []conf.CollectConf) {
	// 1. 如果 conf 中没有已存在 tailMgr.objs 中的配置，需要删除
	// 2. 如果 conf 中有新增的，则新创建一个任务
	set1 := make(map[string]bool)
	for _, obj := range tailMgr.RangeObjs() {
		set1[obj.conf.LogPath] = true
	}

	set2 := make(map[string]bool)
	for _, conf := range confs {
		set2[conf.LogPath] = true
	}

	tailMgr.lock.Lock()
	defer tailMgr.lock.Unlock()

	for _, conf := range confs {
		if !set1[conf.LogPath] { // 需要新增
			err := createNewTask(conf)
			if err != nil {
				logs.Error("tailf updateConfig err: %v", err)
				continue
			}
		}
	}

	for _, obj := range tailMgr.RangeObjs() {
		if !set2[obj.conf.LogPath] { // 需要删除配置并停止任务
			deletedObj := tailMgr.DelObj(obj)
			close(deletedObj.exitChan)
		}
	}
}

func InitTail(confs []conf.CollectConf, chanSize int) error {
	tailMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	if len(confs) == 0 {
		logs.Error("invalid collect conf")
		return nil
	}

	for _, v := range confs {
		err := createNewTask(v)
		if err != nil {
			return err
		}
	}

	logs.Debug("init tail success")
	return nil
}

func createNewTask(conf conf.CollectConf) error {
	obj := &TailObj{
		conf: conf,
		exitChan: make(chan int),
	}
	tail, err := getTail(conf.LogPath)
	if err != nil {
		return err
	}
	obj.tail = tail
	tailMgr.AppendObj(obj)

	go readFromTail(obj)

	return nil
}

func getTail(filename string) (*tail.Tail, error) {
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
		select {
		case line, ok := <- obj.tail.Lines:
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
		case <- obj.exitChan:
			logs.Warn("tail obj will exited, conf: %#v", obj.conf)
			return
		}
	}
}