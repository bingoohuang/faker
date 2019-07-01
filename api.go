package faker

import (
	"errors"
	"reflect"
)

func FakeColumnWithTag(v interface{}, tag string) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("pointer type required for argument")
	}

	rv := reflect.ValueOf(v)
	elem := rv.Elem()
	typ := elem.Type()

	tags := decodeTags(tag, typ)
	switch {
	case tags.Mapper == SKIP:
		return nil
	default:
		return setDataWithTag(rv, tags)
	}
}
