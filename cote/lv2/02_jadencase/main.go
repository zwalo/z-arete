package main

import (
	"fmt"
	"strings"
)

/*
JadenCase란 모든 단어의 첫 문자가 대문자이고, 그 외의 알파벳은 소문자인 문자열입니다. 단, 첫 문자가 알파벳이 아닐 때에는 이어지는 알파벳은 소문자로 쓰면 됩니다. (첫 번째 입출력 예 참고)
문자열 s가 주어졌을 때, s를 JadenCase로 바꾼 문자열을 리턴하는 함수, solution을 완성해주세요.

제한 조건
s는 길이 1 이상 200 이하인 문자열입니다.
s는 알파벳과 숫자, 공백문자(" ")로 이루어져 있습니다.
숫자는 단어의 첫 문자로만 나옵니다.
숫자로만 이루어진 단어는 없습니다.
공백문자가 연속해서 나올 수 있습니다.
*/
// func solution(s string) string {
// 	words := strings.Split(s, " ")

// 	for i, word := range words {
// 		runes := []rune(word)
// 		if len(runes) > 0 {
// 			if unicode.IsLetter(runes[0]) {
// 				for i := range runes {
// 					if i == 0 {
// 						runes[i] = unicode.ToUpper(runes[i])
// 					} else {
// 						runes[i] = unicode.ToLower(runes[i])
// 					}
// 				}
// 			} else {
// 				for i := range runes {
// 					runes[i] = unicode.ToLower(runes[i])
// 				}
// 			}
// 		}

// 		words[i] = string(runes)
// 	}

// 	return strings.Join(words, " ")
// }

func solution(s string) string {
	return strings.Title(strings.ToLower(s))
}

func main() {
	fmt.Println(solution("3people unFollowed me"))
	fmt.Println(solution("for the last week"))
}
