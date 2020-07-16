package main

type fToBool func(int) bool
type twoArrays func([]int)([]int, []int)

func filterFactory(f fToBool) twoArrays {
	return func(arr []int)(t, fl []int) {
		for _, val := range arr {
			if f(val) {
				t = append(t, val)
			} else {
				fl = append(fl, val)
			}
		}
		return
	}
}
