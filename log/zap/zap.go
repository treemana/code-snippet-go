package zap

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogFile    string // 日志文件路径, 不传则不写文件
	LogLevel   string // debug | info | warn | error
	MaxAge     int    // 保存的天数, 默认不删除
	MaxSize    int    // 单个文件大小 MB
	MaxBackups int    // 最多保留的备份数
	Compress   bool   // 是否压缩
	JsonFormat bool   // 是否用 json 格式
}

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func Init(config Config) {

	var ws zapcore.WriteSyncer
	if len(config.LogFile) > 0 {
		hook := lumberjack.Logger{
			Filename:   config.LogFile, // 日志文件路径
			MaxSize:    config.MaxSize, // megabytes
			MaxAge:     config.MaxAge,
			MaxBackups: config.MaxBackups, // 最多保留300个备份
			LocalTime:  false,
			Compress:   config.Compress, // 是否压缩 disabled by default
		}
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	var enc zapcore.Encoder
	if config.JsonFormat {
		enc = zapcore.NewJSONEncoder(cfg)
	} else {
		enc = zapcore.NewConsoleEncoder(cfg)
	}

	var level zapcore.Level
	switch config.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	Logger = zap.New(zapcore.NewCore(enc, ws, level), zap.AddCaller())
	Sugar = Logger.Sugar()
}
