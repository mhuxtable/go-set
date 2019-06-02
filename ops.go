package set

// Union returns the set union of sets s1 and s2 by mutating s1
func (s1 *set) Union(s2 Set) {
	s1.Add(s2.Slice()...)
}

// Subtract subtracts set s2 from set s1, mutating the set s1 in-place.
func (s1 *set) Subtract(s2 Set) {
	if s1 == nil {
		return
	}

	for k, _ := range *s1 {
		if s2.Has(k) {
			delete(*s1, k)
		}
	}
}
