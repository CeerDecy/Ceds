package log

import (
	"fmt"
	"time"
)

type TextFormatter struct {
}

// LevelColor 获取对应等级颜色
func (t *TextFormatter) LevelColor(level Level) string {
	switch level {
	case LevelDebug:
		return yellow
	case LevelInfo:
		return green
	case LevelError:
		return red
	default:
		return white
	}
}

// Format 文本日志格式化
func (t *TextFormatter) Format(param *LoggerFormatterParam) string {
	now := time.Now()
	if param.IsColor {
		return fmt.Sprintf("[Ceds] %v | %s%s%s | [ %s ]:%s",
			now.Format("2006/01/02 - 15:04:05"),
			t.LevelColor(param.Level), param.Level.String(), reset,
			param.Tag, param.Msg,
		)
	} else {
		return fmt.Sprintf("[Ceds] %v | %s | [ %s ]:%s",
			now.Format("2006/01/02 - 15:04:05"),
			param.Level.String(),
			param.Tag, param.Msg,
		)
	}
}
