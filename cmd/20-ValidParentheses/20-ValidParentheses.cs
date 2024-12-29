public class Solution {
    public bool IsValid(string s) {
        var stack = new Stack<char>();
        var samebracket = new Dictionary<char, char> {
            {'(', ')'},
            {'[', ']'},
            {'{', '}'},
        };

        foreach (char ch in s)
        {if (samebracket.ContainsKey(ch)) {
            stack.Push(ch); //mo ngoac thi push to stack
        }
         else 
            if (stack.Count == 0 || samebracket[stack.Pop()] != ch) {
                    return false; // dong ngoac check vs top stack
         }
    }
    return stack.Count == 0;
}
}