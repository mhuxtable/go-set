package set

type set map[interface{}]struct{}

// Set is a basic set data structure. It is not safe for concurrent use. The
// only requirement is that elements added to Set are comparable in the sense
// defined in the Go language spec. It is the user's responsibility to manage
// type safety for any elements retrieved from the set.
type Set struct {
	// embed set to hide the underlying map
	set
}

// New instantiates and returns a new Set with the provided initial elements.
func New(x ...interface{}) Set {
	var s Set
	for _, item := range x {
		s.Add(item)
	}
	return s
}

// Add adds the provided element(s) to the Set.
func (s *set) Add(xs ...interface{}) {
	if len(xs) == 0 {
		return
	}

	if *s == nil {
		*s = make(set)
	}

	for _, x := range xs {
		(*s)[x] = struct{}{}
	}
}

// Count returns the number of elements in the Set.
func (s Set) Count() int {
	if s.set == nil {
		return 0
	}
	return len(s.set)
}

// Has returns whether the Set contains the specified element.
func (s Set) Has(x interface{}) bool {
	if len(s.set) == 0 {
		return false
	}

	_, exists := (s.set)[x]
	return exists
}

// Remove removes the specified element from the Set. This is a no-op if the
// Set does not contain the element.
func (s *set) Remove(x interface{}) {
	if *s == nil {
		return
	}

	delete(*s, x)
}

// Slice returns a slice of the elements contained in the Set.
func (s Set) Slice() []interface{} {
	sl := make([]interface{}, 0, len(s.set))
	for k, _ := range s.set {
		sl = append(sl, k)
	}
	return sl
}
