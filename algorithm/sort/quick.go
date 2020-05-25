package sort

type Quick struct{}

func (q *Quick) Name() string {
	return "quick"
}

func (q *Quick) Sort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	q.Sort(data[:head])
	q.Sort(data[head+1:])
}

type Quick2 struct{}

func (q *Quick2) Name() string {
	return "quick-2"
}

func (q *Quick2) Sort(data []int) {
	q.sort(data, 0, len(data)-1)
}

func (q *Quick2) sort(data []int, head, tail int) {
	if head < tail {
		mid := q.partition(data, head, tail)
		q.sort(data, head, mid-1)
		q.sort(data, mid+1, tail)
	}
}

func (q *Quick2) partition(data []int, head, tail int) int {
	mid := data[head]
	for i := head + 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	return head
}
