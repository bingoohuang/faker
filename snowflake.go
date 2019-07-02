package faker

import (
	"reflect"

	"github.com/bwmarrin/snowflake"
)

var snow = makeSnowImpl()

func makeSnowImpl() GenV2 {
	impl := &SnowImpl{}

	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	impl.node = node
	return impl
}

// SnowImpl struct
type SnowImpl struct {
	node *snowflake.Node
}

// Gen returns the fake value the matches the regex
func (r SnowImpl) Gen(v reflect.Value, tag Tag) (interface{}, error) {
	format := tag.Opts["format"]
	id := r.node.Generate()
	switch format {
	case "string":
		return id.String(), nil
	case "base64":
		return id.Base64(), nil
	case "base32":
		return id.Base32(), nil
	default:
		return int64(id), nil
	}
}
