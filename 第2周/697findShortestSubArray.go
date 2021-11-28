package main

func main() {

}

type entry struct{ cnt, l, r int }

func findShortestSubArray(nums []int) int {
	mp := map[int]entry{}
	for i, v := range nums {
		if e, ok := mp[v]; ok {
			e.cnt++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{cnt: 1, l: i, r: i}
		}
	}
	maxCnt := 0
	ans := len(nums)
	for _, v := range mp {
		if v.cnt > maxCnt {
			maxCnt, ans = v.cnt, v.r-v.l+1
		} else if v.cnt == maxCnt {
			ans = min(ans, v.r-v.l+1)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
