# goLog
基于原生go语言的日志包
## 用法
下载包 <br>
`cd $GOPATH/src/github.com/djq8888` <br>
`git clone https://github.com/djq8888/goLog.git` <br>
导入包 <br>
`import "github.com/djq8888/goLog"`
## 功能
+ 程序启动后，在当前文件夹创建log文件夹，并在log文件夹内创建INFO+time.txt，WARN+time.txt，ERROR+time.txt
+ goLog.LogInfo方法：在INFO+time.txt中写入内容
+ goLog.LogWarn方法：在WARN+time.txt&INFO+time.txt中写入内容
+ goLog.LogError方法：在ERROR+time.txt&WARN+time.txt&INFO+time.txt中写入内容
+ 日志格式：时间（时间戳） 等级 文件 包名.函数名 行数 日志内容<br>
样例：`2020-03-11_00_05_23(1583856323300) WARN E:/Go/src/test/handler.go main.testWarn 18 testWarn`
## 特性
+ 异步写入
+ 协程安全