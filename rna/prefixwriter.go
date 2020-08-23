package main

import (
	"bytes"
	"io"
)

type Color string

const (
	Reset Color = "\u001b[0m"

	Red     Color = "\u001b[31m"
	Black   Color = "\u001b[30m"
	Green   Color = "\u001b[32m"
	Yellow  Color = "\u001b[33m"
	Blue    Color = "\u001b[34m"
	Magenta Color = "\u001b[35m"
	Cyan    Color = "\u001b[36m"
	White   Color = "\u001b[37m"
)

func NewPrefixWriter(w io.Writer, prefix string) PrefixWriter {
	return PrefixWriter{
		Color:  White,
		w:      w,
		prefix: []byte(prefix),
	}
}

type PrefixWriter struct {
	Color Color

	w      io.Writer
	prefix []byte
}

func (pw PrefixWriter) Write(p []byte) (int, error) {
	sum := 0

	prefix := pw.prefix
	if pw.Color != White {
		prefix = append([]byte(pw.Color), prefix...)
		prefix = append(prefix, []byte(Reset)...)
	}

	n, err := pw.w.Write(prefix)
	sum += n
	if err != nil {
		return sum, err
	}

	p = bytes.TrimSpace(p)
	p = append(p, []byte("\n")...)
	n, err = pw.w.Write(p)
	sum += n
	return n, err
}
