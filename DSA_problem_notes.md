# Blind 75
[**Longest common subsequence**](https://leetcode.com/problems/longest-common-subsequence/)
 
*  Example: text1 = "abcd" text2 = "affceed"
    ```
         a f f c e e d
       a *              
       b   \ >          
       c   v > *        
       d         \ > *   
    ```
* The answer is 3
* Explaintion: 
    * First you will traverse both the strings
    * If text1.charAt(i) == text2.charAt(j) then increment both in next recursion call
    * Else increment one in one call and increment the other in the other call

[**Longest palindromic subsequence**](https://leetcode.com/problems/longest-palindromic-subsequence/)
 
*  Example: s = "hanna"
    ```
         a n n a h 
       h v >     
       a *         
       n   *       
       n     *     
       a       *   
    ```
* The answer is 4
* Explaintion: 
    * Same as the above question except the second string is just the first string reversed here 

[**Edit Distance**](https://leetcode.com/problems/edit-distance/)
 
* Examplt word1 = "horse", word2 = "ros" 
* Explaintion: 
    * Similar to the two questions above.
    * Now first treverse the two strings using recursion
    * If the characters at indices are equal increment both the indices
    * Else ther are 3 paths you can take
        * First you replace where you do a 1 + dfs(word1, word2, i + 1, j + 1)
        * Second you delete where you do a 1 + dfs(word1, word2, i + 1, j)
        * Third you insert where you do a 1 + dfs(word1, word2, i, j + 1)
    * If (i >= w1.length() && j >= w2.length()) { return 0; }

[**Decode Ways**](https://leetcode.com/problems/decode-ways/)
 
* Examplt s = "226" 
* Explaintion: 
    * There are basically two branches to the decision tree
    * One where you only consider the current character
    * The other where you consider the current character and the next character only if they do not exceed 26 i.e. 'Z'
    * The base considitons are that if idx >= s.length() return 1 
    * And if s.charAt(idx) == '0' return 0 (because the string cannot be decoded if it starts with leading zeros)

[**Course Schedule**](https://leetcode.com/problems/course-schedule/)
 
* Explaintion: 
    * First convert edges array to Map<Integer, Set\<Integer>> map
    * Then do a dfs on the map
        * If dfs returns true create a set named visited and add that node to that set, it means that the current course can be completed 
        * While doing dfs if you encounter a node that has an empty set in map or a node that exists in the visited set (to optimize) then return true
        * While doing dfs if you encounter a node that is present in the seen set then return false because that means a cicrular dependency
    * While doing dfs on nodes 0 to numCourses if you encounter false even once return false else return true
