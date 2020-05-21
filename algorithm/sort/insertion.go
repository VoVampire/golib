package sort

type Insertion struct{ sorter }

func (s *Insertion) Name() string {
	return "insertion"
}

func (s *Insertion) Sort(a []int) {
	l := len(a)
	for i := 1; i < l; i++ {
		for j := i; j > 0 && s.less(a[j], a[j-1]); j-- {
			s.swap(&a[j], &a[j-1])
		}
	}
}
