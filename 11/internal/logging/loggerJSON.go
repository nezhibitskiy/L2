package logging

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"
)

type LoggerEx interface {
	WriteErr(err error)
	WriteInfo(message string) error
	MustImplement()
	SaveWriting() chan struct{}
}

type LoggerJSON struct {
	sync.RWMutex
	Logger    log      `json:"logger"`
	Lmessages []string `json:"logger_messages"`
}

type log struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func NewLogging(chStop <-chan struct{}, chSaveModel chan<- struct{}, path string) (chan string, error) {

	chMessages := make(chan string, 1)

	js := NewLoggerJSON(time.Now().Format(TIMEFORMAT))

	go js.listenMessages(chMessages, chStop, chSaveModel, path)

	return chMessages, nil
}

func NewLoggerJSON(startTime string) *LoggerJSON {
	log := log{StartTime: startTime}

	return &LoggerJSON{
		RWMutex:   sync.RWMutex{},
		Logger:    log,
		Lmessages: nil,
	}
}

func (l *LoggerJSON) listenMessages(chMessages <-chan string, chExit <-chan struct{}, chSaveModel chan<- struct{}, path string) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
LOOP:
	for {
		select {
		case m := <-chMessages:
			l.RWMutex.Lock()
			l.Lmessages = append(l.Lmessages, m)
			l.RWMutex.Unlock()
		case <-chExit:
			wg.Done()
			break LOOP
		}
	}
	wg.Wait()
	l.Logger.EndTime = time.Now().Format(TIMEFORMAT)
	res, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, res, 0666)
	if err != nil {
		panic(err)
	}
	chSaveModel <- struct{}{}
	return
}
