public class Solution {
    public int[] PlusOne(int[] digits)  {
        for (int i= digits.Length -1; i>=0; i--  ) {
            if( digits[i]+1 < 10) {
                digits[i] += 1;
            break;
            }
            else {
                digits[i]=0;
                if (i==0){
                    int[] newDigits = new int[digits.Length + 1];
newDigits[0] = 1;                                // set the prepended value
Array.Copy(digits, 0, newDigits, 1, digits.Length); // copy the old values
return newDigits;
                }    
            }
        }
        return digits;
    }
}