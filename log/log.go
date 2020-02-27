package log

import (
	"fmt"
	"net/http"
	"os"
)

// Level is an enum of log levels
type Level int

// Values for LogLevel
const (
	Debug Level = 0 + iota
	Info
	Warning
	Error
	Critical
)

// Message is an encapsulated logged message
type Message struct {
	Level  Level
	Format string
	Values []interface{}
}

// Logger is an object that receives messages and forwards them to destinations
type Logger struct {
	BufferSize   int
	destinations map[string]destination
}

type destination struct {
	Name     string
	Receiver MessageReceiver
	Chan     chan<- Message
}

// MessageReceiver is an interface for an object that can receive a Message and do something with it
// It must never block on receiving on the channel
type MessageReceiver interface {
	Receive(<-chan Message)
}

// RegisterReceiver adds a MessageReceiver as a destination for the logger
func (l Logger) RegisterReceiver(name string, recv MessageReceiver) error {
	if _, ok := l.destinations[name]; ok {
		return fmt.Errorf("receiver already registered with name: %s", name)
	}
	ch := make(chan Message, l.BufferSize)
	l.destinations[name] = destination{name, recv, ch}
	go recv.Receive(ch)
	return nil
}

// DeregisterReceiver removes a destination from the logger
func (l Logger) DeregisterReceiver(name string) error {
	if _, ok := l.destinations[name]; !ok {
		return fmt.Errorf("receiver not found: %s", name)
	}
	close(l.destinations[name].Chan)
	delete(l.destinations, name)
	return nil
}

// Debugf prints a debug message to the appropriate logger
func (l Logger) Debugf(format string, values ...interface{}) {
	l.print(Message{Debug, format, values})
}

// Infof prints a debug message to the appropriate logger
func (l Logger) Infof(format string, values ...interface{}) {
	l.print(Message{Info, format, values})
}

// Warningf prints a debug message to the appropriate logger
func (l Logger) Warningf(format string, values ...interface{}) {
	l.print(Message{Warning, format, values})
}

// Errorf prints a debug message to the appropriate logger
func (l Logger) Errorf(format string, values ...interface{}) {
	l.print(Message{Error, format, values})
}

// HTTP logs a received HTTP request
func (l Logger) HTTP(status int, r *http.Request) {
	l.print(Message{Info, "HTTP %d: %s %s referer=%s remote=%s user-agent=%s", []interface{}{
		status,
		r.Method,
		r.URL.String(),
		r.Referer(),
		r.RemoteAddr,
		r.UserAgent(),
	}})
}

// Criticalf prints a debug message to the appropriate logger
func (l Logger) Criticalf(format string, values ...interface{}) {
	l.print(Message{Critical, format, values})
}

func (l Logger) print(m Message) {
	for _, dest := range l.destinations {
		select {
		case dest.Chan <- m:
			// OK
		default:
			fmt.Fprintf(os.Stderr, "[LOG FAILURE] receiver `%s` did not receive message", dest.Name)
		}
	}
}

// ProvideLogger creates a logger for dependency injection
func ProvideLogger() *Logger {
	return &Logger{
		BufferSize:   10,
		destinations: map[string]destination{},
	}
}
