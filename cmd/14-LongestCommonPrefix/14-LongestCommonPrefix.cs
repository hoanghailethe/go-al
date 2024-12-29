public class Solution {
    public string LongestCommonPrefix(string[] strs) {
        if (strs == null || strs.Length == 0)
        {
            return "";
        }
        string pref = strs[0]; // assume first string la prefix dai nhat
        for (int i = 1; i < strs.Length; i++) // itterate qua cac thu tu con lai trong string
        {
            while (!strs[i].StartsWith(pref)) // compare string vs prefix
            pref = pref.Substring(0, pref.Length - 1); //shorten prefix tu char cuoi
            if (string.IsNullOrEmpty(pref))
                {
                    return ""; // No common prefix found
            }
        }
        return pref;
    }
}