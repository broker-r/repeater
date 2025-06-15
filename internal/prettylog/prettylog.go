package prettylog

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"

	"github.com/fatih/color"
)

func PrettyError(err error) slog.Attr {
	return slog.String("error", err.Error())
}

type PrettyHandler struct {
	w    io.Writer
	opts *slog.HandlerOptions
	l    *stdLog.Logger
}

func NewPrettyHandler(w io.Writer, opts *slog.HandlerOptions) PrettyHandler {
	return PrettyHandler{
		w:    w,
		opts: opts,
		l:    stdLog.New(w, "", 0),
	}
}

func (h PrettyHandler) Handle(ctx context.Context, record slog.Record) error {
	log_time := record.Time.Format("2006-01-02 [15:04:05]")
	level := "[" + record.Level.String() + "]"
	message := record.Message

	switch record.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
		message = color.HiMagentaString(message)
	case slog.LevelInfo:
		level = color.BlueString(level)
		message = color.HiBlueString(message)
	case slog.LevelWarn:
		level = color.YellowString(level)
		message = color.HiYellowString(message)
	case slog.LevelError:
		level = color.RedString(level)
		message = color.HiRedString(message)
	}

	attrs := make(map[string]interface{}, record.NumAttrs())

	record.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.Any()
		return true
	})

	var b []byte
	var err error

	if len(attrs) > 0 {
		b, err = json.MarshalIndent(attrs, "", "  ")
		if err != nil {
			return err
		}
	}

	h.l.Println(
		log_time,
		level,
		message,
		color.HiWhiteString(string(b)),
	)

	return nil
}

func (h PrettyHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (h PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return nil
}
func (h PrettyHandler) WithGroup(name string) slog.Handler {
	return nil
}
