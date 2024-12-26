package cmd

func plusOne(digits []int) []int {
    l := len(digits)
    for i, _ := range (digits) {
        last := digits[l-1-i]
        if last + 1 < 10 {
            digits[l-1-i] = last + 1
            break
        } else {
            digits[l-1-i] = 0
            if (l-1-i == 0) {
                digits=append([]int{1},digits...)
            }
        }
    }
    return digits
    
}