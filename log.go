package goLog

import (
	"os"
	"sync"
	"time"
)

var infoLog, warnLog, errorLog *secureFile
var infoQueue, warnQueue, errorQueue chan string
var wg sync.WaitGroup
var createInterval time.Duration

//createinterval：切分日志文件的时间间隔（分钟），如果设为0，则会将运行期间所有日志写入同一文件
func Init(createinterval time.Duration) {
	createInterval = createinterval
	//创建日志文件夹
	os.Mkdir("./log",0777)
	//创建队列
	infoQueue = make(chan string)
	warnQueue = make(chan string)
	errorQueue = make(chan string)
	//开启日志写入协程
	go logInfo()
	go logWarn()
	go logError()
}

func Wait() {
	wg.Wait()
}