package faker

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestFakeColumnWithTag(t *testing.T) {
	var s string
	assert.Nil(t, FakeColumnWithTag(&s, "regex=\\d{15}"))
	assert.True(t, regexp.MustCompile("\\d{15}").MatchString(s))

	var s2 string
	assert.Nil(t, FakeColumnWithTag(&s2, "enum=x/y/z"))
	assert.True(t, regexp.MustCompile("[xyz]").MatchString(s2))

}
