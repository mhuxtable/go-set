// vim: syntax=go

package {{.PackageName}}

import (
	"fmt"
	"reflect"
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

// ZZ_GoSetContents returns a copy of the elements stored in this set in a
// generic form. This is used by generic operations on the set in the go-set
// package.
//
// This is not a stable interface.
func (s {{.SetTypeName}}) ZZ_GoSetContents() []interface{} {
	xs := make([]interface{}, 0, len(s.{{.InternalSetTypeName}}))
	for k := range s.{{.InternalSetTypeName}} {
		xs = append(xs, k)
	}

	return xs
}

// ZZ_GoSetType returns the type of elements stored in this set. This is used
// by generic operations on the set in the go-set package.
//
// This is not a stable interface.
func (*{{.SetTypeName}}) ZZ_GoSetType() reflect.Type {
	var x {{.DataType}}
	return reflect.TypeOf(x)
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

// Intersect computes the set intersection of sets s1 and s2, leaving the
// result in set s1.
func (s1 *{{.InternalSetTypeName}}) Intersect(s2 {{.SetTypeName}}) {
	for k := range *s1 {
		if !s2.Has(k) {
			delete(*s1, k)
		}
	}
}

// Union computes the set union of sets s1 and s2, leaving the result in set s1.
func (s1 *{{.InternalSetTypeName}}) Union(s2 {{.SetTypeName}}) {
	s1.Add(s2.Slice()...)
}

// Subtract subtracts set s2 from set s1, mutating the set s1 in-place.
func (s1 *{{.InternalSetTypeName}}) Subtract(s2 {{.SetTypeName}}) {
	for k := range *s1 {
		if s2.Has(k) {
			delete(*s1, k)
		}
	}
}

// Intersect returns the set intersection of sets s1 and s2 in a new {{.SetTypeName}},
// without mutating the input {{.SetTypeName}}s.
func Intersect{{.SetTypeName}}(s1, s2 {{.SetTypeName}}) {{.SetTypeName}} {
	if s1.Count() == 0 || s2.Count() == 0 {
		return {{.SetTypeName}}{}
	}

	var s {{.SetTypeName}}
	for _, el := range s1.Slice() {
		if s2.Has(el) {
			s.Add(el)
		}
	}
	return s
}

// Union returns the set union of sets s1 and s2 in a new {{.SetTypeName}}, without mutating
// the input {{.SetTypeName}}s.
func Union{{.SetTypeName}}(s1, s2 {{.SetTypeName}}) {{.SetTypeName}} {
	if s1.Count() == 0 && s2.Count() == 0 {
		return {{.SetTypeName}}{}
	}

	var s {{.SetTypeName}}
	s.Add(s1.Slice()...)
	s.Add(s2.Slice()...)
	return s
}

// Subtract returns the subtraction of set s2 from set s1 in a new {{.SetTypeName}}, without
// mutating the input {{.SetTypeName}}s.
func Subtract{{.SetTypeName}}(s1, s2 {{.SetTypeName}}) {{.SetTypeName}} {
	if s1.Count() == 0 {
		return {{.SetTypeName}}{}
	}
	if s2.Count() == 0 {
		return s1
	}

	var s {{.SetTypeName}}
	s.Add(s1.Slice()...)
	s.Subtract(s2)
	return s
}
