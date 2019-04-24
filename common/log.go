package common

import (
	"github.com/pkg/errors"
)

// LogLevel denotes the level of a log; higher is more serious
type LogLevel int

// LogDebug, LogInfo, LogWarning, and LogError are valid values for LogLevel
const (
	LogDebug = LogLevel(iota)
	LogInfo
	LogWarning
	LogError
)

// LogEntry is a simple container for a log entry
type LogEntry struct {
	Level   LogLevel
	Message string
}

// LogDispatcher is a wrapper around log.Logger for a more configurable logging system
type LogDispatcher struct {
	listeners []chan LogEntry
}

// AddListener registers a channel to listen for log events
func (l *LogDispatcher) AddListener(listenChan chan LogEntry) {
	l.listeners = append(l.listeners, listenChan)
}

// RemoveListener deregisters a channel listening for log events
func (l *LogDispatcher) RemoveListener(listenChan chan LogEntry) error {
	loggerPosition := -1
	for i := range l.listeners {
		if l.listeners[i] == listenChan {
			loggerPosition = i
			break
		}
	}

	if loggerPosition == -1 {
		return errors.New("Listener not found")
	}

	l.listeners = append(l.listeners[:loggerPosition], l.listeners[loggerPosition+1:]...)
	return nil
}

// Log sends the given message with the given level to all listeners
func (l *LogDispatcher) Log(level LogLevel, message string) {
	entry := LogEntry{Level: level, Message: message}
	for _, listener := range l.listeners {
		go func(c chan LogEntry) {
			c <- entry
		}(listener)
	}
}

// Debug is a shorthand for sending a debug message
func (l *LogDispatcher) Debug(message string) {
	l.Log(LogDebug, message)
}

// Info is a shorthand for sending an info message
func (l *LogDispatcher) Info(message string) {
	l.Log(LogInfo, message)
}

// Warn is a shorthand for sending a warning message
func (l *LogDispatcher) Warn(message string) {
	l.Log(LogWarning, message)
}

// Error is a shorthand for sending an error message
func (l *LogDispatcher) Error(message string) {
	l.Log(LogError, message)
}

// NewLogDispatcher creates a new logger module
func NewLogDispatcher() *LogDispatcher {
	return &LogDispatcher{
		listeners: []chan LogEntry{},
	}
}
