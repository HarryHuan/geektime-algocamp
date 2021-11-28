package main

func main() {

}

type entry struct{ sum, l, r int }

func subarraySum(nums []int, k int) int {
	cnt := 0
	for r := 0; r < len(nums); r++ {
		sum := 0
		for l := r; l >= 0; l-- {
			sum += nums[l]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}
