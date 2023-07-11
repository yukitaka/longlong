package cli

import "io"

type IOStream struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}
