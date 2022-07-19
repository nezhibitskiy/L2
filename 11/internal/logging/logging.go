package logging

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	TIME       = "TIME"
	MESSAGE    = "MESSAGE"
	LEVEL      = "LEVEL"
	TIMEFORMAT = "2006-02-01 15:04:05"
	INFO       = "INFO"
	ERR        = "ERROR"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

type Logger struct {
	writerINFO  *bufio.Writer
	writerERR   *bufio.Writer
	stopChan    chan struct{}
	messageChan chan string
	saveDone    chan struct{}
}

type loggerError struct {
	string
}

func (l *loggerError) Error() string {
	return l.string
}

func newLoggerErr(s string) *loggerError {
	return &loggerError{s}
}

func NewLogger(writingPath string) (res *Logger, err error) {
	result := Logger{
		writerINFO: bufio.NewWriter(os.Stdout),
		writerERR:  bufio.NewWriter(os.Stderr),
	}
	if writingPath != "n" {
		result.stopChan = make(chan struct{})
		result.saveDone = make(chan struct{})
		result.messageChan, err = NewLogging(result.stopChan, result.saveDone, writingPath)
		if err != nil {
			return nil, err
		}
	}
	res = &result
	return res, nil
}

func (l *Logger) WriteInfo(message string) error {
	mess := fmt.Sprintf("%s[%s] %s[%s] %s[%s]\n", fmt.Sprintf(Cyan+TIME+Reset), time.Now().Format(TIMEFORMAT), fmt.Sprintf(Purple+LEVEL+Reset),
		fmt.Sprintf(Blue+INFO+Reset), fmt.Sprintf(Green+MESSAGE+Reset), message)
	_, err := l.writerINFO.WriteString(mess)
	if err != nil {
		return newLoggerErr(err.Error())
	}

	l.messageChan <- fmt.Sprintf(time.Now().Format(TIMEFORMAT) + " " + fmt.Sprintf("%s[%s]:'%s'", LEVEL, INFO, message))

	err = l.writerINFO.Flush()
	if err != nil {
		return newLoggerErr(err.Error())
	}
	return nil
}

func (l *Logger) WriteErr(err error) {
	mess := fmt.Sprintf("%s[%s] %s[%s] %s[%s]\n", fmt.Sprintf(Cyan+TIME+Reset), time.Now().Format(TIMEFORMAT), fmt.Sprintf(Purple+LEVEL+Reset),
		fmt.Sprintf(Red+ERR+Reset), fmt.Sprintf(Green+MESSAGE+Reset), err.Error())
	_, _ = l.writerERR.WriteString(mess)
	l.messageChan <- fmt.Sprintf(time.Now().Format(TIMEFORMAT) + " " + fmt.Sprintf("%s[%s]:'%s'", LEVEL, ERR, err.Error()))

	err = l.writerERR.Flush()
	if err != nil {

	}

	<-l.SaveWriting()

	os.Exit(1)
}

func (l *Logger) SaveWriting() chan struct{} {
	ret := make(chan struct{})
	l.stopChan <- struct{}{}
	go func() {
		for {
			select {
			case <-l.saveDone:
				ret <- struct{}{}
			}
		}
	}()
	return ret
}

func (l *Logger) MustImplement() {}
