package faker

import (
	"github.com/lucasjones/reggen"
	"reflect"
	"strconv"
)

// Regex struct
type Regexer interface {
	Gen(v reflect.Value, tag FakerTag) (interface{}, error)
}

var regexer Regexer

type RegexImpl struct {
}

// GetPrice returns a new Money interface of Price struct
func GetRegex() Regexer {
	mu.Lock()
	defer mu.Unlock()

	if regexer == nil {
		regexer = &RegexImpl{}
	}
	return regexer
}

func (r RegexImpl) Gen(v reflect.Value, tag FakerTag) (interface{}, error) {
	len := 64
	if l, ok := tag.Opts["len"]; ok {
		len, _ = strconv.Atoi(l)
	}
	val, err := reggen.Generate(tag.Opts["regex"], len)
	return val, err
}
