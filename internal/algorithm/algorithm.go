package algorithm

import (
	"strings"
)

//36 SYMBOLS

func GetShortLink(n int) string {
	var res string = ""
	for n > 0 {
		switch mod36 := n % 36; mod36 {
		case 0:
			res = "0" + res
		case 1:
			res = "1" + res
		case 2:
			res = "2" + res
		case 3:
			res = "3" + res
		case 4:
			res = "4" + res
		case 5:
			res = "5" + res
		case 6:
			res = "6" + res
		case 7:
			res = "7" + res
		case 8:
			res = "8" + res
		case 9:
			res = "9" + res
		case 10:
			res = "a" + res
		case 11:
			res = "b" + res
		case 12:
			res = "c" + res
		case 13:
			res = "d" + res
		case 14:
			res = "e" + res
		case 15:
			res = "f" + res
		case 16:
			res = "g" + res
		case 17:
			res = "h" + res
		case 18:
			res = "i" + res
		case 19:
			res = "j" + res
		case 20:
			res = "k" + res
		case 21:
			res = "l" + res
		case 22:
			res = "m" + res
		case 23:
			res = "n" + res
		case 24:
			res = "o" + res
		case 25:
			res = "p" + res
		case 26:
			res = "q" + res
		case 27:
			res = "r" + res
		case 28:
			res = "s" + res
		case 29:
			res = "t" + res
		case 30:
			res = "u" + res
		case 31:
			res = "v" + res
		case 32:
			res = "w" + res
		case 33:
			res = "x" + res
		case 34:
			res = "y" + res
		case 35:
			res = "z" + res

		}
		n = n / 36
	}
	if len(res) <= 4 {
		res = "0/" + res
	} else {
		res = res[0:len(res)-4] + "/" + res[len(res)-4:]
	}
	return res
}

func GetIdFromLink(code string) int {
	slashIndex := strings.Index(code, "/")
	num := code[0:slashIndex] + code[slashIndex+1:]
	rate36 := 1
	res := 0
	var t int
	for i := len(num) - 1; i >= 0; i-- {
		switch string(num[i]) {
		case "0":
			t = 0
		case "1":
			t = 1
		case "2":
			t = 2
		case "3":
			t = 3
		case "4":
			t = 4
		case "5":
			t = 5
		case "6":
			t = 6
		case "7":
			t = 7
		case "8":
			t = 8
		case "9":
			t = 9
		case "a":
			t = 10
		case "b":
			t = 11
		case "c":
			t = 12
		case "d":
			t = 13
		case "e":
			t = 14
		case "f":
			t = 15
		case "g":
			t = 16
		case "h":
			t = 17
		case "i":
			t = 18
		case "j":
			t = 19
		case "k":
			t = 20
		case "l":
			t = 21
		case "m":
			t = 22
		case "n":
			t = 23
		case "o":
			t = 24
		case "p":
			t = 25
		case "q":
			t = 26
		case "r":
			t = 27
		case "s":
			t = 28
		case "t":
			t = 29
		case "u":
			t = 30
		case "v":
			t = 31
		case "w":
			t = 32
		case "x":
			t = 33
		case "y":
			t = 34
		case "z":
			t = 35
		}
		res = res + t*rate36
		rate36 = rate36 * 36
	}
	return res
}
