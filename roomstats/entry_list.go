package main

import "github.com/unixpickle/cufinals"

type EntryList []cufinals.Entry

func (e EntryList) Len() int {
	return len(e)
}

func (e EntryList) Less(i, j int) bool {
	e1 := e[i]
	e2 := e[j]

	e1HourOff := 0
	e2HourOff := 0
	if !e1.Time.AM {
		e1HourOff = 12
	}
	if !e2.Time.AM {
		e2HourOff = 12
	}

	e1Nums := []int{e1.Date.Month, e1.Date.Day, e1.Time.Hour + e1HourOff, e1.Time.Minute}
	e2Nums := []int{e2.Date.Month, e2.Date.Day, e2.Time.Hour + e2HourOff, e2.Time.Minute}
	for i := 0; i < 4; i++ {
		if e1Nums[i] < e2Nums[i] {
			return true
		} else if e1Nums[i] > e2Nums[i] {
			return false
		}
	}
	return true
}

func (e EntryList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
