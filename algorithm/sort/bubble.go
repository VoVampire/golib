package sort

type Bubble struct{ sorter }

func (s *Bubble) Name() string {
	return "bubble"
}

func (s *Bubble) Sort(a []int) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if s.less(a[j], a[j-1]) {
				s.swap(&a[j], &a[j-1])
			}
		}
	}
}

// 优化：提前终止
func (s *Bubble) Sort2(a []int) {
	l := len(a)
	swapped := true
	for i := 0; swapped; i++ {
		swapped = false
		for j := 1; j < l-i; j++ {
			if s.less(a[j], a[j-1]) {
				s.swap(&a[j], &a[j-1])
				swapped = true
			}
		}
	}
}
