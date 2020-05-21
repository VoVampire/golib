package sort

// https://visualgo.net/en/sorting

type Sorter interface {
	Name() string
	Sort([]int)
}

type sorter struct{}

func (s *sorter) less(i, j int) bool {
	return i < j
}

func (s *sorter) swap(i, j *int) {
	*i, *j = *j, *i
}
