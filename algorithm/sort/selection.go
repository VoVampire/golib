package sort

type Selection struct{ sorter }

func (s *Selection) Name() string {
	return "selection"
}

func (s *Selection) Sort(a []int) {
	l := len(a)
	for i := 0; i < l; i++ {
		m := i
		for j := i + 1; j < l; j++ {
			if s.less(a[j], a[m]) {
				m = j
			}
		}
		s.swap(&a[i], &a[m])
	}
}
