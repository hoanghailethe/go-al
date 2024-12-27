package cmd

import "fmt"

// 100%
func removeElement(nums []int, val int) int {
	count := 0
	right := len(nums)-1
	for i := 0; i < right; i++ {
		for ;right>=0&&nums[right]==val; {
			right--
		}
			
		if i < right {
			tempt := nums[i]
			nums[i] = nums[right]
			nums[right] = tempt
		}
	} 

	for _ , v :=range nums {
		if v == val {
			break
		}
		count++
	}
	fmt.Println(count )
	fmt.Println(nums )
	return count
}

// better
func removeElemen2(nums []int, val int) int {
    left := 0
    right := len(nums) - 1

    for left <= right {
        if nums[left] == val {
            // Swap with the right element
            nums[left] = nums[right]
            right-- // Decrease the right pointer
        } else {
            left++ // Move the left pointer only if no swap occurred
        }
    }

    // The new length is the position of the left pointer
    return left
}