package exercices

/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

//my solution
//BigO O(n^2)

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*

To solve the problem with a time complexity better than O(n²), we can use a hash map (or dictionary) to store the numbers we've already seen along with their indices. This allows us to check if the complement of the current number (the number that, when added to the current number, equals the target) is in the hash map in O(1) time. This reduces the overall time complexity to O(n).

Explanation of the Algorithm (Hash Map Approach):
Create a hash map (or dictionary) to store the numbers we’ve seen so far and their indices.
Iterate through the array:
For each element, compute the complement (i.e., target - nums[i]).
If the complement is already in the hash map, we have found the two numbers, so return their indices.
Otherwise, store the current number and its index in the hash map for future reference.
Return the indices as soon as we find the two numbers that add up to the target.
This approach only requires a single pass through the array, resulting in O(n) time complexity, where n is the number of elements in the array.

*/

func twoSumChatGPT(nums []int, target int) []int {
	// Create a map to store the numbers we've seen and their indices
	seen := make(map[int]int)

	// Loop through the array of numbers
	for i, num := range nums {
		// Calculate the complement (target - current number)
		complement := target - num

		// If the complement is in the map, return the indices
		if index, ok := seen[complement]; ok {
			return []int{index, i}
		}

		// Otherwise, store the current number and its index in the map
		seen[num] = i
	}

	// In case no solution is found (though problem guarantees one solution)
	return []int{}
}
