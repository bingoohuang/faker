package faker

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeColumnWithTag(t *testing.T) {
	s := interface{}("")

	assert.Equal(t, reflect.TypeOf(s), reflect.TypeOf(""))

	assert.Nil(t, FakeColumnWithTag(&s, `regex=\d{15}`))
	assert.True(t, regexp.MustCompile(`\d{15}`).MatchString(s.(string)))

	x := interface{}("")
	assert.Nil(t, FakeColumnWithTag(&x, "enum=0/1"))
	assert.True(t, regexp.MustCompile("[01]").MatchString(x.(string)))

	i := 123
	assert.Nil(t, FakeColumnWithTag(&i, "enum=0/1"))
	assert.True(t, i == 0 || i == 1)

	a1, err := FakeColumnWithType(reflect.TypeOf(""), "enum=0/1")
	assert.Nil(t, err)
	assert.True(t, a1 == "0" || a1 == "1")

	a2, err := FakeColumnWithType(reflect.TypeOf(0), "enum=10/11")
	assert.Nil(t, err)
	assert.True(t, a2 == 10 || a2 == 11)
}
