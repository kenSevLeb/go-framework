package log

import (
	"testing"

	"git.yingxiong.com/platform/go-framework/component/trace"
)

func TestMain(m *testing.M) {
	m.Run()
}

// TestInfo 测试Info
func TestInfo(t *testing.T) {
	trace.Set(trace.Get())
	defer trace.GC()
	Info("test", Content{
		"id": 1,
	})
}

// TestInfo 测试Error
func TestError(t *testing.T) {
	trace.Set(trace.Get())
	defer trace.GC()
	Error("test", Content{
		"id": 1,
	})
}

// TestCustom 测试 自定义文件路径
func TestCustom(t *testing.T) {
	err := Init(&Config{
		Switch:     true,
		Mode:       ModeTimeRotate,
		LogPath:    "",
		FileName:   "app.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     28,
		Compress:   false,
	})
	if err != nil {
		t.Fatal(err)
	}
	trace.Set(trace.Get())
	defer trace.GC()
	Info("test", Content{
		"id": "没有自定义路径",
	})
	Info("test", Content{
		"id": "自定义了路径",
	}, "test")
}
