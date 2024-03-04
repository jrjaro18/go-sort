package gosort

import (
	"sync"
	"golang.org/x/exp/constraints"
)


type numType interface {
	constraints.Ordered
}


// Takes a pointer to the slice of type integers, unsigned integers and floats and sorts the slice in ascending order using quick sort algorithm in go routines.
func Sort[T numType](arr *[]T) {
	goQuickSort(arr, 0, len(*arr)-1)
	wg.Wait()
}

func goQuickSort[T numType](arr *[]T, l int, r int) {
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

func getGoPivot[T numType](arr *[]T, l int, r int) int {
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
