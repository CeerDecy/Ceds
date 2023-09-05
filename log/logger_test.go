package log

import (
	"testing"
)

func TestName(t *testing.T) {
	logger := Default()
	logger.Error("set", "no value")
}
