package skiplist

import (
	"math/rand"
	"testing"
)

func TestSkipListGet(t *testing.T) {
	items := ItemSlice([]Item{
		Item{3, "a"},
		Item{5, "a"},
		Item{30, "a"},
		Item{13, "a"},
		Item{8, "a"},
		Item{1, "a"},
		Item{23, "a"},
		Item{6, "a"},
		Item{17, "b"}, // this one's not ike the others!
		Item{11, "a"},
		Item{10, "a"},
		Item{2, "a"},
	})

	rand.Seed(17)
	headNode := New(items, 0.1)
	if headNode.Depth() != 2 {
		t.Fail()
	}

	//headNode.PrintKeys()
	val, err := headNode.Get(17)
	if err != nil {
		t.Error()
	}

	if val.value.(string) != "b" {
		t.Fail()
	}
}
