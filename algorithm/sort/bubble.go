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

type Bubble2 struct{ sorter }

func (s *Bubble2) Name() string {
	return "bubble-2"
}

// 提前终止
func (s *Bubble2) Sort(a []int) {
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
