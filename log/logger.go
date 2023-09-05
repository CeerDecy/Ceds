package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"Ceds/internal/util"
)

type Level int

var Log *Logger

func init() {
	Log = Default()
	Log.WriterInFile("logs/running")
}

// Levels
const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

// LevelFromString returns a Level from a string.
func (l *Level) String() string {
	switch *l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	default:
		return ""
	}
}

type Logger struct {
	Level       Level
	Writers     []*LoggerWriter
	Formatter   LoggerFormatter
	LogFilePath string
	LogFileSize int64
}

type LoggerWriter struct {
	Level Level
	Out   io.Writer
}

// WriterInFile 将日志写入文件
func (logger *Logger) WriterInFile(logPath string) {
	logger.LogFilePath = logPath
	logger.AddWriter(FileWriter(path.Join(logPath, "all.log")), -1)
	logger.AddWriter(FileWriter(path.Join(logPath, "debug.log")), LevelDebug)
	logger.AddWriter(FileWriter(path.Join(logPath, "info.log")), LevelInfo)
	logger.AddWriter(FileWriter(path.Join(logPath, "error.log")), LevelError)
}

// AddWriter 添加日志类型
func (logger *Logger) AddWriter(writer io.Writer, level Level) {
	logger.Writers = append(logger.Writers, &LoggerWriter{
		Level: level,
		Out:   writer,
	})
}

type LoggerFormatter interface {
	Format(param *LoggerFormatterParam) string
}

type LoggerFormatterParam struct {
	Level   Level
	IsColor bool
	Tag     string
	Msg     any
	Fields  map[string]any
}

func Default() *Logger {
	level := LevelDebug
	logger := NewLogger()
	logger.Level = level
	logger.Writers = append(logger.Writers, &LoggerWriter{Level: level, Out: os.Stdout})
	logger.Formatter = &TextFormatter{}
	return logger
}

func (logger *Logger) Print(level Level, tag string, msg any, fields map[string]any) {
	if logger.Level > level {
		// 当前级别大于输入级别，则不打印
		return
	}
	param := &LoggerFormatterParam{
		Level:  level,
		Tag:    tag,
		Msg:    msg,
		Fields: fields,
	}
	param.Level = level
	commonStr := logger.Formatter.Format(param)
	for _, writer := range logger.Writers {
		if writer.Out == os.Stdout {
			param.IsColor = true
			colorStr := logger.Formatter.Format(param)
			_, _ = fmt.Fprintln(writer.Out, colorStr)
		} else if writer.Level == -1 || level == writer.Level {
			logger.CheckFileSize(writer)
			_, _ = fmt.Fprintln(writer.Out, commonStr)
		}
	}
}

// CheckFileSize 检查文件大小
func (logger *Logger) CheckFileSize(writer *LoggerWriter) {
	logFile := writer.Out.(*os.File)
	if logFile != nil {
		stat, err := logFile.Stat()
		if err != nil {
			log.Println(err)
			return
		}
		size := stat.Size()
		if logger.LogFileSize <= 0 {
			logger.LogFileSize = 100 << 20
		}
		if size >= logger.LogFileSize {
			_, file := path.Split(stat.Name())
			index := strings.Index(file, ".")
			if strings.Index(file, " ") != -1 && strings.Index(file, " ") < index {
				index = strings.Index(file, " ")
			}
			filename := file[0:index]
			fmt.Println(filename)
			out := FileWriter(path.Join(logger.LogFilePath, util.JoinStrings(
				filename,
				" ",
				time.Now().Format("06-01-02 15:04:05"),
				".log",
			)))
			writer.Out = out
		}
	}
}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) InfoFields(tag string, msg any, fields map[string]any) {
	logger.Print(LevelInfo, tag, msg, fields)
}

func (logger *Logger) Info(tag string, msg any) {
	logger.Print(LevelInfo, tag, msg, nil)
}

func (logger *Logger) DebugFields(tag string, msg any, fields map[string]any) {
	logger.Print(LevelDebug, tag, msg, fields)
}

func (logger *Logger) Debug(tag string, msg any) {
	logger.Print(LevelDebug, tag, msg, nil)
}

func (logger *Logger) ErrorFields(tag string, msg any, fields map[string]any) {
	logger.Print(LevelError, tag, msg, fields)
}

func (logger *Logger) Error(tag string, msg any) {
	logger.Print(LevelError, tag, msg, nil)
}

func FileWriter(name string) io.Writer {
	writer, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	return writer
}
