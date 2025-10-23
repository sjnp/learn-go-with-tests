package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numsList ...[]int) []int {
	var sums []int
	for _, nums := range numsList {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numsList ...[]int) []int {
	var sums []int
	for _, nums := range numsList {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}

	return sums
}
