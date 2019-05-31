package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

const (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	DPanicLevel = zapcore.DPanicLevel
	PanicLevel  = zapcore.PanicLevel
	FatalLevel  = zapcore.FatalLevel
)

func LevelStr(str string) zapcore.Level {
	switch strings.ToLower(str) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "Fatal":
		return FatalLevel
	}
	return DebugLevel
}

type Config struct {
	Filename   string        // 日志文件路径
	Level      zapcore.Level // 日志等级
	MaxSize    int           // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int           // 日志文件最多保存多少个备份
	MaxAge     int           // 文件最多保存多少天
	Json       bool          // 是否json格式输出
}

var std *zap.SugaredLogger

func init() {
	std = NewLogger("", zapcore.DebugLevel, 0, 0, 0, false)
}

func Set(cfg *Config) {
	std = NewLogger(cfg.Filename, cfg.Level, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge, cfg.Json)
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
