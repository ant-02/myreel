package errno

import (
	"fmt"
	"runtime"
)

type Frame uintptr

type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	//nolint
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := Frame(pc)
				_, _ = fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

const (
	depth = 32
	skip  = 3
)

func callers() *stack {
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
