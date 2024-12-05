package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input.txt")
	str := string(b)

	groups := strings.Split(str, "\n\n")
	or := strings.Split(groups[0], "\n")
	pu := strings.Split(groups[1], "\n")
	p1(or, pu)
}

func p1(or, pu []string) {
	total := 0
	incOrdered := [][]string{}
	nt := 0

OuterLoop:
	for _, u := range pu {
		updI := strings.Split(u, ",")
		for current := 0; current < len(updI)-1; current++ {
			for next := current + 1; next < len(updI); next++ {
				tr := updI[current] + "|" + updI[next]
				if !doesExist(tr, or) {
					incOrdered = append(incOrdered, updI)
					continue OuterLoop
				}
			}
		}
		m := int(math.Floor(float64(len(updI) / 2)))
		pN, _ := strconv.Atoi(updI[m])
		total += pN
	}

	fmt.Println(total)

	for _, incO := range incOrdered {
		for current := 0; current < len(incO)-1; current++ {
			for next := current + 1; next < len(incO); next++ {
				tr := incO[current] + "|" + incO[next]
				if !doesExist(tr, or) {
					c := incO[current]
					incO[current] = incO[next]
					incO[next] = c
				}
			}
		}
		m := int(math.Floor(float64(len(incO) / 2)))
		pN, _ := strconv.Atoi(incO[m])
		nt += pN
	}

	fmt.Println(nt)
}

func doesExist(tr string, or []string) bool {

	for _, r := range or {
		if tr == r {
			return true
		}
	}

	return false
}
