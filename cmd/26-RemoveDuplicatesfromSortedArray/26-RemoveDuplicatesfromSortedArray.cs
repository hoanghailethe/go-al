public class Solution
{
    public int RemoveDuplicates(int[] nums)
    {
        int i = 1;

        foreach (int n in nums)
        {
            if (nums[i - 1] != n) nums[i++] = n;
        }

        return i;
    }
}