package set

// Intersect computes the set intersection of sets s1 and s2, leaving the
// result in set s1.
func (s1 *set) Intersect(s2 Set) {
	for k := range *s1 {
		if !s2.Has(k) {
			delete(*s1, k)
		}
	}
}

// Union computes the set union of sets s1 and s2, leaving the result in set s1.
func (s1 *set) Union(s2 Set) {
	s1.Add(s2.Slice()...)
}

// Subtract subtracts set s2 from set s1, mutating the set s1 in-place.
func (s1 *set) Subtract(s2 Set) {
	if s1 == nil {
		return
	}

	for k := range *s1 {
		if s2.Has(k) {
			delete(*s1, k)
		}
	}
}

// Intersect returns the set intersection of sets s1 and s2 in a new Set,
// without mutating the input Sets.
func Intersect(s1, s2 Set) Set {
	if s1.Count() == 0 || s2.Count() == 0 {
		return Set{}
	}

	var s Set
	for _, el := range s1.Slice() {
		if s2.Has(el) {
			s.Add(el)
		}
	}
	return s
}

// Union returns the set union of sets s1 and s2 in a new Set, without mutating
// the input Sets.
func Union(s1, s2 Set) Set {
	if s1.Count() == 0 && s2.Count() == 0 {
		return Set{}
	}

	var s Set
	s.Add(s1.Slice()...)
	s.Add(s2.Slice()...)
	return s
}

// Subtract returns the subtraction of set s2 from set s1 in a new Set, without
// mutating the input Sets.
func Subtract(s1, s2 Set) Set {
	if s1.Count() == 0 {
		return Set{}
	}
	if s2.Count() == 0 {
		return s1
	}

	var s Set
	s.Add(s1.Slice()...)
	s.Subtract(s2)
	return s
}
