package log

import (
	"context"
	"log/slog"
	"os"
	"runtime"

	"github.com/pkg/errors"
)

// NewJSONHandler -> generates json
// TextHandler -> generates text form, logfmt standard compliance

type Logger struct {
	slog.Handler
}

func NewLogger() *Logger {
	return &Logger{
		Handler: slog.NewJSONHandler(os.Stdout, nil),
	}
}

func (l *Logger) Info(msg string) {
	slog.New(l.Handler).Info(msg)
}

func (l *Logger) InfoFromCtx(ctx context.Context) {
	slog.New(l.Handler).InfoContext(ctx, "invo from context", "mak kau", "hijau")
}

func (l *Logger) Error(err error) {
	slog.New(l.Handler).Error("error mang", "err", err.Error())
}

type StackFrame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func (l *Logger) ErrorStack(err error) {
	var stack []StackFrame

	var st stackTracer

	// error interface only has Error() which does not has stack trace. errors.Wrap() 
	// returns *withStack, which embeds *stack. The *stack implements StackTrace() which 
	// makes *stack stasifies stackTracer interface. We use errors.As instead of direct type 
	// assertion so it walks full error chain, in case outermost wrapper does not implement
	//  stackTracer but an inner error does (error returned by service in this case)
	if errors.As(err, &st) {
		for _, f := range st.StackTrace() {
			// Use runtime.FuncForPC to get file/line
			pc := uintptr(f) - 1
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				file, line := fn.FileLine(pc)
				stack = append(stack, StackFrame{
					Function: fn.Name(),
					File:     file,
					Line:     line,
				})
			}
		}
	}

	slog.New(l.Handler).Error(err.Error(), "stack_trace", stack)
}
