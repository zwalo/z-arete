package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
문자열 s에는 공백으로 구분된 숫자들이 저장되어 있습니다. str에 나타나는 숫자 중 최소값과 최대값을 찾아 이를 "(최소값) (최대값)"형태의 문자열을 반환하는 함수, solution을 완성하세요.
예를들어 s가 "1 2 3 4"라면 "1 4"를 리턴하고, "-1 -2 -3 -4"라면 "-4 -1"을 리턴하면 됩니다.

제한 조건
s에는 둘 이상의 정수가 공백으로 구분되어 있습니다.
*/
func solution(s string) string {
	slice := strings.Split(s, " ")

	sort.Slice(slice, func(i, j int) bool {
		x, _ := strconv.Atoi(slice[i])
		y, _ := strconv.Atoi(slice[j])

		return x < y
	})

	return fmt.Sprintf("%s %s", slice[0], slice[len(slice)-1])
}

func main() {
	fmt.Println(solution("4 3 27 38"))
}
