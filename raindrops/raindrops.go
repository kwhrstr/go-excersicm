package raindrops

import "strconv"

func Convert (num int) string {
	ans := ""
	if num % 3 == 0 {
		ans += "Pling"
	}
	if num % 5 == 0 {
		ans += "Plang"
	}
	if num % 7 == 0 {
		ans += "Plong"
	}
	if ans == "" {
		ans = strconv.Itoa(num)
	}
	return ans
}