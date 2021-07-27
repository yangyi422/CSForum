package logger

import (
	"log"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

// InitLogger 初始化日志
func InitLogger() {
	// 编码器，用于设定生成日志的格式
	var encoder = getEncoder()
	// 设定日志写入文件的信息，例如文件存放位置，分割方式
	var writeSyncer = getLogWriter()
	// 设置日志的编码器，写入位置，日志等级-严重情况由轻至重，ErrorLevel以上的日志输出会附带堆栈信息，FatalLevel还会调用os.Exit退出程序（DebugLevel，InfoLevel，WarnLevel，ErrorLevel，DPanicLevel，PanicLevel，FatalLevel）
	var core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 创建日志对象
	var logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	// 对日志对象进行封装，使其可以使用格式化输出，类似fmt.Printf的输出
	sugarLogger = logger.Sugar()
}

// getEncoder 设置编码器
func getEncoder() zapcore.Encoder {
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
	}
	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}
	// 设置编码格式
	var encoderConf = zapcore.EncoderConfig{
		MessageKey:     "msg",                          // 日志内容对应的key名，此参数必须不为空，否则日志主体不处理
		LevelKey:       "level_name",                   // 日志级别对应的key名
		TimeKey:        "ts",                           // 时间对应的key名
		CallerKey:      "caller_line",                  // 日志输出行数对应的key名
		FunctionKey:    "function",                     // 日志输出方法名对应的key名
		StacktraceKey:  "stacktrace",                   // 栈追踪的key名
		LineEnding:     zapcore.DefaultLineEnding,      // 默认换行，即使不设置
		EncodeLevel:    customLevelEncoder,             // 自定义日志等级格式
		EncodeTime:     customTimeEncoder,              // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 日期转为s
		EncodeCaller:   customCallerEncoder,            // 自定义文件：行号输出项格式
		//EncodeName:     zapcore.FullNameEncoder, //
		//ConsoleSeparator: "",
	}
	return zapcore.NewConsoleEncoder(encoderConf)
}

// 设置日志文件写入信息
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/test.log", // 日志文件路径，默认 os.TempDir()--创建临时目录
		MaxSize:    10,                // 每个日志文件保存10M，默认 100M
		MaxBackups: 5,                 // 保留5个备份，默认不限
		MaxAge:     30,                // 保留30天，默认不限
		Compress:   false,             // 是否压缩，默认不压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 下面方法是针对zap原有方法进行二次封装
// 普通信息输出
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

// 普通信息格式化输出
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

// 警告信息输出
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

// 警告信息格式化输出
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

// 错误信息输出
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

// 错误信息格式化输出
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

// Panic错误信息输出，Panic会导致进程退出
func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

// Panic错误信息格式化输出
func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}
