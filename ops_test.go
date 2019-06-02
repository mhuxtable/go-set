package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetUnion(t *testing.T) {
	t.Run("Union of two uninitialised Sets is a no-op", func(t *testing.T) {
		var s1, s2 Set
		require.Equal(t, 0, s1.Count())
		require.Equal(t, 0, s2.Count())
		s1.Union(s2)
		require.Equal(t, 0, s1.Count())
		require.Equal(t, 0, s2.Count())
	})

	t.Run("Union of an initialised Set with an empty Set is a no-op", func(t *testing.T) {
		var s1, s2 Set
		s1.Add(1, 2, 3, 4, 5)
		require.Equal(t, 5, s1.Count())
		s1.Union(s2)
		require.Equal(t, 5, s1.Count())
	})

	t.Run("Union of various initialised Sets returns the expected result", func(t *testing.T) {
		tcs := []struct {
			leftElements, rightElements, expect []interface{}
		}{{
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

		for _, tc := range tcs {
			var s1, s2 Set

			s1.Add(tc.leftElements...)
			require.Equal(t, len(tc.leftElements), s1.Count())

			s2.Add(tc.rightElements...)
			require.Equal(t, len(tc.rightElements), s2.Count())

			s1.Union(s2)

			// s2 is unmodified
			require.Equal(t, len(tc.rightElements), s2.Count())

			require.Equal(t, len(tc.expect), s1.Count())
			require.ElementsMatch(t, tc.expect, s1.Slice())
		}
	})
}

func TestSetSubtract(t *testing.T) {
	t.Run("subtracting two empty uninitialised is a no-op", func(t *testing.T) {
		var s1, s2 Set
		s1.Subtract(s2)
		require.Equal(t, 0, s1.Count())
		require.Equal(t, 0, s2.Count())
	})

	t.Run("subtracting two sets produces the expected result", func(t *testing.T) {
		tcs := []struct {
			leftElements, rightElements, expect []interface{}
		}{{
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

		for _, tc := range tcs {
			var s1, s2 Set
			s1.Add(tc.leftElements...)
			s2.Add(tc.rightElements...)
			require.Equal(t, len(tc.leftElements), s1.Count())
			require.Equal(t, len(tc.rightElements), s2.Count())

			s1.Subtract(s2)
			require.Equal(t, len(tc.rightElements), s2.Count())
			require.Equal(t, len(tc.expect), s1.Count())

			require.ElementsMatch(t, tc.expect, s1.Slice())
		}
	})
}
