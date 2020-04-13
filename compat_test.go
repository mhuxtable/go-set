package set

import (
	"fmt"
	"testing"

	"github.com/mhuxtable/go-set/genericset/testitems"
)

var NewItem = testitems.NewSetItem

func TestGenericOpsBackwardsCompatibility(t *testing.T) {
	for i, m := range []func(t *testing.T){
		TestIntersectionSet,
		TestUnionSet,
		TestSubtractSet,
	} {
		t.Run(fmt.Sprintf("test case %d", i), m)
	}
}
