package templates

//go:generate ./generate.sh

var (
	tpl_set = `// vim: syntax=go

package {{.PackageName}}

import (
	"fmt"
	"strings"
)

type {{.InternalSetTypeName}} map[{{.DataType}}]struct{}

// {{.SetTypeName}} is a basic set data structure for elements of type {{.DataType}}. It is
// not safe for concurrent use. Where appropriate, it is the user's
// responsibility to ensure type safety for elements retrieved from the set.
type {{.SetTypeName}} struct {
	// embed set to hide the underlying map
	{{.InternalSetTypeName}}
}

// New instantiates and returns a new {{.SetTypeName}} with the provided initial elements.
func New{{.SetTypeName}}(x ...{{.DataType}}) {{.SetTypeName}} {
	var s {{.SetTypeName}}
	for _, item := range x {
		s.Add(item)
	}
	return s
}

// Add adds the provided element(s) to the {{.SetTypeName}}.
func (s *{{.InternalSetTypeName}}) Add(xs ...{{.DataType}}) {
	if len(xs) == 0 {
		return
	}

	if *s == nil {
		*s = make({{.InternalSetTypeName}})
	}

	for _, x := range xs {
		(*s)[x] = struct{}{}
	}
}

// Count returns the number of elements in the {{.SetTypeName}}.
func (s {{.SetTypeName}}) Count() int {
	if s.{{.InternalSetTypeName}} == nil {
		return 0
	}
	return len(s.{{.InternalSetTypeName}})
}

// Has returns whether the {{.SetTypeName}} contains the specified element.
func (s {{.SetTypeName}}) Has(x {{.DataType}}) bool {
	if len(s.{{.InternalSetTypeName}}) == 0 {
		return false
	}

	_, exists := (s.{{.InternalSetTypeName}})[x]
	return exists
}

// Remove removes the specified element from the {{.SetTypeName}}. This is a no-op if the
// {{.SetTypeName}} does not contain the element.
func (s *{{.InternalSetTypeName}}) Remove(x {{.DataType}}) {
	if *s == nil {
		return
	}

	delete(*s, x)
}

// Slice returns a slice of the elements contained in the {{.SetTypeName}}.
func (s {{.SetTypeName}}) Slice() []{{.DataType}} {
	sl := make([]{{.DataType}}, 0, len(s.{{.InternalSetTypeName}}))
	for k := range s.{{.InternalSetTypeName}} {
		sl = append(sl, k)
	}
	return sl
}

// String implements fmt.Stringer
func (s {{.SetTypeName}}) String() string {
	str := make([]string, 0, len(s.{{.InternalSetTypeName}}))
	for el := range s.{{.InternalSetTypeName}} {
		str = append(str, fmt.Sprintf("%v", el))
	}
	return fmt.Sprintf("{{.SetTypeName}}{%s}", strings.Join(str, ", "))
}`

	tpl_set_test = `// vim: syntax=go

package {{.PackageName}}

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
 	rand.Seed(time.Now().Unix())
}

// Assert existence of item getter function for use in tests. If your tests
// fail here, you need to define the function New{{.SetTypeName}}Item and have
// it return a new instance of the item type {{.DataType}} in your set, so the
// tests have content to manipulate in the {{.SetTypeName}}.
var _ = New{{.SetTypeName}}Item

func get{{.SetTypeName}}Items(n int) []{{.DataType}} {
	items := make([]{{.DataType}}, n)
	for i := 0; i < n; i++ {
		items[i] = New{{.SetTypeName}}Item()
	}

	return items
}

func Assert{{.SetTypeName}}ElementsMatch(t assert.TestingT, expected []{{.DataType}}, set {{.SetTypeName}}) (ok bool) {
	if h, ok := t.(interface {
		Helper()
	}); ok {
		h.Helper()
	}
	return assert.ElementsMatch(t, expected, set.Slice())
}

func Require{{.SetTypeName}}ElementsMatch(t require.TestingT, expected []{{.DataType}}, set {{.SetTypeName}}) {
	if h, ok := t.(interface {
		Helper()
	}); ok {
		h.Helper()
	}
	require.ElementsMatch(t, expected, set.Slice())
}

func Test{{.SetTypeName}}New(t *testing.T) {
	t.Run("New creates an empty set", func(t *testing.T) {
		s := New{{.SetTypeName}}()
		assert.Equal(t, 0, s.Count())
	})

	t.Run("Items passed to New are added to the {{.SetTypeName}}", func(t *testing.T) {
		items := get{{.SetTypeName}}Items(3)
		s := New{{.SetTypeName}}(items...)
		require.Equal(t, 3, s.Count())
		assert.True(t, s.Has(items[0]))
		assert.True(t, s.Has(items[1]))
		assert.True(t, s.Has(items[2]))
	})
}

func Test{{.SetTypeName}}Add(t *testing.T) {
	t.Run("Add adds items to an existing {{.SetTypeName}}", func(t *testing.T) {
		s := New{{.SetTypeName}}()
		items := get{{.SetTypeName}}Items(3)

		s.Add(items[0])
		s.Add(items[1:]...)
		require.Equal(t, 3, s.Count())
		assert.True(t, s.Has(items[0]))
		assert.True(t, s.Has(items[1]))
		assert.True(t, s.Has(items[2]))
		Assert{{.SetTypeName}}ElementsMatch(t, items, s)
	})

	t.Run("Add initialises an uninitialised {{.SetTypeName}}", func(t *testing.T) {
		var s {{.SetTypeName}}
		items := get{{.SetTypeName}}Items(3)
		s.Add(items[0])
		s.Add(items[1])
		s.Add(items[2])
		require.Equal(t, 3, s.Count())
		assert.True(t, s.Has(items[0]))
		assert.True(t, s.Has(items[1]))
		assert.True(t, s.Has(items[2]))
		Assert{{.SetTypeName}}ElementsMatch(t, items, s)
	})
}

func Test{{.SetTypeName}}Count(t *testing.T) {
	t.Run("Count returns 0 on an empty uninitialised {{.SetTypeName}}", func(t *testing.T) {
		var s {{.SetTypeName}}
		require.Equal(t, 0, s.Count())
	})

	t.Run("Count returns the number of items in the {{.SetTypeName}}", func(t *testing.T) {
		var s {{.SetTypeName}}
		n := rand.Intn(2048) + 100
		for i := 0; i < n; i++ {
			assert.Equal(t, i, s.Count())
			s.Add(get{{.SetTypeName}}Items(1)[0])
			assert.Equal(t, i+1, s.Count())
		}
		require.Equal(t, n, s.Count())
	})
}

func Test{{.SetTypeName}}Has(t *testing.T) {
	t.Run("Has returns false when called on an uninitialised {{.SetTypeName}}", func(t *testing.T) {
		var s {{.SetTypeName}}
		require.False(t, s.Has(get{{.SetTypeName}}Items(1)[0]))
	})

	t.Run("Has returns true for an item previously added to the {{.SetTypeName}} and false otherwise", func(t *testing.T) {
		var s {{.SetTypeName}}

		type setTestItem struct {
			dt {{.DataType}}
			include bool
		}

		// Generate some random items (integers) to add to the set and
		// others not to add but use in checking
		n := rand.Intn(2048) + 100
		items := make([]setTestItem, n)
		for i := 0; i < n; i++ {
			items[i] = setTestItem{
				dt: get{{.SetTypeName}}Items(1)[0],
				include: rand.Int31()&0x01 == 0x01,
			}
		}

		// add the items to the set
		for _, item := range items {
			if item.include {
				s.Add(item.dt)
			}
		}

		for _, item := range items {
			assert.Equal(t, item.include, s.Has(item.dt))
		}
	})
}

func Test{{.SetTypeName}}Remove(t *testing.T) {
	t.Run("Remove returns cleanly on an uninitialised {{.SetTypeName}}", func(t *testing.T) {
		var s {{.SetTypeName}}
		s.Remove(get{{.SetTypeName}}Items(1)[0])
		require.Equal(t, 0, s.Count())
		Assert{{.SetTypeName}}ElementsMatch(t, nil, s)
	})

	t.Run("Remove removes an item previously added to the {{.SetTypeName}} s.t. Has(item) returns false", func(t *testing.T) {
		var s {{.SetTypeName}}
		items := get{{.SetTypeName}}Items(5)
		s.Add(items...)
		require.Equal(t, 5, s.Count())
		s.Remove(items[2])
		require.Equal(t, 4, s.Count())
		require.False(t, s.Has(items[2]))

		expectItems := append(items[:2], items[3:]...)
		for _, expectItem := range expectItems {
			require.True(t, s.Has(expectItem))
		}

		Require{{.SetTypeName}}ElementsMatch(t, expectItems, s)
	})

	t.Run("Remove is a no-op for an element not in the set", func(t *testing.T) {
		var s {{.SetTypeName}}
		items := get{{.SetTypeName}}Items(3)
		s.Add(items...)
		require.Equal(t, 3, s.Count())
		s.Remove(get{{.SetTypeName}}Items(1)[0])
		require.Equal(t, 3, s.Count())
		Require{{.SetTypeName}}ElementsMatch(t, items, s)
	})
}

func Test{{.SetTypeName}}Slice(t *testing.T) {
	t.Run("Slice on an uninitialised {{.SetTypeName}} returns slice of length 0", func(t *testing.T) {
		var s {{.SetTypeName}}
		require.Equal(t, 0, len(s.Slice()))
	})

	t.Run("Slice returns a list of items in the {{.SetTypeName}}", func(t *testing.T) {
		var s {{.SetTypeName}}
		items := get{{.SetTypeName}}Items(5)
		s.Add(items...)
		sl := s.Slice()
		assert.ElementsMatch(t, items, sl)
	})
}

func Test{{.SetTypeName}}String(t *testing.T) {
	tcs := [][]{{.DataType}}{
		{},
		get{{.SetTypeName}}Items(5),
		get{{.SetTypeName}}Items(1),
	}

	const offset = len(` + "`" + `{{.SetTypeName}}` + "`" + `)
	for _, tc := range tcs {
		s := strings.TrimSpace(New{{.SetTypeName}}(tc...).String())
		assert.Equal(t, "{{.SetTypeName}}{", s[:offset+1])
		assert.Equal(t, '}', rune(s[len(s)-1]))

		contents := s[offset+1 : len(s)-1]
		var parts []string
		if contents != "" {
			parts = strings.Split(contents, ", ")
		}
		strElts := make([]string, 0, len(tc))
		for _, elt := range tc {
			strElts = append(strElts, fmt.Sprintf("%v", elt))
		}
		assert.ElementsMatch(t, strElts, parts)
	}
}`
)
