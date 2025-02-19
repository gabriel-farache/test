package logger

import (
	"context"
	"os"
	"testing"

	"github.com/IaC/go-kcloutie/pkg/params/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name          string
		logLevel      string
		cldLog        string
		debugMode     bool
		expectedLevel zapcore.Level
	}{
		{
			name:          "Test with LOG_LEVEL=DEBUG, CLD_LOG=CONSOLE, DebugModeEnabled=false",
			logLevel:      "DEBUG",
			cldLog:        "CONSOLE",
			debugMode:     false,
			expectedLevel: zapcore.DebugLevel,
		},
		{
			name:          "Test with LOG_LEVEL=INFO, CLD_LOG=FILE, DebugModeEnabled=false",
			logLevel:      "INFO",
			cldLog:        "FILE",
			debugMode:     false,
			expectedLevel: zapcore.InfoLevel,
		},
		{
			name:          "Test with LOG_LEVEL=ERROR, CLD_LOG=CONSOLE, DebugModeEnabled=true",
			logLevel:      "ERROR",
			cldLog:        "CONSOLE",
			debugMode:     true,
			expectedLevel: zapcore.ErrorLevel,
		},
		{
			name:          "Test with LOG_LEVEL=INFO, CLD_LOG=FILE, DebugModeEnabled=true",
			logLevel:      "INFO",
			cldLog:        "FILE",
			debugMode:     true,
			expectedLevel: zapcore.DebugLevel,
		},
		{
			name:          "Test with LOG_LEVEL=INVALID, CLD_LOG=CONSOLE, DebugModeEnabled=false",
			logLevel:      "INVALID",
			cldLog:        "CONSOLE",
			debugMode:     false,
			expectedLevel: zapcore.InfoLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("LOG_LEVEL", tt.logLevel)
			os.Setenv("CLD_LOG", tt.cldLog)
			settings.DebugModeEnabled = tt.debugMode

			logger := Get()
			core := logger.Core()
			levelEnabler := core.Enabled(tt.expectedLevel)

			if !levelEnabler {
				t.Errorf("Get() = %v, want %v", levelEnabler, true)
			}
		})
	}
}

func TestFromCtx(t *testing.T) {
	// Create a new logger for testing
	testLogger := zap.NewExample()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			name: "Test with context containing logger",
			args: args{
				ctx: context.WithValue(context.Background(), ctxKey{}, testLogger),
			},
			wantNil: false,
		},
		{
			name: "Test with context not containing logger",
			args: args{
				ctx: context.Background(),
			},
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromCtx(tt.args.ctx)

			if got == nil && !tt.wantNil {
				t.Errorf("FromCtx() nil = %v, want %v", (got == nil), tt.wantNil)
			}
		})
	}
}

func TestWithCtx(t *testing.T) {
	// Create a new logger for testing
	testLogger := zap.NewExample()
	anotherLogger := zap.NewExample()

	type args struct {
		ctx context.Context
		l   *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want *zap.Logger
	}{
		{
			name: "Test with context containing logger",
			args: args{
				ctx: context.WithValue(context.Background(), ctxKey{}, testLogger),
				l:   testLogger,
			},
			want: testLogger,
		},
		{
			name: "Test with context containing different logger",
			args: args{
				ctx: context.WithValue(context.Background(), ctxKey{}, testLogger),
				l:   anotherLogger,
			},
			want: anotherLogger,
		},
		{
			name: "Test with context not containing logger",
			args: args{
				ctx: context.Background(),
				l:   testLogger,
			},
			want: testLogger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtx := WithCtx(tt.args.ctx, tt.args.l)
			got := FromCtx(gotCtx)

			if got != tt.want {
				t.Errorf("WithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeveledLogger_Error(t *testing.T) {
	type fields struct {
		logger *zap.Logger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test with message and no additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
			},
		},
		{
			name: "Test with message and one additional field",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
				},
			},
		},
		{
			name: "Test with message and multiple additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
					"key2", "value2",
					"key3", "value3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observedZapCore, observedLogs := observer.New(zap.InfoLevel)
			observedLogger := zap.New(observedZapCore)

			l := &LeveledLogger{
				logger: observedLogger,
			}
			l.Error(tt.args.msg, tt.args.keysAndValues...)

			if observedLogs.All()[0].Level != zapcore.ErrorLevel {
				t.Errorf("WithCtx() Level = %v, want %v", observedLogs.All()[0].Level, zapcore.ErrorLevel)
			}
			if observedLogs.All()[0].Message != tt.args.msg {
				t.Errorf("WithCtx() Message = %v, want %v", observedLogs.All()[0].Message, tt.args.msg)
			}
		})
	}
}

func TestLeveledLogger_Info(t *testing.T) {
	type fields struct {
		logger *zap.Logger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test with message and no additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
			},
		},
		{
			name: "Test with message and one additional field",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
				},
			},
		},
		{
			name: "Test with message and multiple additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
					"key2", "value2",
					"key3", "value3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observedZapCore, observedLogs := observer.New(zap.InfoLevel)
			observedLogger := zap.New(observedZapCore)

			l := &LeveledLogger{
				logger: observedLogger,
			}
			l.Info(tt.args.msg, tt.args.keysAndValues...)

			if observedLogs.All()[0].Level != zapcore.InfoLevel {
				t.Errorf("WithCtx() Level = %v, want %v", observedLogs.All()[0].Level, zapcore.InfoLevel)
			}
			if observedLogs.All()[0].Message != tt.args.msg {
				t.Errorf("WithCtx() Message = %v, want %v", observedLogs.All()[0].Message, tt.args.msg)
			}
		})
	}
}

func TestLeveledLogger_Debug(t *testing.T) {
	type fields struct {
		logger *zap.Logger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test with message and no additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
			},
		},
		{
			name: "Test with message and one additional field",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
				},
			},
		},
		{
			name: "Test with message and multiple additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
					"key2", "value2",
					"key3", "value3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observedZapCore, observedLogs := observer.New(zap.DebugLevel)
			observedLogger := zap.New(observedZapCore)

			l := &LeveledLogger{
				logger: observedLogger,
			}
			l.Debug(tt.args.msg, tt.args.keysAndValues...)

			if observedLogs.All()[0].Level != zapcore.DebugLevel {
				t.Errorf("WithCtx() Level = %v, want %v", observedLogs.All()[0].Level, zapcore.DebugLevel)
			}
			if observedLogs.All()[0].Message != tt.args.msg {
				t.Errorf("WithCtx() Message = %v, want %v", observedLogs.All()[0].Message, tt.args.msg)
			}
		})
	}
}

func TestLeveledLogger_Warn(t *testing.T) {
	type fields struct {
		logger *zap.Logger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test with message and no additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
			},
		},
		{
			name: "Test with message and one additional field",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
				},
			},
		},
		{
			name: "Test with message and multiple additional fields",
			fields: fields{
				logger: zap.NewExample(),
			},
			args: args{
				msg: "test message",
				keysAndValues: []interface{}{
					"key1", "value1",
					"key2", "value2",
					"key3", "value3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			observedLogger := zap.New(observedZapCore)

			l := NewLeveledLogger(observedLogger)

			l.Warn(tt.args.msg, tt.args.keysAndValues...)

			if observedLogs.All()[0].Level != zapcore.WarnLevel {
				t.Errorf("WithCtx() Level = %v, want %v", observedLogs.All()[0].Level, zapcore.WarnLevel)
			}
			if observedLogs.All()[0].Message != tt.args.msg {
				t.Errorf("WithCtx() Message = %v, want %v", observedLogs.All()[0].Message, tt.args.msg)
			}
		})
	}
}
