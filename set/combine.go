package set

// Union returns a new set with the combined values of sets a and b, where a.Union(b)
func (s1 set[T]) Union(s2 set[T]) set[T] {
	newSet := New[T]()

	newSet.addFromSet(s1)
	newSet.addFromSet(s2)

	return newSet
}

// Intersection returns a new with the common values of of sets a and b, where a.Intersection(b)
func (s1 set[T]) Intersection(s2 set[T]) set[T] {
	newSet := New[T]()

	for key := range s1 {
		if s2.Contains(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

// Difference returns a new set with values unique to set A (ie not in set B), where A.Difference(B)
func (s1 set[T]) Difference(s2 set[T]) set[T] {
	newSet := New[T]()

	for key := range s1 {
		if !s2.Contains(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

// SymmetricDifference returns a new with the common values of of sets A and B, where A.SymmetricDifference(B)
func (s1 set[T]) SymmetricDifference(s2 set[T]) set[T] {
	newSet := New[T]()

	for key := range s1 {
		if !s2.Contains(key) {
			newSet.Add(key)
		}
	}

	for key := range s2 {
		if !s1.Contains(key) {
			newSet.Add(key)
		}
	}

	return newSet
}