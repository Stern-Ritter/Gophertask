package server

import (
	"bytes"
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLoggerInterceptor(t *testing.T) {
	var buf bytes.Buffer
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.DebugLevel,
	))
	interceptor := LoggerInterceptor(logger)

	info := &grpc.UnaryServerInfo{
		FullMethod: "/test.Service/Method",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "response", nil
	}

	t.Run("successful unary method call", func(t *testing.T) {
		buf.Reset()

		resp, err := interceptor(context.Background(), nil, info, handler)
		assert.NoError(t, err, "Expected no error from the handler")
		assert.Equal(t, "response", resp, "Expected handler to return 'response'")

		log := buf.String()
		assert.Contains(t, log, `"level":"info"`, "Expected log message with level 'info'")
		assert.Contains(t, log, `"protocol":"grpc"`, "Expected log message contains protocol name")
		assert.Contains(t, log, `"msg":"finished call"`, "Expected log message indicate finish call")
		assert.Contains(t, log, `"grpc.method_type":"unary"`, "Expected log message contains method type")
		assert.Contains(t, log, `"grpc.service":"test.Service"`, "Expected log message contains service name")
		assert.Contains(t, log, `"grpc.method":"Method"`, "Expected log message contains method name")
		assert.Contains(t, log, `"grpc.code":"OK"`, "Expected log message contains response status code")

		startTimeRegexp := regexp.MustCompile(`"grpc.start_time":".+"`)
		timeMsRegexp := regexp.MustCompile(`"grpc.time_ms":"\d+(\.\d+)?"`)

		assert.Regexp(t, startTimeRegexp, log, "Expected log message contains start time")
		assert.Regexp(t, timeMsRegexp, log, "Expected log message contains time taken")
	})

	t.Run("unary method call with error", func(t *testing.T) {
		buf.Reset()

		handlerWithError := func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, status.Errorf(codes.Internal, "test error")
		}

		resp, err := interceptor(context.Background(), nil, info, handlerWithError)
		assert.Error(t, err, "Expected error from the handler")
		assert.Nil(t, resp, "Expected no response due to error")

		log := buf.String()
		assert.Contains(t, log, `"level":"error"`, "Expected log message with level 'error'")
		assert.Contains(t, log, `"protocol":"grpc"`, "Expected log message contains protocol name")
		assert.Contains(t, log, `"msg":"finished call"`, "Expected log message indicate finish call")
		assert.Contains(t, log, `"grpc.method_type":"unary"`, "Expected log message contains method type")
		assert.Contains(t, log, `"grpc.service":"test.Service"`, "Expected log message contains service name")
		assert.Contains(t, log, `"grpc.method":"Method"`, "Expected log message contains method name")
		assert.Contains(t, log, `"grpc.code":"Internal"`, "Expected log message contains response status code")
		assert.Contains(t, log, `"grpc.error":"rpc error: code = Internal desc = test error"`, "Expected log to contain error message")

		startTimeRegexp := regexp.MustCompile(`"grpc.start_time":".+"`)
		timeMsRegexp := regexp.MustCompile(`"grpc.time_ms":"\d+(\.\d+)?"`)

		assert.Regexp(t, startTimeRegexp, log, "Expected log message contains start time")
		assert.Regexp(t, timeMsRegexp, log, "Expected log message contains time taken")
	})
}
