package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * json 是否json格式输出
 */
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, json bool) *zap.SugaredLogger {
	//日志文件路径配置2

	var hook *lumberjack.Logger
	if filePath != "" {
		hook = &lumberjack.Logger{
			Filename:   filePath,   // 日志文件路径
			MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: maxBackups, // 日志文件最多保存多少个备份
			MaxAge:     maxAge,     // 文件最多保存多少天
			Compress:   true,       // 是否压缩
		}
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 编码器配置
	var encoder zapcore.Encoder
	if json {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 打印到控制台和文件
	var writer zapcore.WriteSyncer
	if hook == nil {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))

	} else {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook))
	}

	core := zapcore.NewCore(encoder, writer, atomicLevel)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1)).Sugar()
}
