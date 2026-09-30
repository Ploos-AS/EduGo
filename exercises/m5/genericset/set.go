package genericset

type Set[T comparable] map[T]struct{}

func New[T comparable](values ...T) Set[T] {
	s := make(Set[T], len(values))
	for _, value := range values { s.Add(value) }
	return s
}
func (s Set[T]) Add(value T) { s[value] = struct{}{} }
func (s Set[T]) Contains(value T) bool { _, ok := s[value]; return ok }
func Map[A, B any](in []A, fn func(A) B) []B {
	out := make([]B, len(in))
	for i, value := range in { out[i] = fn(value) }
	return out
}
