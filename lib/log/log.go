// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

func (l *Logger) Debug(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Debug(args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Print(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Warn(args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Panic(args...)
}

func (l *Logger) Success(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.WithField(SuccessField, "").Info(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Debugf(format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Printf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Warnf(format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Fatalf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Panicf(format, args...)
}

func (l *Logger) Successf(format string, args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.WithField(SuccessField, "").Infof(format, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Debugln(args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Println(args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Infoln(args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Warnln(args...)
}

func (l *Logger) Warningln(args ...interface{}) {
	l.Warnln(args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Errorln(args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Fatalln(args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.Panicln(args...)
}

func (l *Logger) Successln(args ...interface{}) {
	l.writer.mu.RLock()
	defer l.writer.mu.RUnlock()
	l.entry.WithField(SuccessField, "").Infoln(args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
// Implemented for compatibility with grpclog. Always returns true.
func (l *Logger) V(verbosityLevel int) bool {
	return true
}
