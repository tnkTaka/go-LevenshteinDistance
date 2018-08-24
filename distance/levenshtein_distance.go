package distance

import (
	"strings"
	"math"
)

func LevenshteinDistance(t1,t2 string) int {
	if t1 == t2{
		return 0
	}

	t1Array := strings.Split(t1,"")
	t2Array := strings.Split(t2,"")

	lenT1Array := len(t1Array)
	lenT2Array := len(t2Array)

	if lenT1Array == 0 {
		return lenT2Array
	}

	if lenT2Array == 0 {
		return lenT1Array
	}

	m := make([][]int, lenT1Array+1)
	for i := range m {
		m[i] = make([]int, lenT2Array+1)
	}

	for i := 0; i < lenT1Array+1; i++ {
		for j := 0; j < lenT2Array+1; j++ {
			if i == 0 {
				m[i][j] = j
			} else if j == 0 {
				m[i][j] = i
			} else {
				if t1Array[i-1] == t2Array[j-1] {
					m[i][j] = m[i-1][j-1]
				} else {
					m[i][j] = Min(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+1)
				}
			}
		}
	}

	return m[lenT1Array][lenT2Array]
}

func Min(is ...int) int {
	min := int(math.MaxInt64)

	for _, v := range is {
		if min > v {
			min = v
		}
	}

	return min
}