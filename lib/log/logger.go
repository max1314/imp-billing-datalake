// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Package log is small wrapper around github.com/sirupsen/logrus.
// It adds a @tags fields to the logging, logstash hook and
// a small set of helper functions.
//
// Remote logging to Logstash is supported. See the Remote and RemoteHostAndPort
// methods.
package log

import (
	"context"
	"io"
	golog "log"
	"os"
	"runtime"
	"sync"

	"imp-billing-datalake/lib/errors"
	"imp-billing-datalake/lib/sharedflags"

	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/tags/logrus"
	httplogrus "github.com/improbable-eng/go-httpwares/logging/logrus/ctxlogrus"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	ErrorKey           = "error"
	ErrorStackKey      = "stack"
	ErrorCodeKey       = "code"
	ErrorCodeStringKey = "code_string"
	TagsField          = "@tags"
	SuccessField       = "success"
)

var (
	// rootLoggers is a list of all Loggers directly constructed with New(). These are used to
	// update the underlying logrus Loggers when flags are updated by cobra. The mutex guards the
	// list as flag updates (dynflags in particular) can happen concurrently. These loggers hold
	// "weak" references to the shared logWriter: when no other references to the logWriter exist,
	// a finalizer will remove the logger with that logWriter from this list.
	rootLoggers     []*Logger
	rootLoggersLock sync.Mutex

	fLogFormat = sharedflags.Set.String(
		"log_format",
		"logfmt",
		"the output format of the logger",
	)
)

// Fields wraps a logrus.Fields.
type Fields logrus.Fields

// Entry wraps a logrus.Entry.
type Entry logrus.Entry

// Formatter wraps a logrus.Formatter.
type Formatter logrus.Formatter

// logWriter defines output options that are shared between all loggers.
//
// nolint: structcheck
// Due to this struct's fields only being used via embedded types, false positives are being
// raised by the `structcheck` linter. See https://github.com/golangci/golangci-lint/issues/572.
type logWriter struct {
	// logrus loggers do not expect to have their properties (notably level) dynamically updated.
	// we wrap access in an rwlock to ensure the dynflag handler can update the level safely.
	mu sync.RWMutex
	*logrus.Logger
	// if set, this logger has remote logging enabled
	remote bool
	// if false, the level of this logger will be updated when the --log_level flag changes
	levelOverride bool
	// if false, the formatter of the logger may be updated when the --log_force_colors flag changes
	formatterOverride bool
}

// logWriterRef is used to provide weak/strong reference semantics.
type logWriterRef struct {
	*logWriter
}

// Logger is used to log entries. You can get a new with logger := New() and then
// call logger.Infof and friends. See github.com/sirupsen/logrus for more information.
type Logger struct {
	writer *logWriterRef
	entry  *logrus.Entry
	Data   Fields
}

// A hook to be fired when logging on the logging levels returned from
// `Levels()` on your implementation of the interface.
type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}

// New returns a new logger with the default fields filled in. It
// takes a list of tag strings that will be used in the @tags field.
// It defaults to logging to standard output in plain text.
func New(tags ...string) *Logger {
	logger := NewWithWriter(os.Stdout, tags...)
	return logger
}

// NewNullLogger returns returns an improbable logger that discards the messages.
func NewNullLogger() *Logger {
	logrusLogger, _ := test.NewNullLogger()
	return newLogger(logrus.NewEntry(logrusLogger))
}

// NewNullLoggerWithHook returns returns an improbable logger that doesn't log to any output and the logrus hook which is useful for testing.
func NewNullLoggerWithHook() (*Logger, *test.Hook) {
	logrusLogger, hook := test.NewNullLogger()
	return newLogger(logrus.NewEntry(logrusLogger)), hook
}

// Returns a new logger whose output goes through the given writer.
func NewWithWriter(w io.Writer, tags ...string) *Logger {
	l := logrus.New()
	lvl := fLogLevel.Get()
	if lvl == "" {
		lvl = os.Getenv("LOG_LEVEL")
	}

	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		l.Level = logrus.InfoLevel
	} else {
		l.Level = level
	}
	setOutput(l, w)

	logger := newLogger(logrus.NewEntry(l))
	if *fRemoteLoggingEnabled {
		logger.enableRemote()
	}

	fields := Fields{
		TagsField: tags,
	}

	switch *fLogFormat {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: DefaultTimestampFormat,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "@timestamp",
			},
		})
		hostname, _ := os.Hostname()
		fields["HOSTNAME"] = hostname
	case "logfmt":
	default:
		logger.Fatalf("unrecognized log format %s", *fLogFormat)
	}

	return logger.WithFields(fields)
}

// FieldsFromContext returns log fields from the grpc/http logger that is currently on the context.
func FieldsFromContext(context context.Context) Fields {
	gRPCFields := grpclogrus.Extract(context).Data
	if len(gRPCFields) > 0 {
		return (Fields)(gRPCFields)
	}

	httpFields := httplogrus.Extract(context).Data
	if len(httpFields) > 0 {
		return (Fields)(httpFields)
	}

	return map[string]interface{}{}
}

// FieldsToContext adds log fields to the grpc/http logger that is currently on the context.
func FieldsToContext(ctx context.Context, fields Fields) {
	logrusFields := logrus.Fields(fields)
	httplogrus.AddFields(ctx, logrusFields)
	grpclogrus.AddFields(ctx, logrusFields)
}

func setOutput(l *logrus.Logger, w io.Writer) {
	l.Out = w
	l.Formatter = &logrus.TextFormatter{DisableColors: true}

	f, isWriterFile := w.(*os.File)
	isTerminal := isWriterFile && isatty.IsTerminal(f.Fd())
	isCygwinTerminal := isWriterFile && isatty.IsCygwinTerminal(f.Fd())

	if isTerminal || isCygwinTerminal  { //|| fForceColorLogs.Get()
		l.Formatter = &humanFormatter{}
	}
	if runtime.GOOS == "windows" && isTerminal {
		// For standard windows terminals, we have to use a special writer that translates the color
		// codes into syscalls that update the terminal state.
		l.Out = termdecoWriter{w: w}
	}
}

func newLogger(logger *logrus.Entry) *Logger {
	writer := &logWriter{
		Logger: logger.Logger,
	}

	// Create a ref to the writer that removes the writer from the root loggers list when the ref is
	// garbage collected.
	strongRef := &logWriterRef{
		logWriter: writer,
	}
	runtime.SetFinalizer(strongRef, func(ref *logWriterRef) {
		rootLoggersLock.Lock()
		defer rootLoggersLock.Unlock()
		// filtering without allocating, see:
		// https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
		newSlice := rootLoggers[:0]
		for _, l := range rootLoggers {
			if l.writer.logWriter == ref.logWriter {
				// drop the rootlogger matching the garbage-collected ref
				l.Close()
				continue
			}
			newSlice = append(newSlice, l)
		}
		rootLoggers = newSlice
	})

	// Put a separate ref to the one above in the root loggers list. This represents a weak ref as
	// it is removed when the strong ref is garbage collected.
	weakRef := &logWriterRef{
		logWriter: writer,
	}
	rootLogger := &Logger{
		writer: weakRef,
		entry:  logger,
	}
	rootLoggersLock.Lock()
	rootLoggers = append(rootLoggers, rootLogger)
	rootLoggersLock.Unlock()

	// Return the strong ref to the user
	return &Logger{
		writer: strongRef,
		entry:  logger,
	}
}

// WithFields wraps logrus.WithFields. Basic usage:
//
//	import "improbable.io/log"
//
//	logger := log.New(...)
//	logger.WithFields(log.Fields{"animal": "walrus"}).Info("A walrus appears")
//
func (l *Logger) WithFields(fields Fields) *Logger {
	entry := l.entry.WithFields(logrus.Fields(fields))
	return &Logger{writer: l.writer, entry: entry}
}

// IsTerminal returns true if the given loggers file descriptor is a terminal.
func (l *Logger) IsTerminal() bool {
	w := l.writer.Out

	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

func (l *Logger) Entry() *logrus.Entry {
	return l.entry
}

// Convert map[string]string to log.Fields
func FieldsFromStringMap(m map[string]string) Fields {
	f := Fields{}
	for k, v := range m {
		f[k] = v
	}
	return f
}

func FieldsFromError(err error) Fields {
	return Fields{
		ErrorKey:      errors.Desc(err),
		ErrorCodeKey:  errors.Code(err),
		ErrorStackKey: errors.Details(err),
		// Use a separate field for string code because the "code" field is shared with http
		// code. ElasticSearch magic means that if the first logged code looks like an int
		// later codes with strings will be dropped.
		ErrorCodeStringKey: errors.Code(err).String(),
	}
}

// WithField wraps logrus.WithField.
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return l.WithFields(Fields{key: value})
}

// WithError wraps logrus.WithError.
func (l *Logger) WithError(err error) *Logger {
	if err == nil {
		return l
	}
	return l.WithFields(FieldsFromError(err))
}

// WithTag appends a tag to the @tags field
func (l *Logger) WithTag(tag string) *Logger {
	var tags []string
	if tagsField, ok := l.entry.Data[TagsField]; ok {
		tags = tagsField.([]string)
	}
	tags = append(tags, tag)
	return l.WithField(TagsField, tags)
}

// SetLevel sets the standard logger level.
func (l *Logger) SetLevel(level Level) {
	l.writer.mu.Lock()
	defer l.writer.mu.Unlock()
	l.writer.Level = logrus.Level(level)
	l.writer.levelOverride = true
}

// SetLevel sets the standard logger level.
func (l *Logger) GetLevel() Level {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	return Level(l.writer.Level)
}

// DEPRECATED: SetOutput sets the standard logger output
func (l *Logger) SetOut(out io.Writer) {
	l.writer.Logger.Out = out
}

// SetFormatter sets the formatter for the logger
func (l *Logger) SetFormatter(formatter Formatter) {
	l.writer.mu.Lock()
	defer l.writer.mu.Unlock()
	l.writer.Logger.Formatter = formatter
	l.writer.formatterOverride = true
}

// AddHook adds a new hook to listen for log events.
func (l *Logger) AddHook(hook Hook) {
	l.writer.Logger.Hooks.Add(&logrusHook{hook: hook})
}

// Close will close any attached LogStashHook, so that entries in its buffer are flushed
func (l *Logger) Close() {
	for _, levelHooks := range l.writer.Logger.Hooks {
		for _, hook := range levelHooks {
			lrh, ok := hook.(*logrusHook)
			if ok {
				h, ok := lrh.hook.(*LogStashHook)
				if ok {
					h.Close()
				}
			}
		}
	}
}

// CopyHooksFrom performs a deep copy of the otherLogger's hooks into this logger's hooks
func (l *Logger) CopyHooksFrom(otherLogger *Logger) {
	for level, hooks := range otherLogger.writer.Logger.Hooks {
		// Deep copy the hooks.
		hooksCopy := make([]logrus.Hook, len(hooks))
		copy(hooksCopy, hooks)
		// Append to existing hooks.
		l.writer.Logger.Hooks[level] = append(l.writer.Logger.Hooks[level], hooksCopy...)
	}
}

func (l *Logger) Copy() *Logger {
	logger := New()
	logger = logger.WithFields(Fields(l.entry.Data))
	logger.CopyHooksFrom(l)
	if l.writer.formatterOverride {
		logger.SetFormatter(l.writer.Logger.Formatter)
	}
	if l.writer.levelOverride {
		logger.SetLevel(l.GetLevel())
	}
	return logger
}

// ToGolangLogger returns a Golang `log` that logs at a Warn level.
func (l *Logger) ToGolangLogger() *golog.Logger {
	return l.ToGolangLoggerWithLevel(WarnLevel)
}

// ToGolangLogger returns a Golang `log` that logs at the given level.
func (l *Logger) ToGolangLoggerWithLevel(level Level) *golog.Logger {
	return golog.New(l.ToWriterWithLevel(level), "", 0)
}

// ToWriterWithLevel returns an io.Writer that logs at the Warn level.
func (l *Logger) ToWriter() io.Writer {
	return l.ToWriterWithLevel(WarnLevel)
}

// ToWriterWithLevel returns an io.Writer that logs at the given level.
func (l *Logger) ToWriterWithLevel(level Level) io.Writer {
	return &golangLoggerWriter{level: level, l: l}
}

// WarnIfErr will log a warning if the given function returns an error.
// It can be used to safely defer functions that return an error without losing their warnings.
// Example:
//     defer logger.WarnIfErr(conn.Close, "Closing connection failed, may cause connection leak")
func (l *Logger) WarnIfErr(f func() error, msg string) {
	if err := f(); err != nil {
		l.WithError(err).Warn(msg)
	}
}

// golangLoggerWriter is needed to use a Writer so that you can get a golang.Logger.
type golangLoggerWriter struct {
	level Level
	l     *Logger
}

func (w *golangLoggerWriter) Write(p []byte) (n int, err error) {
	switch w.level {
	case DebugLevel:
		w.l.Debug(string(p))
	case WarnLevel:
		w.l.Warn(string(p))
	case ErrorLevel:
		w.l.Error(string(p))
	default:
		w.l.Info(string(p))
	}
	return len(p), nil
}

// logrusHook wraps our own hook interface in a logrus hook.
type logrusHook struct {
	hook Hook
}

func (a *logrusHook) Levels() []logrus.Level {
	result := []logrus.Level{}
	for _, level := range a.hook.Levels() {
		result = append(result, logrus.Level(level))
	}
	return result
}

func (a *logrusHook) Fire(logrusEntry *logrus.Entry) error {
	entry := (*Entry)(logrusEntry)
	return a.hook.Fire(entry)
}

func (entry *Entry) String() (string, error) {
	logrusEntry := (*logrus.Entry)(entry)
	return logrusEntry.String()
}
