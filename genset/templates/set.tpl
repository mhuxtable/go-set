// vim: syntax=go

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
}
