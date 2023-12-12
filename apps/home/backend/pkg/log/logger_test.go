package log

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"testing"

	sm "github.com/cch123/supermonkey"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type fakeWriteSyncer struct {
	buf bytes.Buffer
}

type fields struct {
	Level   string
	Time    string
	Message string
	Name    string
	Context string
	Age     int
}

func (fws *fakeWriteSyncer) Sync() error {
	return nil
}

func (fws *fakeWriteSyncer) Write(p []byte) (int, error) {
	return fws.buf.Write(p)
}

func (fws *fakeWriteSyncer) bytes() (p []byte) {
	s := fws.buf.Bytes()
	p = make([]byte, len(s))
	copy(p, s)
	fws.buf.Reset()
	return
}

func unmarshalLogMessage(t *testing.T, data []byte) *fields {
	var f fields
	err := json.Unmarshal(data, &f)
	assert.Nil(t, err, "failed to unmarshal log message: ", err)
	return &f
}

func TestLogger(t *testing.T) {
	for level := range levelMap {
		t.Run("test log with level "+level, func(t *testing.T) {
			fws := &fakeWriteSyncer{}
			logger, err := NewLogger(
				WithLogLevel(level),
				WithWriteSyncer(fws),
				WithContext("test-logger"),
			)
			assert.Nil(t, err, "failed to new logger: ", err)
			defer logger.Close()

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

			rv := reflect.ValueOf(logger)

			handler := rv.MethodByName(http.CanonicalHeaderKey(level))
			handler.Call([]reflect.Value{reflect.ValueOf("hello")})

			if level == "fatal" {
				assert.True(t, existed, "os.Exit was not called")
			}

			assert.Nil(t, logger.Sync(), "failed to sync logger")

			fields := unmarshalLogMessage(t, fws.bytes())
			assert.Equal(t, fields.Level, level, "bad log level ", fields.Level)
			assert.Equal(t, fields.Message, "hello", "bad log message ", fields.Message)
			assert.Equal(t, fields.Context, "test-logger", "bad context")

			handler = rv.MethodByName(http.CanonicalHeaderKey(level) + "f")
			handler.Call([]reflect.Value{reflect.ValueOf("hello I am %s"), reflect.ValueOf("alex")})

			assert.Nil(t, logger.Sync(), "failed to sync logger")

			fields = unmarshalLogMessage(t, fws.bytes())
			assert.Equal(t, fields.Level, level, "bad log level ", fields.Level)
			assert.Equal(t, fields.Message, "hello I am alex", "bad log message ", fields.Message)
			assert.Equal(t, fields.Context, "test-logger", "bad context")

			handler = rv.MethodByName(http.CanonicalHeaderKey(level) + "w")
			handler.Call([]reflect.Value{reflect.ValueOf("hello"), reflect.ValueOf(zap.String("name", "alex")), reflect.ValueOf(zap.Int("age", 3))})

			assert.Nil(t, logger.Sync(), "failed to sync logger")

			fields = unmarshalLogMessage(t, fws.bytes())
			assert.Equal(t, fields.Level, level, "bad log level ", fields.Level)
			assert.Equal(t, fields.Message, "hello", "bad log message ", fields.Message)
			assert.Equal(t, fields.Name, "alex", "bad name field ", fields.Name)
			assert.Equal(t, fields.Age, 3, "bad age field ", fields.Age)
			assert.Equal(t, fields.Context, "test-logger", "bad context")
		})
	}
}

func TestLogLevel(t *testing.T) {
	fws := &fakeWriteSyncer{}
	logger, err := NewLogger(WithLogLevel("error"), WithWriteSyncer(fws))
	assert.Nil(t, err, "failed to new logger: ", err)
	defer logger.Close()

	logger.Warn("this message should be dropped")
	assert.Nil(t, logger.Sync(), "failed to sync logger")

	p := fws.bytes()
	assert.Len(t, p, 0, "saw a message which should be dropped")
}

func TestWithTimeEncoder(t *testing.T) {
	fws := &fakeWriteSyncer{}
	logger, err := NewLogger(WithTimeEncoder("2006-01-02 15:04:00.000"), WithWriteSyncer(fws))
	assert.Nil(t, err, "failed to new logger: ", err)
	defer logger.Close()

	logger.Warn("this message should be dropped")
	assert.Nil(t, logger.Sync(), "failed to sync logger")

	p := fws.bytes()
	fields := unmarshalLogMessage(t, p)
	reg := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}`)
	match := reg.MatchString(fields.Time)
	assert.Equal(t, match, true, "bad log time layout ", fields.Level)
}

func TestZapLogger(t *testing.T) {
	fws := &fakeWriteSyncer{}
	logger, err := NewLogger(WithLogLevel("error"), WithWriteSyncer(fws))
	assert.Nil(t, err, "failed to new logger: ", err)
	defer logger.Close()

	zl := logger.ZapLogger()
	zl.Warn("this message should be dropped")
	assert.Nil(t, logger.Sync(), "failed to sync logger")

	p := fws.bytes()
	assert.Len(t, p, 0, "saw a message which should be dropped")

	zl.Error("this message should be seen")

	p = fws.bytes()
	assert.Contains(t, string(p), "this message should be seen")
}
