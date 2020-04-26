package zap

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger("")
	Logger.Debug("Debug")
	Logger.Warn("Warn")
	Logger.Info("Info")
	Logger.Error("Error")
}
