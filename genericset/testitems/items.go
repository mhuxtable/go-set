package testitems

import (
	"fmt"
	"math/rand"
)

var setItemFactories []func() interface{}

func init() {
	setItemFactories = makeSetItemFactories()
}

func makeSetItemFactories() []func() interface{} {
	return []func() interface{}{
		func() interface{} { return rand.Int() },
		func() interface{} { return fmt.Sprintf("%d", rand.Int()) },
		func() interface{} {
			type x struct {
				x interface{}
			}

			return x{randomSetItem()}
		}}
}

func randomSetItem() interface{} {
	// We return a variety of primitive types to test the generic set with
	// multiple datatypes in the same tests
	r := rand.Float32()

	// return from one of the factories with psuedo-random uniformity
	for i := 0; i < len(setItemFactories); i++ {
		if r < (1.0 / float32(len(setItemFactories)-i)) {
			return setItemFactories[i]()
		}
	}

	panic("programmer error: this code should be unreachable")
}

func NewSetItem() interface{} {
	return randomSetItem()
}
