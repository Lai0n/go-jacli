package arg

import (
	"fmt"
	"github.com/lai0n/go-jacli/pkg/iterator"
)

type CommandArg struct {
	fmt.Stringer

	value       *string
	hyphenCount uint8
}

func (a *CommandArg) String() string {
	return (*(a.value))[a.hyphenCount:]
}

func (a *CommandArg) IsFlag() bool {
	return a.hyphenCount > 0
}

func (a *CommandArg) HyphenCount() uint8 {
	return a.hyphenCount
}

func newArg(str *string) *CommandArg {
	hyphenCount := uint8(0)
	for _, char := range *str {
		if char == '-' {
			hyphenCount += 1
			continue
		}
		break
	}

	return &CommandArg{
		value:       str,
		hyphenCount: hyphenCount,
	}
}

// argsIterator iterates an array of string and producing CommandArg as a result.
type argsIterator struct {
	iterator.Iterator[*CommandArg]

	Args []string
	// Constructor should set value to -1 indicating that iteration did not start
	iteratorIndex int
}

func (a *argsIterator) HasNext() bool {
	l := len(a.Args)
	return l > 0 && a.iteratorIndex+1 < l
}

func (a *argsIterator) Next() (elem *CommandArg, ok bool) {
	if a.HasNext() {
		a.iteratorIndex += 1
		return newArg(&a.Args[a.iteratorIndex]), true
	}
	return nil, false
}

func (a *argsIterator) Peek() (elem *CommandArg, ok bool) {
	if a.HasNext() {
		return newArg(&a.Args[a.iteratorIndex+1]), true
	}
	return nil, false
}

func NewArgIterator(argSet []string) iterator.Iterator[*CommandArg] {
	return &argsIterator{
		Args:          argSet,
		iteratorIndex: -1,
	}
}
