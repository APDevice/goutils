package set

// IsDisjoint returns true if there is no intersection between sets
func (s1 set[T]) IsDisjoint(s2 set[T]) bool {
	for key := range s1 {
		if _, ok := range s2[key]; ok {
			return false
		}
	}

	return true
}

// IsSubset returns true if all of set A is contained in set B, where A.IsSubset(B)
func (s1 set[T]) IsSubset(s2 set[T]) bool {
	for key := range s1 {
		if _, ok := range s2[key]; !ok {
			return false
		}
	}

	return true
}

// IsSuperset returns true if the values of set A contains all of set B, where A.IsSubset(B)
func (s1 set[T]) IsSubset(s2 set[T]) bool {
	for key := range s2 {
		if _, ok := range s1[key]; !ok {
			return false
		}
	}

	return true
}