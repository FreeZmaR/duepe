package fxutils

import (
	"log/slog"
)

type Finalizer interface {
	Append(item FinalizerItem)
	Close()
}

type finalizer struct {
	items []FinalizerItem
}

func NewFinalizer() Finalizer {
	return &finalizer{}
}

func (f *finalizer) Append(item FinalizerItem) {
	f.items = append(f.items, item)
}

func (f *finalizer) Close() {
	for _, item := range f.items {
		err := item.Close()
		if err != nil {
			slog.Error("Finalizer error: ", slog.String("err", err.Error()))
		}
	}
}
