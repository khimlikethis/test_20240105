package main

import "fmt"

type testcase struct {
	number int
	expect bool
}

func main() {

	firstCase := []int{-1, 4, 30, 2, -4}
	secondCase := []int{ 3, 4, 5, 6, 7}
	thirdCase := []int{3, 4, 5, 6, 7, 7}

	maxValueFirstCase := max(firstCase...)
	maxValueSecondCase := max(secondCase...)

	fmt.Println("A------------------------")

	fmt.Println("The maximum value is:", maxValueFirstCase)
	fmt.Println("The maximum value is:", maxValueSecondCase)

	fmt.Println("A Bonus----------------------")
	secondMaxValueFirstCase := secondMax(firstCase...)
	secondMaxValueSecondCase := secondMax(secondCase...)
	secondMaxValueThirdCase := secondMax(thirdCase...)

	fmt.Println("The secondMax value is:", secondMaxValueFirstCase)
	fmt.Println("The secondMax value is:", secondMaxValueSecondCase)
	fmt.Println("The secondMax value is:", secondMaxValueThirdCase)

	fmt.Println("B----------------------")
	bFirstCase := []int{1, 4, -1, 2, 3}
	fmt.Println("The sum by 3 value is:", fn(bFirstCase, 3))

	bSecondtCase := []int{1, 4, -1, 2, 3}
	fmt.Println("The sum by 3 value is:", fn(bSecondtCase, 2))
}

func max(values ...int) int {
	maxValue := values[0]

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func secondMax(values ...int) int {
	maxValue := 0
	secondMaxValue := 0
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	keepMaxTime := 0

	for _, value := range values {
		if value < maxValue {
			if value > secondMaxValue {
				secondMaxValue = value
			}
		}

		if value == maxValue {
			keepMaxTime++
		}
	}

	if keepMaxTime > 1 {
		return maxValue
	}

	return secondMaxValue
}

func fn(inputs []int, k int) int {
	maxSum := 0
	currentSum := 0

	for i := 0; i < k; i++ {
		currentSum += inputs[i]
	}

	for i := k; i < len(inputs); i++ {
		currentSum = currentSum + inputs[i] - inputs[i-k]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}