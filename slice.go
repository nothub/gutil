package gutil

func Dedup[T any](a []T) []T {
	m := make(map[any]bool, len(a))
	u := make([]T, len(m))
	for _, s := range a {
		if !m[s] {
			u = append(u, s)
			m[s] = true
		}
	}
	return u
}
