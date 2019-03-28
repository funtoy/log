package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Level int8

const (
	DebugLevel = Level(iota)
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

const (
	logLevelDebugStr = "DEBUG"
	logLevelInfoStr  = "INFO0"
	logLevelWarnStr  = "WARN0"
	logLevelErrorStr = "ERROR"
	logLevelFatalStr = "FATAL"
)

const (
	callDepth  = 4
	timeFormat = "2006/01/02 15:04:05"
)

var std = NewWithWriter(os.Stdout)

type logger struct {
	out            io.Writer
	path           string
	level          Level
	jsonFormat     bool
	noHighlighting bool
	rotateByDay    bool
	date           string
	chOut          chan []byte
}

type message struct {
	Level   string `json:"LEVEL"`
	Time    string `json:"TIME"`
	Dir     string `json:"DIR"`
	Message string `json:"MSG"`
}

func newLogFileName(path string) string {
	return time.Now().Format(path + ".2006.01.02.log")
}

func getHighlightColor(t string) string {
	switch t {
	case logLevelFatalStr:
		return "[0;31"
	case logLevelErrorStr:
		return "[0;31"
	case logLevelWarnStr:
		return "[0;33"
	case logLevelInfoStr:
		return "[0;34"
	case logLevelDebugStr:
		return "[0;36"
	}
	return "[0;37"
}

func getFileLine(call int) string {
	var file string
	var line int
	var ok bool
	if _, file, line, ok = runtime.Caller(call); !ok {
		file = "???"
		line = 0
	} else {
		if s := strings.Split(file, "/"); len(s) > 0 {
			file = s[len(s)-1]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func New() *logger {
	return NewWithWriter(os.Stdout)
}

func NewWithWriter(out io.Writer) *logger {
	l := &logger{out: out, chOut: make(chan []byte, 10000)}
	go l.handling()
	return l
}

func NewWithFile(pathToFile string) *logger {
	os.MkdirAll(filepath.Dir(pathToFile), 0755)
	f, err := os.OpenFile(newLogFileName(pathToFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	l := NewWithWriter(f)
	l.path = pathToFile
	l.noHighlighting = true
	return l
}

func (r *logger) SetRotateByDay() {
	r.rotateByDay = true
}

func (r *logger) SetLevel(lv Level) {
	r.level = lv
}

func (r *logger) SetJson() {
	r.jsonFormat = true
}

func (r *logger) write(p []byte) {
	r.chOut <- p
}

func (r *logger) handling() {
	defer r.write([]byte("<<<logger dead>>>>:"))

	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			if r.path != "" && r.rotateByDay {
				if f, ok := r.out.(*os.File); ok {
					var name = newLogFileName(r.path)
					if f.Name() != name {
						newOut, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
						if err != nil {
							r.write([]byte("<<<logger error>>>>:" + err.Error()))
						} else {
							r.out = newOut
							f.Close()
						}
					}
				}
			}

		case data := <-r.chOut:
			r.out.Write(data)
		}
	}
}

func (r *logger) doLog(lv, msg string) {
	timeStr := time.Now().Format(timeFormat)
	dir := getFileLine(callDepth)
	if !r.jsonFormat {
		var data string
		if r.noHighlighting {
			data = fmt.Sprintf("[%v] %v | %25s | %v\n", lv, timeStr, dir, msg)
		} else {
			color := getHighlightColor(lv)
			data = fmt.Sprintf("\033%vm[%v] %v | %25s | %v\033[0m\n", color, lv, timeStr, dir, msg)
		}
		r.write([]byte(data))
		return
	}

	var data = &message{
		Level:   lv,
		Time:    time.Now().Format(timeFormat),
		Dir:     dir,
		Message: msg,
	}
	if data, err := json.Marshal(data); err == nil {
		buf := new(bytes.Buffer)
		buf.Write(data)
		buf.WriteByte('\n')
		r.write(buf.Bytes())
	}
}

func (r *logger) Debug(val ...interface{}) {
	if r.level > DebugLevel {
		return
	}
	r.doLog(logLevelDebugStr, fmt.Sprint(val...))
}

func (r *logger) Info(val ...interface{}) {
	if r.level > InfoLevel {
		return
	}
	r.doLog(logLevelInfoStr, fmt.Sprint(val...))
}

func (r *logger) Warn(val ...interface{}) {
	if r.level > WarnLevel {
		return
	}
	r.doLog(logLevelWarnStr, fmt.Sprint(val...))
}

func (r *logger) Error(val ...interface{}) {
	if r.level > ErrorLevel {
		return
	}
	r.doLog(logLevelErrorStr, fmt.Sprint(val...))
}

func (r *logger) Fatal(val ...interface{}) {
	if r.level > FatalLevel {
		return
	}
	r.doLog(logLevelFatalStr, fmt.Sprint(val...))
	os.Exit(1)
}

func (r *logger) Debugf(format string, val ...interface{}) {
	if r.level > DebugLevel {
		return
	}
	r.doLog(logLevelDebugStr, fmt.Sprintf(format, val...))
}

func (r *logger) Infof(format string, val ...interface{}) {
	if r.level > InfoLevel {
		return
	}
	r.doLog(logLevelInfoStr, fmt.Sprintf(format, val...))
}

func (r *logger) Warnf(format string, val ...interface{}) {
	if r.level > WarnLevel {
		return
	}
	r.doLog(logLevelWarnStr, fmt.Sprintf(format, val...))
}

func (r *logger) Errorf(format string, val ...interface{}) {
	if r.level > ErrorLevel {
		return
	}
	r.doLog(logLevelErrorStr, fmt.Sprintf(format, val...))
}

func (r *logger) Fatalf(format string, val ...interface{}) {
	if r.level > FatalLevel {
		return
	}
	r.doLog(logLevelFatalStr, fmt.Sprintf(format, val...))
	os.Exit(1)
}

func Debug(val ...interface{}) {
	std.Debug(val...)
}

func Info(val ...interface{}) {
	std.Info(val...)
}

func Warn(val ...interface{}) {
	std.Warn(val...)
}

func Error(val ...interface{}) {
	std.Error(val...)
}

func Fatal(val ...interface{}) {
	std.Fatal(val...)
}

func Debugf(format string, val ...interface{}) {
	std.Debugf(format, val...)
}

func Infof(format string, val ...interface{}) {
	std.Infof(format, val...)
}

func Warnf(format string, val ...interface{}) {
	std.Warnf(format, val...)
}

func Errorf(format string, val ...interface{}) {
	std.Errorf(format, val...)
}

func Fatalf(format string, val ...interface{}) {
	std.Fatalf(format, val...)
}

func SetJson() {
	std.SetJson()
}

func SetLevel(lv Level) {
	std.SetLevel(lv)
}

func SetLevelByName(lv string) {
	switch strings.ToLower(lv) {
	case "info":
		std.SetLevel(InfoLevel)

	case "warn":
		std.SetLevel(WarnLevel)

	case "error":
		std.SetLevel(ErrorLevel)

	}
}

func SetNoColor() {
	std.noHighlighting = true
}

func SetRotateByDay() {
	std.SetRotateByDay()
}

func WriteToFile(pathToFile string) {
	std = NewWithFile(pathToFile)
}
