package set

// func (s set[T]) Iterator() {
// 	temp := make([]T, 0, len(s))

// 	for key := range s {
// 		temp = append(temp, key)
// 	}

// 	return iterator{temp, 0}
// }


type iterator[T comparable] struct {
	keys []T
	idx int
}

func (it iterator) Next() {
	return it.idx < len(it.keys)
}

func (it iterator) HasNext() {
	return it.idx < len(it.keys)
}

