package set

import "github.com/mhuxtable/go-set/genericset"

// compat contains methods for backwards compatibility with old package interfaces.

// Deprecated. Set is a generic set for items of type interface{}. This exists
// for backwards compatibility; new code should use genericset.Set directly.
type Set = genericset.Set

// Deprecated. New returns a new genericset containing the supplied items. This
// exists for backwards compatibility; new code should call genericset.NewSet
// directly.
func New(xs ...interface{}) Set {
	return genericset.NewSet(xs...)
}

// Deprecated. Intersect computes the intersection of two sets, returning the
// result in a third set. This exists for backwards compatibility; new code
// should call genericset.IntersectSet directly.
func Intersect(s1, s2 Set) Set {
	return genericset.IntersectSet(s1, s2)
}

// Deprecated. Union computes the set union of two sets, returning the result
// in a third set. This exists for backwards compatibility; new code should
// call genericset.UnionSet directly.
func Union(s1, s2 Set) Set {
	return genericset.UnionSet(s1, s2)
}

// Deprecated. Subtract computes the subtraction of set s2 from set s1,
// returning the result in a third set. This exists for backwards
// compatibility; new code should call genericset.SubtractSet directly.
func Subtract(s1, s2 Set) Set {
	return genericset.SubtractSet(s1, s2)
}
