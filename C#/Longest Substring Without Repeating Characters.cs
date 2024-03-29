public class Solution3 {
    public int LengthOfLongestSubstring(string s) {
        
        int max = 0;
        Dictionary<char,int> map = new Dictionary<char, int>();
        
        for(int i=0,j=0;i<s.Length;i++)
        {
            if(map.ContainsKey(s[i]))
            {
                j = Math.Max(j,map[s[i]]+1);
            }    

            map[s[i]] = i;
            max = Math.Max(max,i-j+1);
        }

        return max;
    }
}