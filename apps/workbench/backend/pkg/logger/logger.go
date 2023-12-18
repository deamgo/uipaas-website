package logger

import (
	"os"
	"sync"

	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	mutex   = &sync.Mutex{}
	hasInit = false
	encoder = zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			CallerKey:      "line",
			NameKey:        "logger",
			FunctionKey:    "func",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.0000"),
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		})
	logger    *zap.Logger
	sugar     *zap.SugaredLogger
	cmdLogger *zap.Logger
	cmdSugar  *CmdSugarLogger
)

// CmdSugarLogger wraps zap.SugaredLogger and zapcore.WriteSyncer in order to use Sugar
// while being able to use low-level writers.
type CmdSugarLogger struct {
	*zap.SugaredLogger
	// wrap ws to print directly
	ws zapcore.WriteSyncer
}

func (log *CmdSugarLogger) Print(s string) {
	_, _ = log.ws.Write([]byte(s))
}

func Init() {
	mutex.Lock()
	defer mutex.Unlock()
	if hasInit {
		return
	}
	hasInit = true

	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	logger = zap.New(core)

	// flushes buffer, if any
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)

	sugar = logger.Sugar()

	// Make sure that logger statements internal to gRPC library are logged using the zapLogger as well.
	grpcZap.ReplaceGrpcLoggerV2(logger)
}

func InitCmdSugar(ws zapcore.WriteSyncer) {
	mutex.Lock()
	defer mutex.Unlock()

	core := zapcore.NewCore(encoder, ws, zap.DebugLevel)
	cmdLogger = zap.New(core)
	defer func(cmdLogger *zap.Logger) {
		err := cmdLogger.Sync()
		if err != nil {
			logger.Error(err.Error())
		}
	}(cmdLogger) // flushes buffer, if any
	cmdSugar = &CmdSugarLogger{
		SugaredLogger: cmdLogger.Sugar(),
		ws:            ws,
	}
}

func Sugar() *zap.SugaredLogger {
	if sugar == nil {
		Init()
	}
	return sugar
}

func Logger() *zap.Logger {
	if logger == nil {
		Init()
	}
	return logger
}

func CmdSugar() *CmdSugarLogger {
	if cmdSugar == nil {
		InitCmdSugar(os.Stdout)
	}
	return cmdSugar
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./installer.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}
