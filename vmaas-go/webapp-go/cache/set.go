/* Custom set implementation */

package cache

var exists = struct{}{}

type intSet struct {
    m map[int]struct{}
}

func NewIntSet() *intSet {
    s := intSet{}
    s.m = map[int]struct{}{}
	return &s
}

func (s *intSet) Add(value int) {
    s.m[value] = exists
}

func (s *intSet) Remove(value int) {
    delete(s.m, value)
}

func (s *intSet) Contains(value int) bool {
    _, c := s.m[value]
    return c
}

func (s *intSet) Len() int {
    return len(s.m)
}
