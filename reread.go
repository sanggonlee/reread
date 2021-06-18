package reread

import (
	"bytes"
	"io"
)

// With provides an interface for duplicating a reader.
//
// It returns a new reader which is a concatenation of
// buffer of the original reader read within the given function,
// and  the remaining buffer.
func With(r io.Reader, fn func(r io.Reader) error) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tr := io.TeeReader(r, buf)
	return io.MultiReader(buf, r), fn(tr)
}
