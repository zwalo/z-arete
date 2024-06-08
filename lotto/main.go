package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	// 상수 정의
	mandatoryNumbers := []int{27, 38}
	optionalNumbers := getBestNumbers()
	allNumbers := make([]int, 45)
	for i := 1; i <= 45; i++ {
		allNumbers[i-1] = i
	}

	// 랜덤 시드 설정
	rand.Seed(time.Now().UnixNano())

	// 5세트 생성
	for i := 0; i < 5; i++ {
		set := generateLottoSet(mandatoryNumbers, optionalNumbers, allNumbers)
		fmt.Println(set)
	}
}

func getBestNumbers() []int {
	// CSV 파일 열기
	bestNumbers := make([]int, 6)
	file, err := os.Open("lotto.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return bestNumbers
	}
	defer file.Close()

	// CSV 읽기
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return bestNumbers
	}

	// 번호 출현 빈도를 저장할 맵
	frequency := make(map[int]int)

	// CSV 데이터 처리
	for i, record := range records {
		if i == 0 {
			continue
		}

		for _, value := range record[1:] {
			number, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting value to int:", err)
				//return
			}
			frequency[number]++
		}
	}

	// 빈도수를 기반으로 번호를 정렬
	type kv struct {
		Key   int
		Value int
	}

	var freqList []kv
	for k, v := range frequency {
		freqList = append(freqList, kv{k, v})
	}

	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i].Value > freqList[j].Value
	})

	for i := 0; i < 6; i++ {
		bestNumbers[i] = freqList[i].Key
	}
	return bestNumbers
}

func generateLottoSet(mandatoryNumbers, optionalNumbers, allNumbers []int) []int {
	chosenNumbers := make(map[int]bool)
	// 필수 숫자 추가
	for _, num := range mandatoryNumbers {
		chosenNumbers[num] = true
	}

	// 선택 숫자에서 2개 추가
	for len(chosenNumbers) < 4 {
		num := optionalNumbers[rand.Intn(len(optionalNumbers))]
		chosenNumbers[num] = true
	}

	// 나머지 숫자 추가
	for len(chosenNumbers) < 6 {
		num := allNumbers[rand.Intn(len(allNumbers))]
		if !chosenNumbers[num] {
			chosenNumbers[num] = true
		}
	}

	// 결과 배열 생성
	var result []int
	for num := range chosenNumbers {
		result = append(result, num)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
