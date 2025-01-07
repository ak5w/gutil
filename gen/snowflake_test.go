package gen_test

import (
	"testing"

	"github.com/ak5w/gutil/gen"
)

func TestSnowflakeId(t *testing.T) {

	var workId int64 = 1
	node, err := gen.NewSnowflakeWorker(workId)
	if err != nil {
		panic(err)
	}
	id1 := node.GetId()
	id2 := node.GetId()
	id3 := node.GetId()

	if !(id3 > id2 && id2 > id1) {
		t.Errorf("id3 <= id2 or id2 <= id1")
	}
}
