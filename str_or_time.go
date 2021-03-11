package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {

	s1 := "字符串"
	s2 := "拼接"

	c1 := s1 + s2
	fmt.Println(c1)

	c2 := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(c2)

	t1 := []string{s1, s2}
	c3 := strings.Join(t1, "")
	fmt.Println(c3)

	var b bytes.Buffer
	b.WriteString(s1)
	b.WriteString(s2)
	c4 := b.String()
	fmt.Println(c4)

	//推荐
	var builder strings.Builder
	builder.WriteString(s1)
	builder.WriteString(s2)
	c5 := builder.String()
	fmt.Println(c5)

	y := time.Now().Unix()
	tm := time.Unix(y, 0)
	x := tm.Format("2006-01-02 03:04:05")
	fmt.Println(x)

	unix, _ := time.Parse("2006-01-02", "2021-02-24")
	z := unix.Format("2006/01/02")

	fmt.Println(z)
}
