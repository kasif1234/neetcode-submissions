class Solution {
    twoSum(numbers: number[], target: number): number[] {
        let left = 0, right = numbers.length - 1;
        
        while (left < right) {
            const currSum = numbers[left] + numbers[right];
            if (currSum === target) {
                return [left + 1, right + 1];
            } else if (currSum < target) {
                left++;
            } else {
                right--;
            }
        }
        
        return [];
    }
}