package cmd

import "fmt"

// UNSOLVED

const TARGET_SUM = 0

func threeSum(nums []int) [][]int {

    cacheSumThreeNumbers := make(map[int]map[int]bool)
	res := [][]int{}
    for i, v := range nums{

		targetSum := TARGET_SUM-v

		cacheSumTwoNumbers := map[int]bool{}
		_, ok := cacheSumThreeNumbers[targetSum]
		if ok {
			cacheSumTwoNumbers = cacheSumThreeNumbers[targetSum]
		}

		rm := remove(nums,i)
		pairsNumber, newCacheMap := findTargetSum(targetSum , rm, cacheSumTwoNumbers )
		fmt.Println(pairsNumber)
		// for _, v2 := range pairsNumber {
		// 	v2 = append(v2,v)
		// 	fmt.Println(v2)
		// }
		for i, v2 := range pairsNumber {
			pairsNumber[i] = append(v2,v)
			fmt.Println(pairsNumber[i])
			// add cache three number level:
			for _,v3 := range v2 {
				cacheSumThreeNumbers[TARGET_SUM-v3]= newCacheMap
			}
		}
		fmt.Println(pairsNumber)

		// add cache three number level:
		cacheSumThreeNumbers[targetSum]= newCacheMap
		res = append(pairsNumber, res...)
	}
    return res
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func findTargetSum(sum int, arr[]int, m map[int]bool) ([][]int, map[int]bool) {
	pairs := [][]int{}
	existMap := map[int]bool{}
	
	for _, v := range arr {
		// check if saved pair existed in result 
		// if _, ok := dict["foo"]; ok {
		if _, ok := m[v]; ok{
			continue
		}
	
		// check exist the pair to form Sum target
		need := sum-v
		if _,ok := existMap[need]; ok{ //exist the pair
			pairs = append(pairs, []int{need, v})
			// save to cache result
			m[need]=true
			m[v]=true
		} else {
			existMap[v]=true
		}
	}
	return pairs, m
}