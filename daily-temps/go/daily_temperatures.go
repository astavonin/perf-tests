package dailytemps

func DailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return result
}

func DailyTemperaturesOpt(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := make([]int, 0, n/4)

	for i := 0; i < n; i++ {
		curr := temperatures[i]
		for len(stack) > 0 {
			top := stack[len(stack)-1] // safely assign once
			if curr <= temperatures[top] {
				break
			}
			stack = stack[:len(stack)-1]
			result[top] = i - top
		}
		stack = append(stack, i)
	}
	return result
}
