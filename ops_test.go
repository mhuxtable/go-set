package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	leftElements, rightElements, expect []interface{}
}

var (
	unionTestCases = []testCase{{
		leftElements:  nil,
		rightElements: nil,
		expect:        nil,
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: nil,
		expect:        []interface{}{1, 2, 3, 4, 5},
	}, {
		leftElements:  nil,
		rightElements: []interface{}{1, 2, 3, 4, 5},
		expect:        []interface{}{1, 2, 3, 4, 5},
	}, {
		leftElements:  []interface{}{},
		rightElements: []interface{}{},
		expect:        []interface{}{},
	}, {
		leftElements:  []interface{}{},
		rightElements: []interface{}{3, 4, 5},
		expect:        []interface{}{3, 4, 5},
	}, {
		leftElements:  []interface{}{1, 2, 3},
		rightElements: []interface{}{},
		expect:        []interface{}{1, 2, 3},
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: []interface{}{1, 2, 3, 4, 5},
		expect:        []interface{}{1, 2, 3, 4, 5},
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: []interface{}{6, 7, 8, 9},
		expect:        []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: []interface{}{4, 5, 6, 7, 8},
		expect:        []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
	}}

	subtractTestCases = []testCase{{
		leftElements:  nil,
		rightElements: nil,
		expect:        nil,
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: nil,
		expect:        []interface{}{1, 2, 3, 4, 5},
	}, {
		leftElements:  nil,
		rightElements: []interface{}{1, 2, 3, 4, 5},
		expect:        nil,
	}, {
		leftElements:  []interface{}{},
		rightElements: []interface{}{},
		expect:        []interface{}{},
	}, {
		leftElements:  []interface{}{1, 2, 3},
		rightElements: []interface{}{},
		expect:        []interface{}{1, 2, 3},
	}, {
		leftElements:  []interface{}{},
		rightElements: []interface{}{4, 5, 6},
		expect:        []interface{}{},
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: []interface{}{6, 7, 8, 9},
		expect:        []interface{}{1, 2, 3, 4, 5},
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: []interface{}{1, 2, 3, 4, 5},
		expect:        []interface{}{},
	}, {
		leftElements:  []interface{}{1, 2, 3, 4, 5},
		rightElements: []interface{}{5, 4, 3, 7, 8, 9},
		expect:        []interface{}{1, 2},
	}}
)

// makeSet initialises and returns a Set with the items provided in xs. If xs
// is nil, the Set is returned uninitialised.
func makeSet(t *testing.T, xs []interface{}) Set {
	var s Set
	if xs != nil {
		s.Add(xs...)
	}
	require.Equal(t, len(xs), s.Count())
	return s
}

func testSetOperation(cases []testCase, op func(s1 *Set, s2 Set)) func(t *testing.T) {
	return func(t *testing.T) {
		for _, tc := range cases {
			s1 := makeSet(t, tc.leftElements)
			s2 := makeSet(t, tc.rightElements)

			op(&s1, s2)

			// s2 is unmodified
			RequireSetElementsMatch(t, tc.rightElements, s2)

			// s1 has been modified appropriately
			RequireSetElementsMatch(t, tc.expect, s1)
		}
	}
}

func TestSetUnion(t *testing.T) {
	testSetOperation(unionTestCases, func(s1 *Set, s2 Set) {
		s1.Union(s2)
	})(t)

}

func TestSetSubtract(t *testing.T) {
	testSetOperation(subtractTestCases, func(s1 *Set, s2 Set) {
		s1.Subtract(s2)
	})(t)
}

func testNonMutatingSetOperation(cases []testCase, op func(s1, s2 Set) Set) func(*testing.T) {
	return func(t *testing.T) {
		for _, tc := range cases {
			s1 := makeSet(t, tc.leftElements)
			s2 := makeSet(t, tc.rightElements)

			out := op(s1, s2)

			// The input sets are not modified
			require.Equal(t, len(tc.leftElements), s1.Count())
			require.ElementsMatch(t, tc.leftElements, s1.Slice())
			require.Equal(t, len(tc.rightElements), s2.Count())
			require.ElementsMatch(t, tc.rightElements, s2.Slice())

			require.Equal(t, len(tc.expect), out.Count())
			require.ElementsMatch(t, tc.expect, out.Slice())
		}
	}
}

func TestUnion(t *testing.T) {
	testNonMutatingSetOperation(unionTestCases, Union)(t)
}

func TestSubtract(t *testing.T) {
	testNonMutatingSetOperation(subtractTestCases, Subtract)(t)
}
