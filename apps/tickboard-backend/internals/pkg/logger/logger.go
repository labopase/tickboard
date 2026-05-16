package logger

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func NewLogger(cfg *Config) (Logger, error) {
	if cfg == nil {
		return nil, errors.New("logger: config is nil")
	}

	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("logger: failed to parse log level: %w", err)
	}
	encoder := buildZapEncoder(cfg.Environment)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)

	var opts []zap.Option

	if cfg.EnableCaller {
		opts = append(opts, zap.AddCaller())
		opts = append(opts, zap.AddCallerSkip(1))
	}

	if cfg.EnableTrace {
		opts = append(opts, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	logger := zap.New(core, opts...)

	return &zapLog{
		logger:        logger,
		sugaredLogger: logger.Sugar(),
	}, nil

}

func (z *zapLog) Sync() error {
	return z.logger.Sync()
}

func (z *zapLog) Debug(msg ...interface{}) {
	z.sugaredLogger.Debug(msg...)
}

func (z *zapLog) Info(msg ...interface{}) {
	z.sugaredLogger.Info(msg...)
}

func (z *zapLog) Warn(msg ...interface{}) {
	z.sugaredLogger.Warn(msg...)
}

func (z *zapLog) Error(msg ...interface{}) {
	z.sugaredLogger.Error(msg...)
}

func (z *zapLog) Fatal(msg ...interface{}) {
	z.sugaredLogger.Fatal(msg...)
}

func (z *zapLog) Debugf(str string, args ...interface{}) {
	z.sugaredLogger.Debugf(str, args...)
}

func (z *zapLog) Infof(str string, args ...interface{}) {
	z.sugaredLogger.Infof(str, args...)
}

func (z *zapLog) Warnf(str string, args ...interface{}) {
	z.sugaredLogger.Warnf(str, args...)
}

func (z *zapLog) Errorf(str string, args ...interface{}) {
	z.sugaredLogger.Errorf(str, args...)
}

func (z *zapLog) Fatalf(str string, args ...interface{}) {
	z.sugaredLogger.Fatalf(str, args...)
}

func (z *zapLog) Debugw(msg string, fields ...Field) {
	z.logger.Debug(msg, z.convertFields(fields)...)
}

func (z *zapLog) Infow(msg string, fields ...Field) {
	z.logger.Info(msg, z.convertFields(fields)...)
}

func (z *zapLog) Warnw(msg string, fields ...Field) {
	z.logger.Warn(msg, z.convertFields(fields)...)
}

func (z *zapLog) Errorw(msg string, fields ...Field) {
	z.logger.Error(msg, z.convertFields(fields)...)
}

func (z *zapLog) Fatalw(msg string, fields ...Field) {
	z.logger.Fatal(msg, z.convertFields(fields)...)
}

func (z *zapLog) DebugCtx(ctx context.Context, msg string, fields ...Field) {
	z.logger.Debug(msg, z.convertFieldsWithContext(ctx, fields)...)
}

func (z *zapLog) InfoCtx(ctx context.Context, msg string, fields ...Field) {
	z.logger.Info(msg, z.convertFieldsWithContext(ctx, fields)...)
}

func (z *zapLog) WarnCtx(ctx context.Context, msg string, fields ...Field) {
	z.logger.Warn(msg, z.convertFieldsWithContext(ctx, fields)...)
}

func (z *zapLog) ErrorCtx(ctx context.Context, msg string, fields ...Field) {
	z.logger.Error(msg, z.convertFieldsWithContext(ctx, fields)...)
}

func (z *zapLog) FatalCtx(ctx context.Context, msg string, fields ...Field) {
	z.logger.Fatal(msg, z.convertFieldsWithContext(ctx, fields)...)
}

func (z *zapLog) convertFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

func (z *zapLog) convertFieldsWithContext(ctx context.Context, fields []Field) []zap.Field {
	zapFields := z.convertFields(fields)

	if traceID := ctx.Value("trace_id"); traceID != nil {
		zapFields = append(zapFields, zap.Any("trace_id", traceID))
	}
	if requestID := ctx.Value("request_id"); requestID != nil {
		zapFields = append(zapFields, zap.Any("request_id", requestID))
	}
	if userID := ctx.Value("span_id"); userID != nil {
		zapFields = append(zapFields, zap.Any("span_id", userID))
	}

	return zapFields
}

func buildZapEncoder(env string) zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if env == constants.EnvTest || env == constants.EnvProduction {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeName = zapcore.FullNameEncoder

		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		encoderConfig.ConsoleSeparator = " | "
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeName = zapcore.FullNameEncoder

		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return encoder
}
