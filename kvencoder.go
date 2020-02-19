package kvencoder

import (
	"encoding/base64"
	"fmt"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"time"
)

type kvlog struct {
	*zapcore.EncoderConfig
	buf *buffer.Buffer
}

func (k *kvlog) AppendBool(b bool) {
	k.buf.AppendBool(b)
}

func (k *kvlog) AppendByteString(bytes []byte) {
	k.buf.AppendString(string(bytes))
}

func (k *kvlog) AppendComplex128(val complex128) {
	r, i := float64(real(val)), float64(imag(val))
	k.buf.AppendFloat(r, 64)
	k.buf.AppendByte('+')
	k.buf.AppendFloat(i, 64)
	k.buf.AppendByte('i')
}

func (k *kvlog) AppendComplex64(c complex64) {
	k.AppendComplex128(complex128(c))
}

func (k *kvlog) AppendFloat64(f float64) {
	k.buf.AppendFloat(f, 64)
}

func (k *kvlog) AppendFloat32(f float32) {
	k.AppendFloat64(float64(f))
}

func (k *kvlog) AppendInt(i int) {
	k.AppendInt64(int64(i))
}

func (k *kvlog) AppendInt64(i int64) {
	k.buf.AppendInt(i)
}

func (k *kvlog) AppendInt32(i int32) {
	k.AppendInt64(int64(i))
}

func (k *kvlog) AppendInt16(i int16) {
	k.AppendInt64(int64(i))
}

func (k *kvlog) AppendInt8(i int8) {
	k.AppendInt64(int64(i))
}

func (k *kvlog) AppendString(s string) {
	k.buf.AppendString(s)
}

func (k *kvlog) AppendUint(u uint) {
	k.AppendUint64(uint64(u))
}

func (k *kvlog) AppendUint64(u uint64) {
	k.buf.AppendUint(u)
}

func (k *kvlog) AppendUint32(u uint32) {
	k.AppendUint64(uint64(u))
}

func (k *kvlog) AppendUint16(u uint16) {
	k.AppendUint64(uint64(u))
}

func (k *kvlog) AppendUint8(u uint8) {
	k.AppendUint64(uint64(u))
}

func (k *kvlog) AppendUintptr(u uintptr) {
	k.AppendUint64(uint64(u))
}

func (k *kvlog) AppendDuration(val time.Duration) {
	k.EncodeDuration(val, k)
}

func (k *kvlog) AppendObject(marshaler zapcore.ObjectMarshaler) error {
	k.buf.AppendByte('{')
	marshaler.MarshalLogObject(k)
	k.buf.AppendByte('}')
	return nil
}

func (k *kvlog) AppendReflected(value interface{}) error {
	k.buf.AppendString(fmt.Sprint(value))
	return nil
}

func (k *kvlog) AppendArray(marshaler zapcore.ArrayMarshaler) error {
	k.buf.AppendByte('[')
	marshaler.MarshalLogArray(k)
	k.buf.AppendByte(']')
	return nil
}

func (k *kvlog) AppendTime(val time.Time) {
	k.EncodeTime(val, k)
}

func (k *kvlog) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	k.addKey(key)
	return k.AppendArray(marshaler)
}

func (k *kvlog) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	k.addKey(key)
	return k.AppendObject(marshaler)
}

func (k *kvlog) AddBinary(key string, value []byte) {
	k.addKey(key)
	k.buf.AppendString(base64.StdEncoding.EncodeToString(value))
}

func (k *kvlog) AddByteString(key string, value []byte) {
	k.addKey(key)
	k.buf.AppendString(string(value))
}

func (k *kvlog) AddBool(key string, value bool) {
	k.addKey(key)
	k.buf.AppendBool(value)
}

func (k *kvlog) AddComplex128(key string, value complex128) {
	k.addKey(key)
	k.AppendComplex128(value)
}

func (k *kvlog) AddComplex64(key string, value complex64) {
	k.AddComplex128(key, complex128(value))
}

func (k *kvlog) AddDuration(key string, value time.Duration) {
	k.addKey(key)
	k.EncodeDuration(value, k)
}

func (k *kvlog) AddFloat64(key string, value float64) {
	k.addKey(key)
	k.AppendFloat64(value)
}

func (k *kvlog) AddFloat32(key string, value float32) {
	k.AddFloat64(key, float64(value))
}

func (k *kvlog) AddInt(key string, value int) {
	k.AddInt64(key, int64(value))
}

func (k *kvlog) AddInt64(key string, value int64) {
	k.addKey(key)
	k.AppendInt64(value)
}

func (k *kvlog) AddInt32(key string, value int32) {
	k.AddInt64(key, int64(value))
}

func (k *kvlog) AddInt16(key string, value int16) {
	k.AddInt64(key, int64(value))
}

func (k *kvlog) AddInt8(key string, value int8) {
	k.AddInt64(key, int64(value))
}

func (k *kvlog) AddString(key, value string) {
	k.addKey(key)
	k.AppendString(value)
}

func (k *kvlog) AddTime(key string, value time.Time) {
	k.addKey(key)
	k.EncodeTime(value, k)
}

func (k *kvlog) AddUint(key string, value uint) {
	k.AddUint64(key, uint64(value))
}

func (k *kvlog) AddUint64(key string, value uint64) {
	k.addKey(key)
	k.AppendUint64(value)
}

func (k *kvlog) AddUint32(key string, value uint32) {
	k.AddUint64(key, uint64(value))
}

func (k *kvlog) AddUint16(key string, value uint16) {
	k.AddUint64(key, uint64(value))
}

func (k *kvlog) AddUint8(key string, value uint8) {
	k.AddUint64(key, uint64(value))
}

func (k *kvlog) AddUintptr(key string, value uintptr) {
	k.AddUint64(key, uint64(value))
}

func (k *kvlog) AddReflected(key string, value interface{}) error {
	k.addKey(key)
	k.AppendReflected(value)
	return nil
}

func (k *kvlog) OpenNamespace(_ string) {
}

func (k *kvlog) Clone() zapcore.Encoder {
	return k.clone()
}
func (k *kvlog) clone() *kvlog {
	return &kvlog{
		EncoderConfig: k.EncoderConfig,
		buf:           _bufferPool.Get(),
	}
}


func (k *kvlog) addKey(key string) {
	k.addElementSeparator()
	k.buf.AppendString(key)
	k.buf.AppendByte('=')
}

func (k *kvlog) addElementSeparator() {
	k.buf.AppendString("||")
}

func (k *kvlog) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	handler := k.clone()
	// [INFO]
	handler.buf.AppendByte('[')
	handler.EncodeLevel(ent.Level, handler)
	handler.buf.AppendByte(']')
	// [2015-12-02T00:00:07.099+0800]
	handler.buf.AppendByte('[')
	handler.EncodeTime(ent.Time, handler)
	handler.buf.AppendByte(']')
	// [foo.go:123]
	handler.buf.AppendByte('[')
	if handler.CallerKey != "" && ent.Caller.Defined {
		handler.EncodeCaller(ent.Caller, handler)
	}
	handler.buf.AppendByte(']')

	// white space required!!
	handler.buf.AppendByte(' ')

	// default dltag
	handler.AppendString("_undef")
	for i := range fields {
		fields[i].AddTo(handler)
	}
	handler.buf.AppendByte('\n')
	return handler.buf, nil
}

func NewKVEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return &kvlog{
		EncoderConfig: &cfg,
		buf: _bufferPool.Get(),
	}
}

var (
	_bufferPool = buffer.NewPool()
)