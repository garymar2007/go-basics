package stub

import (
	"io"
)

func newCloser() io.Closer {
	return closer{} //@suggestedfix("c", "quickfix", "")
}

type closer struct{}
