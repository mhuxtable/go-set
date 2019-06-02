package set

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestSetNew(t *testing.T) {
	t.Run("New creates an empty set", func(t *testing.T) {
		s := New()
		assert.Equal(t, 0, s.Count())
	})

	t.Run("Items passed to New are added to the Set", func(t *testing.T) {
		s := New(1, 2, 3)
		require.Equal(t, 3, s.Count())
		assert.True(t, s.Has(1))
		assert.True(t, s.Has(2))
		assert.True(t, s.Has(3))
	})
}

func TestSetAdd(t *testing.T) {
	t.Run("Add adds items to an existing Set", func(t *testing.T) {
		s := New()
		s.Add(4)
		s.Add(5, 6)
		require.Equal(t, 3, s.Count())
		assert.True(t, s.Has(4))
		assert.True(t, s.Has(5))
		assert.True(t, s.Has(6))
	})

	t.Run("Add initialises an uninitialised Set", func(t *testing.T) {
		var s Set
		s.Add(7)
		s.Add(8)
		s.Add(9)
		require.Equal(t, 3, s.Count())
		assert.True(t, s.Has(7))
		assert.True(t, s.Has(8))
		assert.True(t, s.Has(9))
	})
}

func TestSetCount(t *testing.T) {
	t.Run("Count returns 0 on an empty uninitialised Set", func(t *testing.T) {
		var s Set
		require.Equal(t, 0, s.Count())
	})

	t.Run("Count returns the number of items in the Set", func(t *testing.T) {
		var s Set
		n := rand.Intn(2048)
		for i := 0; i < n; i++ {
			assert.Equal(t, i, s.Count())
			s.Add(i)
			assert.Equal(t, i+1, s.Count())
		}
		require.Equal(t, n, s.Count())
	})
}

func TestSetHas(t *testing.T) {
	t.Run("Has returns false when called on an uninitialised Set", func(t *testing.T) {
		var s Set
		require.False(t, s.Has(12345))
	})

	t.Run("Has returns true for an item previously added to the Set and false otherwise", func(t *testing.T) {
		var s Set

		// Generate some random items (integers) to add to the set
		items := make(map[int]bool)
		n := rand.Intn(2048)
		for i := 0; i < n; i++ {
			items[i] = rand.Int31()&0x01 == 0x01
		}

		// add the items to the set
		for k, v := range items {
			if v {
				s.Add(k)
			}
		}

		for k, v := range items {
			assert.Equal(t, v, s.Has(k))
		}
	})
}

func TestSetRemove(t *testing.T) {
	t.Run("Remove returns cleanly on an uninitialised Set", func(t *testing.T) {
		var s Set
		s.Remove(12345)
	})

	t.Run("Remove removes an item previously added to the Set s.t. Has(item) returns false", func(t *testing.T) {
		var s Set
		s.Add(1, 2, 3, 4, 5)
		require.Equal(t, 5, s.Count())
		s.Remove(3)
		require.Equal(t, 4, s.Count())
		require.True(t, s.Has(1))
		require.True(t, s.Has(2))
		require.False(t, s.Has(3))
		require.True(t, s.Has(4))
		require.True(t, s.Has(5))
	})

	t.Run("Remove is a no-op for an element not in the set", func(t *testing.T) {
		var s Set
		s.Add(1, 2, 3)
		require.Equal(t, 3, s.Count())
		s.Remove(4)
		require.Equal(t, 3, s.Count())
	})
}

func TestSetSlice(t *testing.T) {
	t.Run("Slice on an uninitialised Set returns slice of length 0", func(t *testing.T) {
		var s Set
		require.Equal(t, 0, len(s.Slice()))
	})

	t.Run("Slice returns a list of items in the Set", func(t *testing.T) {
		var s Set
		s.Add(1, 2, 3, 4, 5)
		sl := s.Slice()
		assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, sl)
	})
}