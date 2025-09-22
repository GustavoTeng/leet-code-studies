package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	//param1 := []int{2, 3, 1, 1, 4}
	grid := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1}}

	retorno := numIslands(grid)

	fmt.Println(retorno)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1idx := m - 1
	nums2idx := n - 1
	totalidx := m + n - 1

	for nums2idx >= 0 {
		if nums1idx >= 0 && nums1[nums1idx] > nums2[nums2idx] {
			nums1[totalidx] = nums1[nums1idx]
			nums1idx--
		} else {
			nums1[totalidx] = nums2[nums2idx]
			nums2idx--
		}
		totalidx--
	}
}

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k += 1
		}
	}

	return k
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func removeDuplicatesII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i int = 1

	for j := 1; j < len(nums); j++ {
		if i == 1 || nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func majorityElement(nums []int) int {
	// mapAux := make(map[int]int)
	// maior := 0
	// rep := 0

	// for _, i := range nums {
	// 	mapAux[i] = 1 + mapAux[i]
	// 	if mapAux[i] > maior {
	// 		rep = i
	// 		maior = mapAux[i]
	// 	}
	// }

	// return rep

	maior := 0
	res := 0

	for _, i := range nums {
		if maior == 0 {
			res = i
		}

		if res == i {
			maior++
		} else {
			maior--
		}
	}

	return res
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l

	reverseRotate(nums, 0, l-1)
	reverseRotate(nums, 0, k-1)
	reverseRotate(nums, k, l-1)
}

func reverseRotate(list []int, start int, end int) {
	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}
}

func reverseList(list []any) []any {
	start := 0
	end := len(list) - 1

	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}

	return list
}

func translateHexToString(hexValue string) {
	bytes, err := hex.DecodeString(hexValue)
	if err != nil {
		fmt.Println("Null")
	}

	asciiString := string(bytes)

	fmt.Println("Hexadecimal: \n", hexValue)
	fmt.Println("ASCII: \n", asciiString)
}

func translateStringToHex(stringValue string) {
	bytes := []byte(stringValue)

	hexString := hex.EncodeToString(bytes)

	fmt.Printf("ASCII: %s\n", stringValue)
	fmt.Printf("Hexadecimal: %s\n", hexString)
}

func maxProfit(prices []int) int {
	profit := 0
	minValue := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if profit < prices[i]-minValue {
			profit = prices[i] - minValue
		}
	}

	return profit
}

func maxProfitII(prices []int) int {
	profit := 0
	minValue := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] > minValue {
			profit += prices[i] - minValue
		}
		minValue = prices[i]
	}

	return profit
}

func canJump(nums []int) bool {
	end := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= end {
			end = i
		}
	}

	if end == 0 {
		return true
	} else {
		return false
	}
}

func jump(nums []int) int {
	jumpCount, menorPos, maiorPos := 0, 0, 0

	for maiorPos < len(nums)-1 {
		posMaisLonge := 0
		for i := menorPos; i <= maiorPos; i++ {
			if posMaisLonge < i+nums[i] {
				posMaisLonge = i + nums[i]
			}
		}
		menorPos = maiorPos + 1
		maiorPos = posMaisLonge
		jumpCount++
	}

	return jumpCount
}

func numIslands(grid [][]byte) int {
	isles := 0
	sizeX := len(grid[0])
	sizeY := len(grid)
	backTrack := map[string]bool{}

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if grid[y][x] == 1 && !backTrack[fmt.Sprintf(`%d,%d`, y, x)] {
				isles++
				dfs(x, y, backTrack, grid)
			}
		}
	}

	return isles
}

func dfs(x int, y int, backTrack map[string]bool, grid [][]byte) {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[y][x] == 0 || backTrack[fmt.Sprintf(`%d,%d`, y, x)] {
		return
	}
	backTrack[fmt.Sprintf("%d,%d", y, x)] = true //Put the actual item into a backTracking variable
	dfs(x+1, y, backTrack, grid)                 //Right item
	dfs(x, y+1, backTrack, grid)                 //Down item
	dfs(x-1, y, backTrack, grid)                 //Left item
	dfs(x, y-1, backTrack, grid)                 //Up item
}
