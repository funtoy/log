package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

type Config struct {
	Filename   string       // 日志文件路径
	Level      int8         // 日志等级
	MaxSize    int          // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int          // 日志文件最多保存多少个备份
	MaxAge     int          // 文件最多保存多少天
	Json       bool         // 是否json格式输出
	Target     OutputTarget // 输出到哪里
}

var std *zap.SugaredLogger

func init() {
	std = NewLogger("", zapcore.DebugLevel, 0, 0, 0, false, OutputToConsole)
}

func Set(cfg *Config) {
	std = NewLogger(cfg.Filename, zapcore.Level(cfg.Level), cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge, cfg.Json, cfg.Target)
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
