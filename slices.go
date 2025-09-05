package slices

// SafeSlice is a slice that is safe to modify
type SafeSlice[T any] struct {
	slice    []T
	subslice bool
}

func NewSafeSlice[T any](slice []T) *SafeSlice[T] {
	return &SafeSlice[T]{slice: slice}
}

func (s *SafeSlice[T]) GetSlice() []T {
	return s.slice
}

func (s *SafeSlice[T]) SubSlice(start int, end int) *SafeSlice[T] {
	return &SafeSlice[T]{slice: s.slice[start:end], subslice: true}
}

func (s *SafeSlice[T]) Len() int {
	return len(s.slice)
}

func (s *SafeSlice[T]) Get(index int) T {
	return s.slice[index]
}

func (s *SafeSlice[T]) AppendAll(values ...T) *SafeSlice[T] {
	if s.subslice {
		// Make a copy of the slice and return a new SafeSlice
		newSlice := make([]T, len(s.slice))
		copy(newSlice, s.slice)
		newSlice = append(newSlice, values...)
		return &SafeSlice[T]{slice: newSlice}
	}
	s.slice = append(s.slice, values...)
	return s
}

func (s *SafeSlice[T]) Set(index int, value T) {
	if s.subslice {
		// Make a copy of s.slice and then modify it.
		newSlice := make([]T, len(s.slice))
		copy(newSlice, s.slice)
		s.slice = newSlice
		s.subslice = false
	}
	s.slice[index] = value
}
