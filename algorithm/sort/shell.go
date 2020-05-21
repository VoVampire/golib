package sort

type Shell struct{ sorter }

func (s *Shell) Name() string {
	return "shell"
}

func (s *Shell) Sort(a []int) {
	l := len(a)
	h := 1
	for h < l/3 {
		h = 3*h + 1
	}

	for ; h >= 1; h /= 3 {
		for i := h; i < l; i++ {
			for j := i; j >= h && s.less(a[j], a[j-h]); j -= h {
				s.swap(&a[j], &a[j-h])
			}
		}
	}
}
