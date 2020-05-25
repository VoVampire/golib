package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for _, v := range []Sorter{
		&GoSort{},
		&Shell{},
		&Quick{},
		&Quick2{},
		&QuickInsertion{},
		&Insertion{},
		&Selection{},
		&Bubble{},
		&Bubble2{},
	} {
		arr := doSort(v.Name, v.Sort, randArr())
		assert.True(t, isSorted(arr))
	}
}

func doSort(name func() string, fn func([]int), a []int) []int {
	fmt.Printf("%-10s: ", name())
	now := time.Now()
	fn(a)
	fmt.Println(time.Now().Sub(now).Microseconds(), "μs")
	return a
}

func randArr() []int {
	var a []int
	for i := 0; i < 20000; i++ {
		a = append(a, rand.Intn(1000000))
	}
	return a
}

func ascArr() []int {
	var a []int
	for i := 0; i < 10000; i++ {
		a = append(a, i)
	}
	return a
}

func descArr() []int {
	var a []int
	for i := 10000; i > 0; i-- {
		a = append(a, i)
	}
	return a
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
