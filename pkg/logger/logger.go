package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	config "github.com/MGomed/chat_server/config"
)

const logTimeLayout = "02-01-2006_03:04:05"

var errOpen = errors.New("couldn't open log file")

// InitLogger initiate chat-server service logger
func InitLogger(conf *config.Config) (*log.Logger, error) {
	logFileName := fmt.Sprintf("%s/chat_server_%s.log", conf.OutLogDir, time.Now().Format(logTimeLayout))

	out, err := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR, 0600) //nolint: gosec
	if err != nil {
		return nil, fmt.Errorf("%w - %v: %w", errOpen, conf.OutLogDir, err)
	}

	return log.New(out, "", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile), nil
}
