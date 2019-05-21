package faker

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

// Enumer defines the interface of enum value generator
type Enumer interface {
	Gen(v reflect.Value, tag Tag) (interface{}, error)
}

var enumer Enumer

// EnumImpl defines the struct
type EnumImpl struct {
}

// GetEnum returns a new Enumer interface of EnumImpl struct
func GetEnum() Enumer {
	mu.Lock()
	defer mu.Unlock()

	if enumer == nil {
		enumer = &EnumImpl{}
	}
	return enumer
}

// Gen generates an enum value specified in the field tag.
func (r EnumImpl) Gen(v reflect.Value, tag Tag) (interface{}, error) {
	enums := tag.Opts["enum"]
	values := strings.SplitN(enums, "/", -1)
	i := rand.Intn(len(values))
	randVal := values[i]

	kind := tag.Type.Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		kind = tag.Type.Elem().Kind()
	}

	switch kind {
	case reflect.String:
		return randVal, nil
	case reflect.Int8:
		return cast.ToInt8E(randVal)
	case reflect.Uint8:
		return cast.ToUint8E(randVal)
	case reflect.Int16:
		return cast.ToInt16E(randVal)
	case reflect.Uint16:
		return cast.ToUint16E(randVal)
	case reflect.Int32:
		return cast.ToInt32E(randVal)
	case reflect.Uint32:
		return cast.ToUint32E(randVal)
	case reflect.Int64:
		return cast.ToInt64E(randVal)
	case reflect.Uint64:
		return cast.ToUint64E(randVal)
	case reflect.Int:
		return cast.ToIntE(randVal)
	case reflect.Uint:
		return cast.ToUintE(randVal)
	case reflect.Bool:
		return cast.ToBoolE(randVal)
	case reflect.Float32:
		return cast.ToFloat32E(randVal)
	case reflect.Float64:
		return cast.ToFloat64E(randVal)
	default:
		return nil, fmt.Errorf("unknown type %v", tag.Type)
	}
}
