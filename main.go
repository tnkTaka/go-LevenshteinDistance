package main

import (
	"github.com/go-LevenshteinDistance/distance"
)

func main() {

	t1 := "golang"
	t2 := "go言語"

	print(distance.LevenshteinDistance(t1,t2))
}
