package arrays

func SingleNonDuplicate(nums []int) int {
	if nums[0] != nums[1] {
		return nums[0]
	}

	for i := 2; i <= len(nums)-1; i = i + 2 {
		if i == len(nums)-1 || (nums[i-1] != nums[i] && nums[i+1] != nums[i]) {
			return nums[i]
		}
	}
	return -1
}
