package sort

type Selection struct{ sorter }

func (s *Selection) Name() string {
	return "selection"
}

func (s *Selection) Sort(a []int) {
	l := len(a)
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if s.less(a[j], a[min]) {
				min = j
			}
		}
		s.swap(&a[i], &a[min])
	}
}
