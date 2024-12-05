package infrastructure

import (
	"fmt"
	"io"
	"log/slog"
)

type IOAdapter struct {
	w      io.Writer
	logger *slog.Logger
}

func NewIOAdapter(w io.Writer, logger *slog.Logger) *IOAdapter {
	return &IOAdapter{
		w:      w,
		logger: logger,
	}
}

func (a *IOAdapter) Output(content string) {
	_, err := fmt.Fprint(a.w, content)
	if err != nil {
		a.logger.Error(err.Error())
	}
}
