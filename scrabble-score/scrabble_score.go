package scrabble

import (
	"strings"
)

func ScoreLetter(ch rune) int {
	/** https://golang.org/pkg/strings/#ContainsRune
	 *	に書いてある
	 */
	if strings.ContainsRune("AEIOULNRST", ch) {
		return 1
	}
	if strings.ContainsRune("DG", ch) {
		return 2
	}
	if strings.ContainsRune("BCMP", ch) {
		return 3
	}
	if strings.ContainsRune("FHVWY", ch) {
		return 4
	}
	if strings.ContainsRune("K", ch) {
		return 5
	}
	if strings.ContainsRune("JX", ch) {
		return 8
	}
	if strings.ContainsRune("QZ", ch) {
		return 10
	}
	return 0
}

func Score(s string) int {
	point := 0
	/*文字列を大文字にしてloop*/
	for _, ch := range strings.ToUpper(s) { /* maze.goで使っている内容 */
		point += ScoreLetter(ch)
	}
	//別解
	//for _, ch := range s {
	//  とりたしたruneを大文字にしている
	//	point += ScoreLetter(unicode.ToUpper(ch))
	//}
	return point
}