package faker

import (
	"reflect"
	"strconv"

	"github.com/lucasjones/reggen"
)

// Regexer interface
type Regexer interface {
	Gen(v reflect.Value, tag Tag) (interface{}, error)
}

var regexer Regexer

// RegexImpl struct
type RegexImpl struct {
}

// GetRegex returns a new Regexer interface of RegexImpl struct
func GetRegex() Regexer {
	mu.Lock()
	defer mu.Unlock()

	if regexer == nil {
		regexer = &RegexImpl{}
	}
	return regexer
}

// Gen returns the fake value the matches the regex
func (r RegexImpl) Gen(v reflect.Value, tag Tag) (interface{}, error) {
	len := 64
	if l, ok := tag.Opts["len"]; ok {
		len, _ = strconv.Atoi(l)
	}
	return reggen.Generate(tag.Opts["regex"], len)
}
