package d14

// singleNonDuplicateOffice 有序数组中的单一元素官方解法
// 利用双指针来缩小范围，每次找到中值，利用中值附近（理应匹配的）字符进行比较，判断应该向哪个方向切割
// 这才是二分...
func singleNonDuplicateOffice(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
