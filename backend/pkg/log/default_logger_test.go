package log

import (
	"os"
	"reflect"
	"testing"

	sm "github.com/cch123/supermonkey"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logHandler = map[string][]reflect.Value{
		zapcore.DebugLevel.String(): {
			reflect.ValueOf(Debug),
			reflect.ValueOf(Debugf),
			reflect.ValueOf(Debugw),
		},
		zapcore.InfoLevel.String(): {
			reflect.ValueOf(Info),
			reflect.ValueOf(Infof),
			reflect.ValueOf(Infow),
		},
		zapcore.WarnLevel.String(): {
			reflect.ValueOf(Warn),
			reflect.ValueOf(Warnf),
			reflect.ValueOf(Warnw),
		},
		zapcore.ErrorLevel.String(): {
			reflect.ValueOf(Error),
			reflect.ValueOf(Errorf),
			reflect.ValueOf(Errorw),
		},
		zapcore.PanicLevel.String(): {
			reflect.ValueOf(Panic),
			reflect.ValueOf(Panicf),
			reflect.ValueOf(Panicw),
		},
		zapcore.FatalLevel.String(): {
			reflect.ValueOf(Fatal),
			reflect.ValueOf(Fatalf),
			reflect.ValueOf(Fatalw),
		},
	}
)

func TestDefaultLogger(t *testing.T) {
	for level, handlers := range logHandler {
		t.Run("test log with level "+level, func(t *testing.T) {
			fws := &fakeWriteSyncer{}
			logger, err := NewLogger(WithLogLevel(level), WithWriteSyncer(fws))
			assert.Nil(t, err, "failed to new logger: ", err)
			defer logger.Close()
			// Reset default logger
			DefaultLogger = logger

			defer func() {
				if level == "panic" {
					r := recover()
					assert.Equal(t, r, "hello")
				}
			}()

			existed := false
			if level == "fatal" {
				fakeExit := func(int) {
					existed = true
				}
				patch := sm.Patch(os.Exit, fakeExit)
				defer patch.Unpatch()
			}

			handlers[0].Call([]reflect.Value{reflect.ValueOf("hello")})
			if level == "fatal" {
				assert.Equal(t, true, existed, "os.Exit was not called")
			}

			assert.Nil(t, logger.Sync(), "failed to sync logger")

			fields := unmarshalLogMessage(t, fws.bytes())
			assert.Equal(t, fields.Level, level, "bad log level ", fields.Level)
			assert.Equal(t, fields.Message, "hello", "bad log message ", fields.Message)

			handlers[1].Call([]reflect.Value{reflect.ValueOf("hello I am %s"), reflect.ValueOf("alex")})
			assert.Nil(t, logger.Sync(), "failed to sync logger")

			fields = unmarshalLogMessage(t, fws.bytes())
			assert.Equal(t, fields.Level, level, "bad log level ", fields.Level)
			assert.Equal(t, fields.Message, "hello I am alex", "bad log message ", fields.Message)

			handlers[2].Call([]reflect.Value{reflect.ValueOf("hello"), reflect.ValueOf(zap.String("name", "alex")), reflect.ValueOf(zap.Int("age", 3))})

			assert.Nil(t, logger.Sync(), "failed to sync logger")

			fields = unmarshalLogMessage(t, fws.bytes())
			assert.Equal(t, fields.Level, level, "bad log level ", fields.Level)
			assert.Equal(t, fields.Message, "hello", "bad log message ", fields.Message)
			assert.Equal(t, fields.Name, "alex", "bad name field ", fields.Name)
			assert.Equal(t, fields.Age, 3, "bad age field ", fields.Age)
		})
	}
}
