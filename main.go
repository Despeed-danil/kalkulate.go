// Program to access individual character of string

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* перед тем как проверять код, поверьте я не мазахист что создавал такие функции, как не пытался никак не мог через strconv.aci сменить строковый тип string на целочисленный int,
функции через перечисления фор подстанавливают числа ввместо текста*/

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}
func minus(q string, w string) int {
	l := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ax := 0
	az := 0
	for o := 0; o < 9; o++ {
		if l[o] == w {
			ax = o + 1
		}
		if l[o] == q {
			az = o + 1
		}

	}
	return az - ax
}
func del(q string, w string) int {
	l := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ax := 0
	az := 0
	for o := 0; o < 9; o++ {
		if l[o] == w {
			ax = o + 1
		}
		if l[o] == q {
			az = o + 1
		}

	}
	return ax / az
}
func umn(q string, w string) int {
	l := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ax := 0
	az := 0
	for o := 0; o < 9; o++ {
		if l[o] == w {
			ax = o + 1
		}
		if l[o] == q {
			az = o + 1
		}

	}
	return ax * az
}
func summ(q string, w string) int {
	l := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ax := 0
	az := 0
	for o := 0; o < 9; o++ {
		if l[o] == w {
			ax = o + 1
		}
		if l[o] == q {
			az = o + 1
		}

	}
	return ax + az
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("введите значение ")
	s, _ := reader.ReadString('\n')
	s = strings.ReplaceAll(s, " ", "")
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		panic("Выдача паники, так как строка не является математической операцией.")
		os.Exit(0)
	}
	if (strings.Contains(s, "1") || strings.Contains(s, "2") || strings.Contains(s, "3") || strings.Contains(s, "4") || strings.Contains(s, "5") || strings.Contains(s, "6") || strings.Contains(s, "7") || strings.Contains(s, "8") || strings.Contains(s, "9")) && (strings.Contains(s, "I") || strings.Contains(s, "V") || strings.Contains(s, "X")) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
		os.Exit(0)
	}
	if strings.Contains(s, "+") {
		s := strings.Split(s, "+")
		a := s[0]
		b := s[1]
		if len(s) > 2 {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			os.Exit(0)
		}
		if len(s) == 2 && (a == "I" || a == "II" || a == "III" || a == "IV" || a == "V" || a == "VI" || a == "VII" || a == "VIII" || a == "IX" || a == "X") && (b == "I" || b == "II" || b == "III" || b == "IV" || b == "V" || b == "VI" || b == "VII" || b == "VIII" || b == "IX" || b == "X") {
			ff := romanToInt(a) + romanToInt(b)
			qq := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI"}
			for ww := 0; ww < 20; ww++ {
				if ww+1 == ff {
					fmt.Println(qq[ww])
				}
			}
		} else if len(s) == 2 && (a == "1" || a == "2" || a == "3" || a == "4" || a == "5" || a == "6" || a == "7" || a == "8" || a == "9") && (b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9") {

			fmt.Println(summ(a, b))
		}
	} else if strings.Contains(s, "*") {
		s := strings.Split(s, "*")
		if len(s) > 2 {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			os.Exit(0)
		}
		a := s[0]
		b := s[1]
		if len(s) == 2 && (a == "I" || a == "II" || a == "III" || a == "IV" || a == "V" || a == "VI" || a == "VII" || a == "VIII" || a == "IX" || a == "X") && (b == "I" || b == "II" || b == "III" || b == "IV" || b == "V" || b == "VI" || b == "VII" || b == "VIII" || b == "IX" || b == "X") {
			ff := romanToInt(a) * romanToInt(b)
			qq := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
			for ww := 0; ww < 100; ww++ {
				if ww+1 == ff {
					fmt.Println(qq[ww])
				}
			}
		} else if len(s) == 2 && (a == "1" || a == "2" || a == "3" || a == "4" || a == "5" || a == "6" || a == "7" || a == "8" || a == "9") && (b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9") {

			fmt.Println(umn(a, b))
		}
	} else if strings.Contains(s, "/") {
		s := strings.Split(s, "/")
		if len(s) > 2 {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			os.Exit(0)
		}
		a := s[0]
		b := s[1]
		if len(s) == 2 && (a == "I" || a == "II" || a == "III" || a == "IV" || a == "V" || a == "VI" || a == "VII" || a == "VIII" || a == "IX") && (b == "I" || b == "II" || b == "III" || b == "IV" || b == "V" || b == "VI" || b == "VII" || b == "VIII" || b == "IX") {
			ff := romanToInt(a) / romanToInt(b)
			qq := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI"}
			for ww := 0; ww < 9; ww++ {
				if ww+1 == ff {
					fmt.Println(qq[ww])
				}
			}
		} else if len(s) == 2 && (a == "1" || a == "2" || a == "3" || a == "4" || a == "5" || a == "6" || a == "7" || a == "8" || a == "9") && (b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9") {

			fmt.Println(del(a, b))
		}
	} else if strings.Contains(s, "-") {
		s := strings.Split(s, "-")
		if len(s) > 2 {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			os.Exit(0)
		}
		a := s[0]
		b := s[1]
		if len(s) == 2 && (a == "I" || a == "II" || a == "III" || a == "IV" || a == "V" || a == "VI" || a == "VII" || a == "VIII" || a == "IX") && (b == "I" || b == "II" || b == "III" || b == "IV" || b == "V" || b == "VI" || b == "VII" || b == "VIII" || b == "IX") {
			f := romanToInt(a) - romanToInt(b)
			if f < 0 {
				panic("Выдача паники, так как в римской ссистеме нету отрицательных чисел.")
				os.Exit(0)
			}
			qq := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
			for ww := 0; ww < 9; ww++ {
				if ww+1 == f {
					fmt.Println(qq[ww])
				}
			}
		} else if len(s) == 2 && (a == "1" || a == "2" || a == "3" || a == "4" || a == "5" || a == "6" || a == "7" || a == "8" || a == "9") && (b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9") {

			fmt.Println(minus(a, b))
		}
	} else {
		panic("паника")
	}
}
