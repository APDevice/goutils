package set

import (
	"fmt"
	"bytes"
)

// void is a dummy type used in place of map value
type void struct{}

// set is a datatype that stores unique values
type set[T comparable] map[T]void

/* New returns a new set of a given type T

Usage: variable := set.New[type]()
*/

func New[T comparable]() set[T] {
	return make(set[T])
}

// Contains returns true if value is in set
func (s set[T]) Contains(val T) bool {
	_, contains := s[val]

	return contains == true
}

// Add adds the specified element to set
func (s *set[T]) Add(val T) {
	(*s)[val] = void{}
}

// Clear removes all elements from set without deleting the set itself
func (s *set[T]) Clear() {
	for key := range *s {
		delete(*s, key)
	}
}

// IsEmpty returns true if set is empty
func (s set[T]) IsEmpty() bool {
	return len(s) == 0
}


/* AddAll adds all the elements from another set of type T or other compatible container of T

Input:
 - set[T] or []T

 Output:
 - (optional) error: if the container passed in is not supported, an error will be passed indicating the incompatibility
*/
func (s *set[T]) AddAll(vals interface{}) error {
	switch vals.(type) {
		case set[T]:
			s.addFromSet( vals.(set[T]) )
		case []T:
			s.addFromSlice( vals.([]T) )
		default:
			return fmt.Errorf("type %T unsupported", vals)
	}

	return nil
}

// addFromSlice impliments the AddAll method for other sets
func (s *set[T]) addFromSet(vals set[T]) {
	for v := range vals {
		(*s).Add(v)
	}
}

// addFromSlice impliments the AddAll method for slices
func (s *set[T]) addFromSlice(vals []T) {
	for _, v := range vals {
		(*s).Add(v) 
	}
}



// Remove deletes the specified element from set
func (s *set[T]) Remove(val T) {
	delete(*s, val)
}

/* String implements the Stringer interface, allowing set to be printed out with fmt.Print, etc

Sample output: set[1 2 3]
 */
func (s set[T]) String() string {
	var output bytes.Buffer
	
	output.WriteString("set[")
	
	for key := range s {
		output.WriteString(fmt.Sprintf("%v ", key))
	}

	if len(s) > 1 {
		output.Truncate(output.Len()-1)
	}
	
	output.WriteString("]")
	
	return output.String()
	
}

