package faker

import (
	"errors"
	"reflect"
)

func FakeColumnWithType(rt reflect.Type, tag string) (interface{}, error) {
	tags := decodeTags(tag, rt)
	switch {
	case tags.Mapper == SKIP:
		return reflect.Zero(rt), nil
	default:
		zero := reflect.New(rt)
		err := setDataWithTag(zero, tags)
		return zero.Elem().Interface(), err
	}
}

func FakeColumnWithValue(rv reflect.Value, tag string) error {
	if rv.CanAddr() {
		return errors.New("arg1 should be addressable")
	}

	tags := decodeTags(tag, rv.Type())
	switch {
	case tags.Mapper == SKIP:
		return nil
	default:
		return setDataWithTag(rv, tags)
	}
}

func FakeColumnWithTag(v interface{}, tag string) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("pointer type required for argument")
	}

	rv := reflect.ValueOf(v)
	elem := rv.Elem()

	tags := decodeTags(tag, elem.Type())
	switch {
	case tags.Mapper == SKIP:
		return nil
	default:
		return setDataWithTag(rv, tags)
	}
}
