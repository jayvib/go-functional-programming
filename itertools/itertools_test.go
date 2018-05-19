package itertools

import (
	"testing"
	"reflect"

	"github.com/yanatan16/itertools"
)

func testIterEq(t *testing.T, it1, it2 Iter) {
	t.Log("Start")

	for el1 := range it1 {
		if el2, ok := <-it2; !ok {
			t.Error("it2 is shorter than it1 ", el1)
			return
		} else if !reflect.DeepEqual(el1, el2){
			t.Error("it is not equal with it2")
			return
		} else {
			t.Log(el1, el2)
		}
	}

	if el2, ok := <-it2; ok {
		t.Error("t1 is shorter than t2 ", el2)
	}
	t.Log("Stop")
}

func TestMap(t *testing.T) {
	mapper := func(item interface{}) interface{} {
		return len(item.(string))
	}

	testIterEq(t, New(3, 5, 10), Iter(itertools.Map(mapper, itertools.Iter(New("CRV", "IS250", "Highlander")))))
}