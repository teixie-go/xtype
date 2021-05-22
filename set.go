package xtype

var (
	_ IntCollection    = (*intSet)(nil)
	_ Int64Collection  = (*int64Set)(nil)
	_ StringCollection = (*stringSet)(nil)
)

type intSet struct {
	intCollection

	exists map[int]bool
}

func (s *intSet) Add(members ...int) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *intSet) Contains(val int) bool {
	return s.exists[val]
}

func NewIntSet(members ...int) *intSet {
	set := &intSet{
		intCollection: intCollection{
			members: make([]int, 0),
		},
		exists: make(map[int]bool),
	}
	set.Add(members...)
	return set
}

//------------------------------------------------------------------------------

type int64Set struct {
	int64Collection

	exists map[int64]bool
}

func (s *int64Set) Add(members ...int64) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *int64Set) Contains(val int64) bool {
	return s.exists[val]
}

func NewInt64Set(members ...int64) *int64Set {
	set := &int64Set{
		int64Collection: int64Collection{
			members: make([]int64, 0),
		},
		exists: make(map[int64]bool),
	}
	set.Add(members...)
	return set
}

//------------------------------------------------------------------------------

type stringSet struct {
	stringCollection

	exists map[string]bool
}

func (s *stringSet) Add(members ...string) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *stringSet) Contains(val string) bool {
	return s.exists[val]
}

func NewStringSet(members ...string) *stringSet {
	set := &stringSet{
		stringCollection: stringCollection{
			members: make([]string, 0),
		},
		exists: make(map[string]bool),
	}
	set.Add(members...)
	return set
}
