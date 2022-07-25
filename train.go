package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	r := strings.Split(s, ";")
	i1 := r[0]
	z1 := strings.Split(i1, " ")
	t1 := strings.Join(z1, "")
	y1 := strings.Split(t1, ",")
	u1 := strings.Join(y1, ".")
	i2 := r[1]
	z2 := strings.Split(i2, " ")
	t2 := strings.Join(z2, "")
	y2 := strings.Split(t2, ",")
	u2 := strings.Join(y2, ".")
	ii1, err := strconv.ParseFloat(u1, 64)
	if err != nil {
		panic(err)
	}
	ii2, err := strconv.ParseFloat(u2, 64)
	if err != nil {
		panic(err)
	}
	i3 := ii1 / ii2
	fmt.Printf("%.4f", i3)
}
