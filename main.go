package main

import (
	"fmt"
	log3 "logger/log"
	"os"
	"path"
	"runtime"
	"time"
)

func getFileInfo() {
	pc, fileName, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("获取失败")
		return
	}
	fucName := runtime.FuncForPC(pc).Name()
	fmt.Println(fucName)
	fmt.Println(path.Base(fileName))
	fmt.Println(line)
}
func main() {
	for {
		log := log3.NewFileLogger("error", "D:/goworkspace/src/logger/", "xc.log", 1024)
		//log := log2.NewFileLog("error")
		log.Debug("这是个debug日志")
		log.Info("这是个info日志")
		log.Warning("这是个warning日志")
		id := 10
		name := "xc"
		log.Error("这是个error日志 id=%d name=%s", id, name)
		log.Fatal("这是个fatal日志")
		time.Sleep(time.Second * 3)
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

}
func f() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed %v", err)
	}
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get open file info failed %v", err)
	}
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Name())
}
