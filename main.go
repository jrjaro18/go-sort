package gosort

import (
	"sync"
	"golang.org/x/exp/constraints"
)

func main() {
	arr := []int{5,300,44,3,55,2}
	Sort(&arr)
}

type Num interface {
	constraints.Ordered
}

func Sort[T Num](arr *[]T) {
	goQuickSort(arr, 0, len(*arr)-1)
	wg.Wait()
}

func goQuickSort[T Num](arr *[]T, l int, r int) {
	if l < r {
		pivot := getGoPivot(arr, l, r)
		if pivot-l+1 > 15 && r-pivot+1 > 15 {
			wg.Add(2) //number of routines added in the group
			go func() {
				defer wg.Done() //always remove the goroutine from the go function only
				goQuickSort(arr, l, pivot-1)
			}()
			go func() {
				defer wg.Done()
				goQuickSort(arr, pivot+1, r)
			}()
		} else {
			goQuickSort(arr, l, pivot-1)
			goQuickSort(arr, pivot+1, r)
		}
	}
}

func getGoPivot[T Num](arr *[]T, l int, r int) int {
	pivot := (*arr)[r]
	i := l - 1
	for j := l; j <= r-1; j++ {
		if (*arr)[j] < pivot {
			i++
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = temp
		}
	}
	temp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[r]
	(*arr)[r] = temp
	return i + 1
}

var wg = sync.WaitGroup{}
