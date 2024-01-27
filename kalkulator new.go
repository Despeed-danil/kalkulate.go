package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := bufio.NewReader(os.Stdin)
	z, _ := a.ReadString('\n')
	if strings.Contains(z, "+") {
		x := strings.Split(z, "+")
		for qwe := 0; qwe < 2; qwe++ {
			p := x[qwe]
			if len(p) > 12 {
				panic("ваша строка выходит за рамки 10 символов")
			}
		}
		if len(x) == 2 {
			c := x[0]
			c = strings.TrimSpace(c)
			r := strings.Split(c, "")
			k := len(r)
			qq := r[k-1]
			qq = strings.TrimSpace(qq)
			if r[0] == "\"" && qq == "\"" {
				c = ""
				for te := 1; te < len(r)-1; te++ {
					c = c + r[te]
				}
			} else {
				panic("первый аргумент не является строкой")
			}
			d := x[1]
			d = strings.TrimSpace(d)
			r = strings.Split(d, "")
			k = len(r)
			if r[0] == "\"" && qq == "\"" {
				d = ""
				for te := 1; te < len(r)-1; te++ {
					d = d + r[te]
				}
				fmt.Println(c + d)
			} else {
				panic("числа не могут добавляться к строке")
			}
		}
	}
	if strings.Contains(z, "-") {
		x := strings.Split(z, "-")
		for qwe := 0; qwe < 2; qwe++ {
			p := x[qwe]
			if len(p) > 12 {
				panic("ваша строка выходит за рамки 10 символов")
			}
		}
		if len(x) == 2 {
			c := x[0]
			c = strings.TrimSpace(c)
			r := strings.Split(c, "")
			k := len(r)
			qq := r[k-1]
			qq = strings.TrimSpace(qq)
			if r[0] == "\"" && qq == "\"" {
				c = ""
				for te := 1; te < len(r)-1; te++ {
					c = c + r[te]
				}
			} else {
				panic("первый аргумент не является строкой")
			}
			d := x[1]
			d = strings.TrimSpace(d)
			r = strings.Split(d, "")
			k = len(r)
			if r[0] == "\"" && qq == "\"" {
				d = ""
				for te := 1; te < len(r)-1; te++ {
					d = d + r[te]
				}
				g := strings.Split(c, d)

				fmt.Print("\"", g[0]+g[1], "\"")
			} else {
				panic("числа не могут вычитаться из строки")
			}
		}
	}
	if strings.Contains(z, "*") {
		x := strings.Split(z, "*")
		umn := ""
		p := x[0]
		if len(p) > 12 {
			panic("ваша строкаа выходит за рамки 10 символов")
		}
		if len(x) == 2 {
			c := x[0]
			c = strings.TrimSpace(c)
			r := strings.Split(c, "")

			k := len(r)
			qq := r[k-1]
			qq = strings.TrimSpace(qq)
			if r[0] == "\"" && qq == "\"" {
				for te := 1; te < k-1; te++ {
					umn = umn + r[te]
				}
				d := x[1]
				d = strings.TrimSpace(d)
				if strings.Contains(d, "\"") {
					panic("второй аргумент не является числом")
				} else {
					k, _ := strconv.Atoi(d)
					i := ""
					for q := 0; q < k; q++ {
						i = i + umn
					}
					if k < 11 && len(i) <= 40 {
						fmt.Printf("\"%s\"", i)
					} else if k < 11 && len(i) > 40 {
						zz := ""
						for qw := 0; qw < 40; qw++ {
							qww := strings.Split(i, "")
							d := qww[qw]
							zz = d + zz
						}
						fmt.Printf("\"%s...\"", zz)
					} else {
						panic("число больше чем 10")
					}

				}
			} else {
				panic("первый аргумент не является строкой")
			}

		}
	}
	if strings.Contains(z, "/") {
		x := strings.Split(z, "/")
		p := x[0]
		if len(p) > 12 {
			panic("ваша строкаа выходит за рамки 10 символов")
		}
		umn := ""
		if len(x) == 2 {
			c := x[0]
			c = strings.TrimSpace(c)
			r := strings.Split(c, "")
			k := len(r)
			o := len(r)
			qq := r[k-1]
			qq = strings.TrimSpace(qq)
			if r[0] == "\"" && qq == "\"" {
				for te := 1; te < k-1; te++ {
					umn = umn + r[te]
				}
				d := x[1]
				d = strings.TrimSpace(d)
				if strings.Contains(d, "\"") {
					panic("второй аргумент не является числом")
				} else {
					k, _ := strconv.Atoi(d)
					i := ""
					j := strings.Split(umn, "")
					for q := 0; q < o/k; q++ {
						i = i + j[q]
					}
					if k < 11 {
						fmt.Printf("\"%s\"", i)
					} else {
						panic("число больше чем 10")
					}
				}
			} else {
				panic("первый аргумент не является строкой")
			}

		}
	}
}
