package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const logTimeLayout = "02-01-2006_03:04:05"

var errOpen = errors.New("couldn't open log file")

// Config defines logger config interface
type Config interface {
	OutDir() string
}

// InitLogger initiate chat-server service logger
func InitLogger(conf Config) (*log.Logger, error) {
	logFileName := fmt.Sprintf("%s/chat_server_%s.log", conf.OutDir(), time.Now().Format(logTimeLayout))

	out, err := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR, 0600) //nolint: gosec
	if err != nil {
		return nil, fmt.Errorf("%w - %v: %w", errOpen, conf.OutDir(), err)
	}

	return log.New(out, "", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile), nil
}
