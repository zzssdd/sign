package utils

import (
	"strconv"
	"strings"
)

func StringToInt64Slice(s string) []int64 {
	numsStrings := strings.Split(s, ",")
	nums := []int64{}
	for _, v := range numsStrings {
		num, _ := strconv.ParseInt(v, 10, 64)
		nums = append(nums, num)
	}
	return nums
}

func Int64SliceToString(nums []int64) string {
	var numsStrings strings.Builder
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			numsStrings.WriteString(",")
		}
		numsStrings.WriteString(strconv.FormatInt(nums[i], 10))
	}
	return numsStrings.String()
}

func AddInt64ToString(s string, num int64) string {
	nums := StringToInt64Slice(s)
	nums = append(nums, num)
	return Int64SliceToString(nums)
}

func StringContainInt64(s string, num int64) bool {
	nums := StringToInt64Slice(s)
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}
