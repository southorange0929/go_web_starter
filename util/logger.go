package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var file string
	var line int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}
	level := strings.ToUpper(entry.Level.String())
	content,err := json.Marshal(entry.Data)
	if err != nil {
		return []byte(""),err
	}
	contentString := string(content)
	msg := ""
	// 数据为空时 则不打印
	if contentString == "{}" {
		msg = fmt.Sprintf("%s [%s] [GOID:%d] [%s:%d] %s \n", timestamp, level, getGID(), file, line, entry.Message)
	} else {
		msg = fmt.Sprintf("%s [%s] [GOID:%d] [%s:%d] %s %v\n", timestamp, level, getGID(), file, line, entry.Message, contentString)
	}
	return []byte(msg), nil
}

// 获取当前协程id
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

var Log = log.New()

func LogInit() {
	Log.SetFormatter(&LogFormatter{})
	Log.SetReportCaller(true)
	Log.SetLevel(log.TraceLevel)
}
