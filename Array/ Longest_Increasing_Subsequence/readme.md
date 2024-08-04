# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

Let's walk through finding the Longest Increasing Subsequence (LIS) for the sequence 
    nums = [10, 9, 2, 5, 3, 7, 101, 18]

# Step-by-Step Explanation:
**Initialize the dp array:**
Start with an array dp where each position is initialized to 1, because the smallest increasing subsequence for any single element is the element itself.

```go
dp = [1, 1, 1, 1, 1, 1, 1, 1]
```

# **Process each element:**

- **Element at index 0 (10):**
   
    There are no previous elements to compare with.
    `dp[0] remains 1.`

    ```go
    dp = [1, 1, 1, 1, 1, 1, 1, 1]
    ```


- **Element at index 1 (9):**

        Compare with 10:

        9 < 10, so no update.

        `dp[1] remains 1.`

        ```go
        dp = [1, 1, 1, 1, 1, 1, 1, 1]
        ```
    
- **Element at index 2 (2):**

        Compare with 10 and 9:

        2 < 10, so no update.

        2 < 9, so no update.

        `dp[2] remains 1.`

        ```go
        dp = [1, 1, 1, 1, 1, 1, 1, 1]
        ```
- **Element at index 3 (5):**

        Compare with 10, 9, and 2:

        5 < 10, so no update.

        5 < 9, so no update.

        5 > 2, so update dp[3] to dp[2] + 1 = 2.

        `dp[3] is updated to 2.`

        ```go
        dp = [1, 1, 1, 2, 1, 1, 1, 1]
        ```

- **Element at index 4 (3):**

        Compare with 10, 9, 2, and 5:

        3 < 10, so no update.

        3 < 9, so no update.

        3 > 2, so update dp[4] to dp[2] + 1 = 2.

        3 < 5, so no further update.

        `dp[4] is updated to 2.`


        ```go
        dp = [1, 1, 1, 2, 2, 1, 1, 1]
        ```

- **Element at index 5 (7):**

        Compare with 10, 9, 2, 5, and 3:

        7 < 10, so no update.

        7 < 9, so no update.

        7 > 2, so update dp[5] to dp[2] + 1 = 2.

        7 > 5, so update dp[5] to dp[3] + 1 = 3.

        7 > 3, so no further update.

        `dp[5] is updated to 3.`

        ```go
        dp = [1, 1, 1, 2, 2, 3, 1, 1]
        ```

- **Element at index 6 (101):**

        Compare with 10, 9, 2, 5, 3, and 7:

        101 > 10, so update dp[6] to dp[0] + 1 = 2.

        101 > 9, so no further update.

        101 > 2, so no further update.

        101 > 5, so update dp[6] to dp[3] + 1 = 3.

        101 > 3, so no further update.

        101 > 7, so update dp[6] to dp[5] + 1 = 4.

        `dp[6] is updated to 4.`

        ```go
        dp = [1, 1, 1, 2, 2, 3, 4, 1]

        ```

- **Element at index 7 (18):**

        Compare with 10, 9, 2, 5, 3, 7, and 101:
        
        18 > 10, so update dp[7] to dp[0] + 1 = 2.

        18 > 9, so no further update.

        18 > 2, so no further update.

        18 > 5, so update dp[7] to dp[3] + 1 = 3.

        18 > 3, so no further update.

        18 > 7, so update dp[7] to dp[5] + 1 = 4.

        18 < 101, so no further update.

        
        `dp[7] is updated to 4.`

        ```go
        dp = [1, 1, 1, 2, 2, 3, 4, 4]
        ```


**Final dp array:**

```go
dp = [1, 1, 1, 2, 2, 3, 4, 4]

```

# Result:
**The longest increasing subsequence is the maximum value in the dp array, which is 4.**


# Complexity
- Time complexity:O(n^n)

- Space complexity:O(n)

```go
func lengthOfLIS(nums []int) int {
   if len(nums) == 0 {
		return 0
	}

	// dp array to store the length of the longest increasing subsequence ending at each index
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // each element is an increasing subsequence of length 1
	}

	maxLength := 1

	for j := 1; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
		maxLength = max(maxLength, dp[j])
	}

	return maxLength
}

```
# Minimize Approach

Let's use the example 
```go    
nums := []int{10, 9, 2, 5, 3, 7, 101, 18} 
```
and see how the code works step by step.

**Initialization:**
   stack starts with [10].
**Processing Each Element:**

- **Element 9:**
    binarySearch(stack, 9) returns 0 because 9 should be placed before 10.
    Replace 10 with 9. So, stack becomes [9].

- **Element 2:**
    binarySearch(stack, 2) returns 0 because 2 should replace 9.
    Replace 9 with 2. So, stack becomes [2].

- **Element 5:**
    binarySearch(stack, 5) returns 1 because 5 should be placed at the end of the current stack.
    Append 5 to stack. So, stack becomes [2, 5].

- **Element 3:**
    binarySearch(stack, 3) returns 1 because 3 should replace 5.
    Replace 5 with 3. So, stack becomes [2, 3].

- **Element 7:**
    binarySearch(stack, 7) returns 2 because 7 should be placed at the end of the current stack.
    Append 7 to stack. So, stack becomes [2, 3, 7].

- **Element 101:**
    binarySearch(stack, 101) returns 3 because 101 should be placed at the end of the current stack.
    Append 101 to stack. So, stack becomes [2, 3, 7, 101].
    Element 18:

**binarySearch(stack, 18) returns 3 because 18 should replace 101.
Replace 101 with 18. So, stack becomes [2, 3, 7, 18].**

# Complexity
- Time complexity:O(nlogn)

- Space complexity:O(n)

# Code
```go
func lengthOfLIS(nums []int) int {
if len(nums) == 0 {
		return 0
	}
	temp := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		res, _ := slices.BinarySearch(temp, nums[i])
		if res == len(temp) {
			temp = append(temp, nums[i])
		} else {
			temp[res] = nums[i]
		}
	}
	return len(temp)
}
```