// 文档：https://www.godoc.org/go.uber.org/zap

package zap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strconv"
	"time"
)

var Logger *zap.Logger

func init() {
	Logger = zap.NewExample()
}

func InitLogger(filename string) {
	ll := &lumberjack.Logger{
		Filename:   filename, // 输出文件
		MaxSize:    1024,     // 日志文件最大大小（单位：MB）
		MaxBackups: 10,       // 保留的旧日志文件最大数量
		MaxAge:     10,       // 保存日期
	}

	fileWriter := zapcore.AddSync(ll)
	consoleWriter := zapcore.Lock(os.Stdout)

	// 日志输出流可以添加多个
	//bothWriter = zapcore.NewMultiWriteSyncer(zapcore.Lock(os.Stdout), ll)

	fileEncoder := zapcore.NewJSONEncoder(NewEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(NewEncoderConfig())

	highEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, highEnabler),
		zapcore.NewCore(consoleEncoder, consoleWriter, lowEnabler),
	)

	Logger = zap.New(core, zap.AddCaller())
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",                          // json时时间键
		LevelKey:       "level",                       // json时日志等级键
		NameKey:        "logger",                      // json时日志记录器键
		CallerKey:      "caller",                      // json时日志文件信息键
		MessageKey:     "msg",                         // json时日志消息键
		StacktraceKey:  "stacktrace",                  // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,     // 日志换行符
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 日志大小写
		EncodeTime:     zapcore.EpochNanosTimeEncoder, // 日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 日志文件信息
	}
}

func EncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func EncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(runtime.FuncForPC(caller.PC).Name() + ":" + strconv.FormatInt(int64(caller.Line), 10))
}
