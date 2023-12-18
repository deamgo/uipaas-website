package logger

type LogLevel uint32

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelPanic
	LogLevelFatal
)

func log(level LogLevel, args ...interface{}) {
	switch level {
	case LogLevelDebug:
		Sugar().Debug(args...)
	case LogLevelInfo:
		Sugar().Info(args...)
	case LogLevelWarn:
		Sugar().Warn(args...)
	case LogLevelError:
		Sugar().Error(args...)
	case LogLevelPanic:
		Sugar().Panic(args...)
	case LogLevelFatal:
		Sugar().Fatal(args...)
	}
}

func logFormat(level LogLevel, format string, args ...interface{}) {
	switch level {
	case LogLevelDebug:
		Sugar().Debugf(format, args...)
	case LogLevelInfo:
		Sugar().Infof(format, args...)
	case LogLevelWarn:
		Sugar().Warnf(format, args...)
	case LogLevelError:
		Sugar().Errorf(format, args...)
	case LogLevelPanic:
		Sugar().Panicf(format, args...)
	case LogLevelFatal:
		Sugar().Fatalf(format, args...)
	}
}

func Debug(args ...interface{}) {
	log(LogLevelDebug, args)
}

func Debugf(format string, args ...interface{}) {
	logFormat(LogLevelDebug, format, args...)
}

func Info(args ...interface{}) {
	log(LogLevelInfo, args)
}

func Infof(format string, args ...interface{}) {
	logFormat(LogLevelInfo, format, args...)
}

func Warn(args ...interface{}) {
	log(LogLevelWarn, args)
}

func Warnf(format string, args ...interface{}) {
	logFormat(LogLevelWarn, format, args...)
}

func Error(args ...interface{}) {
	log(LogLevelError, args)
}

func Errorf(format string, args ...interface{}) {
	logFormat(LogLevelError, format, args...)
}

func Panic(args ...interface{}) {
	log(LogLevelPanic, args)
}

func Panicf(format string, args ...interface{}) {
	logFormat(LogLevelPanic, format, args)
}

func Fatal(args ...interface{}) {
	log(LogLevelFatal, args)
}

func Fatalf(format string, args ...interface{}) {
	logFormat(LogLevelFatal, format, args)
}
