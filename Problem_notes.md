# Neetcode 150
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


[**Jump Game**](https://leetcode.com/problems/jump-game/)
 
* Example nums = [2,3,1,1,4]
* Explaintion: 
    * Just look at the problem differently.
    * Instead of solving the problem from index 0 and going forward, go backwards.
    * By backwards I mean start at the 2nd last index beause that last index is our destination and also set a variable prev = nume.length - 1
    * Now if i + nums[i] >= prev then set prev to i
    * When loop is finished if prev == 0 return true else false
    * This approach has O(1) space and O(n) time
* Another approach with O(n) space and O(nk) time, (k is nums[i])
    ```java
    [2, 3, 1, 1, 4] (int)
    [T, T, T, T, T] (boolean)
    ```
    * Same as the above approach but here you initialize a boolean array and traverse backwards
    * The reason it is O(nk) is because inside the loop we have another loop 
        ```java
        boolean[] flag = new boolean[nums.length]
        flag[flag.length - 1] = true;
        for(int i = nums.length - 2; i >= 0; i++){
            int j = 0;
            while(j < 0){
                if(i + nums[i] - j > nums.length - 1 || flag[i + nums[i] - j]){
                flag[i] = true;
                }
            }
        }
        return flag[0]
        ```
* This approach is just a little easier to understand for new folks

[**Last Stone Weight**](https://leetcode.com/problems/last-stone-weight/)
 
* Explaintion: 
    * Easy problem, just use a max heap
    * In the loop "while(!pq.isEmpty())" pop 2 elements from the heap
    * If after one element is popped the heap becomes empty just return the popped value
    * Else return 0 after loop

[**Rotate Image**](https://leetcode.com/problems/rotate-image/)
 
* Example:
    ``` 
    matrix =

    1  2  3
    4  5  6
    7  8  9
    ```
* Explaintion: 
    * First do a transpose of matrix in-place, (since in the testcases of this problem number of rows is equal to the number of columns of the matrix)
    * To do the transpose of a matix in-place refer notes 
    * This is what the matrix will look like after the transpose
        ``` 
        matrix =

        1  4  7
        2  5  8
        3  6  9
        ```
    * Now if you observe carefully you just have to swap the elements at matrix[i][j] with matrix[i][matrix[i].length - j - 1]
    * And that's it
        ``` 
        matrix =

        7  4  1 
        8  5  2 
        9  6  3 
        ```
