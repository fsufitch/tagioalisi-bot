package log

import (
	"fmt"
	"io"
)

type writerMessageReceiver struct {
	Writer   io.Writer
	MinLevel Level
	MaxLevel Level
}

func (w writerMessageReceiver) Receive(messages <-chan Message) {
	go func() {
		for message := range messages {
			w.write(message)
		}
	}()
}

func (w writerMessageReceiver) write(m Message) {
	if m.Level < w.MinLevel || m.Level > w.MaxLevel {
		return
	}

	var levelPrefix string
	switch m.Level {
	case Debug:
		levelPrefix = "[DEBUG]"
	case Info:
		levelPrefix = "[INFO]"
	case Warning:
		levelPrefix = "[WARNING]"
	case Error:
		levelPrefix = "[ERROR]"
	case Critical:
		levelPrefix = "[CRITICAL]"
	}

	message := fmt.Sprintf("%s %s\n", levelPrefix, fmt.Sprintf(m.Format, m.Values...))
	w.Writer.Write([]byte(message))
}
