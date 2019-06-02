package set

func (s1 *set) Union(s2 Set) {
	s1.Add(s2.Slice()...)
}

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
