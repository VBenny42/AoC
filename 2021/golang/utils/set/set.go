package set

type Set[T comparable] map[T]struct{}

// NewSet creates a new set with the given values.
func NewSet[T comparable](values ...T) Set[T] {
	s := make(Set[T])
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

// Add adds the given value to the set.
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Remove removes the given value from the set.
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Contains returns true if the set contains the given value.
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// Values returns a slice of all the values in the set.
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}

// Clone returns a copy of the set.
func (s Set[T]) Clone() Set[T] {
	clone := make(Set[T], len(s))
	for v := range s {
		clone[v] = struct{}{}
	}
	return clone
}
