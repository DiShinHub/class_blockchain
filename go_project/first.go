package main

import "fmt"

func main() {

	// 배열
	// a := 1
	// num2 := [5]int{a, 2, 3, 4, 5}
	// fmt.Println(num2)

	// fruits := make([]string, 4)

	// fruits[0] = "사과"
	// fruits[1] = "바나나"
	// fruits[2] = "딸기"
	// fruits[3] = "배"

	// fmt.Println(fruits)

	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...) // 이어붙이기

	f4 := append(f1[:2], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)

}
