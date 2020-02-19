package kvencoder

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestKVEncoder(t *testing.T) {
	should := require.New(t)
	cfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	buf := bytes.NewBuffer(nil)
	l := zap.New(zapcore.NewCore(NewKVEncoder(cfg), zapcore.AddSync(buf), zap.InfoLevel))
	type User struct {
		Name string
		Age int
	}
	u := User{
		Name: "deen",
		Age: 26,
	}
	l.Info("hello", zap.Duration("du", 2*time.Second), zap.Int("val", 105), zap.Bool("ok", false), zap.ByteString("str", []byte("content")), zap.Reflect("ref", u))
	should.Equal("", buf.String())
}