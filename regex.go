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

var regexer Regexer = &RegexImpl{}

// RegexImpl struct
type RegexImpl struct {
}

// Gen returns the fake value the matches the regex
func (r RegexImpl) Gen(v reflect.Value, tag Tag) (interface{}, error) {
	l := 64
	if le, ok := tag.Opts["len"]; ok {
		l, _ = strconv.Atoi(le)
	}
	return reggen.Generate(tag.Opts["regex"], l)
}
