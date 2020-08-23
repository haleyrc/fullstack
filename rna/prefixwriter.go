package main

import "io"

func NewPrefixWriter(w io.Writer, prefix string) PrefixWriter {
	return PrefixWriter{
		w:      w,
		prefix: []byte(prefix),
	}
}

type PrefixWriter struct {
	w      io.Writer
	prefix []byte
}

func (pw PrefixWriter) Write(p []byte) (int, error) {
	sum := 0
	n, err := pw.w.Write(pw.prefix)
	sum += n
	if err != nil {
		return sum, err
	}
	n, err = pw.w.Write(p)
	sum += n
	return n, err
}
