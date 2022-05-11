package d11

import (
	"fmt"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	c := Constructor()

	tree := c.deserialize("2 1 3")
	ret := c.serialize(tree)
	fmt.Println(ret)
}
