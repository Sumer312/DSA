# Neetcode 150
#### [Longest common subsequence](https://leetcode.com/problems/longest-common-subsequence/)
 
*  Example: text1 = "abcd" text2 = "affceed"
    ```
         a f f c e e d
       a *              
       b   \ >          
       c   v > *        
       d         \ > *   
    ```
* The answer is 3
* Explanation: 
    * First you will traverse both the strings
    * If text1.charAt(i) == text2.charAt(j) then increment both in next recursion call
    * Else increment one in one call and increment the other in the other call

#### [Longest palindromic subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/)
 
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
* Explanation: 
    * Same as the above question except the second string is just the first string reversed here 

#### [Edit Distance](https://leetcode.com/problems/edit-distance/)
 
* Example: word1 = "horse", word2 = "ros" 
* Explanation: 
    * Similar to the two questions above.
    * Now first traverse the two strings using recursion
    * If the characters at indices are equal increment both the indices
    * Else there are 3 paths you can take
        * First you replace where you do a 1 + dfs(word1, word2, i + 1, j + 1)
        * Second you delete where you do a 1 + dfs(word1, word2, i + 1, j)
        * Third you insert where you do a 1 + dfs(word1, word2, i, j + 1)
    * If (i >= w1.length() && j >= w2.length()) { return 0; }

#### [Decode Ways](https://leetcode.com/problems/decode-ways/)
 
* Example: s = "226" 
* Explanation: 
    * There are basically two branches to the decision tree
    * One where you only consider the current character
    * The other where you consider the current character and the next character only if they do not exceed 26 i.e. 'Z'
    * The base considitons are that if idx >= s.length() return 1 
    * And if s.charAt(idx) == '0' return 0 (because the string cannot be decoded if it starts with leading zeros)

#### [Course Schedule](https://leetcode.com/problems/course-schedule/)
 
* Explanation: 
    * First convert edges array to Map<Integer, Set\<Integer>> map
    * Then do a dfs on the map
        * If dfs returns true create a set named visited and add that node to that set, it means that the current course can be completed 
        * While doing dfs if you encounter a node that has an empty set in map or a node that exists in the visited set (to optimize) then return true
        * While doing dfs if you encounter a node that is present in the seen set then return false because that means a cicrular dependency
    * While doing dfs on nodes 0 to numCourses if you encounter false even once return false else return true


#### [Jump Game](https://leetcode.com/problems/jump-game/)
 
* Example: nums = [2,3,1,1,4]
* Explanation: 
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

#### [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)
 
* Explanation: 
    * Easy problem, just use a max heap
    * In the loop "while(!pq.isEmpty())" pop 2 elements from the heap
    * If after one element is popped the heap becomes empty just return the popped value
    * Else return 0 after loop

#### [Rotate Image](https://leetcode.com/problems/rotate-image/)
 
* Example:
    ``` 
    matrix =

    1  2  3
    4  5  6
    7  8  9
    ```
* Explanation: 
    * First do a transpose of matrix in-place, (since in the test cases of this problem number of rows is equal to the number of columns of the matrix)
    * To do the transpose of a matrix in-place refer notes 
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
#### [Binary Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)
 
```
            (1)   
        /         \ 
      (4)         (7)
    /     \     /                  
  (2)     (9) (8)                

Level order treversal = 
[[1],[4,7],[2,9,8]]
```

* Explanation: 
    * Very simple, initialize a queue of type TreeNode, add root to the queue
    * After that start the loop
    * Inside the loop declare a variable size and initialize it to queue.size()
    * Start a nested loop where the condition is size > 0
    * Add q.poll().val to a list and check if it has any left or right and if it does add those to the queue
    * Decrement the size variable
    * Pseudocode
        ```java
        List<List<Integer>> res = new LinkedList<>();
        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
         while(!q.isEmpty()){
            int size = q.size();
            List<Integer> list = new LinkedList<>();
            while(size > 0){
                TreeNode temp = q.poll();
                list.add(temp.val);
                if(temp.left != null){
                    q.offer(temp.left);
                }
                if(temp.right != null){
                    q.offer(temp.right);
                }
                size--;
            }
            res.add(list);
        }
        ```


#### [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)
 
* First off read the question properly and if you do not understand come back to it later, honestly this was such an easy question that I had just ignored for the longest amount of time because I did not put effort into understanding the question
* Explanation: 
    * The question is just asking you to give the right side view of the tree i.e the right most non null elements
    * The solution is simple, just do a level order traversal and then add the right most elements of the tree, i.e. the last element of every level to a list
        ```
                        (1)   
                    /         \ 
                  (4)         (7)
                /     \     /                  
              (2)     (9) (8)                

          Level order treversal = 
          [[1],[4,7],[2,9,8]]
          
          Right Side View = 
          [1,7,8]
        ```

#### [Reorder List](https://leetcode.com/problems/reorder-list/)
 
* There are 2 approaches to solve this problem, both use linear time, one uses constant space and other uses linear space
    * O(n) space approach
        * Very simple, just traverse the list and add each node into an arraylist
        * Then initialize 2 pointers i and j, i = 0 and j = list.size() - 1
        * Create a new ListNode dummy
        * dummy = new ListNode(0)
        * Now first attach the node at i and increment i and then attach node at j to dummy and decrement j
        * return dummy.next
    * O(1) space approach
        * Find the mid node in the list (fast and slow pointer approach)
        * Then store mid.next in another variable of type ListNode 
        * Now do a dfs on the second half (the variable that stores mid.next) of the list using recursion and initialize a pointer to head 
        * Now when you reach the last node assign last node to pointer.next and move the pointer to pointer.next
        * Now the function will return the the function above in the recursive call where the current node is second last node and do the same

#### [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)
 
* Explanation:
    * Create 2 functions, findPivot() and binarySearch()
        ```
        [4,5,6,7,0,1,2]
               p
        [5,6,7,0,1,2,4]
             p  
        [5,1,3]
         p
        ```
    * Here is how you find the pivot
        *  ```java
            int l = 0, r = nums.length - 1;
            while(l <= r){
                int m = l + (r - l) / 2;
                if(m - 1 >= 0 && nums[m - 1] < nums[m] && m + 1 < nums.length && nums[m + 1] < nums[m]){
                    return m;
                } else if(m - 1 >= 0 && nums[m - 1] > nums[m]){
                    return m - 1;
                } else if(nums[l] > nums[m]){
                    r = m - 1;
                } else {
                    l = m + 1;
                }
            }
            return -1;
            ```
    * Now you search from (0 to p) and from (p + 1 to nums.length - 1)
    
#### [Permutations](https://leetcode.com/problems/permutations/)
* Explanation:
    * Easy backtracking problem, traverse through the array
    * Add the current element to a list and do a recursive call
    * In the next call the loop should start from the beginning ( here there is no idx in the parameter )
    * To make sure that there are no duplicates add this check in the loop
     ```java
         if(list.contains(nums[i])){
             continue;
         }
     ```
     * One optimization you can make is use a hashset instead of a list because the problem specifies that there are no duplicates in the array

#### [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)

* Example s = "abab"
    ```
    list =          
    ["a", "b", "a", "b"]
    ["a", "bab"]
    ["aba", "b"]

    res = 
    [["a", "b", "a", "b"]
    ["a", "bab"]
    ["aba", "b"]]
    ```
* Explanation:
    * Backtracking problem, here we are going to have 2 functions, one is gonna be the recursive function and the other function will check if a string is a palindrome or not
    * The palindrome function is easy to implement
    * The recursive function's parameters are (String s, int idx)
    * Start a for loop where i = idx and inside the loop create a substring from idx to i
    ```java
    String sub = s.substring(idx, i)
    ```
    * If sub is a palindrome add sub to list and do a recursive call to the function and idx = i + 1
    * In the base condition you have
    ```java
    if(idx >= s.length()){
         res.add(new LinkedList(list));
         return;
    }
    ```
#### [Minimum Suffix Flips](https://leetcode.com/problems/minimum-suffix-flips/)
* Explanation:
    * The starting string you have is consisting of all zeros you have to return the minimum number of moves to convert that to the target string given
    * So initialize 2 variables 
    ```java
    int count = 0;
    char curr = '0';
    ```
    * Then traverse through the target string
    * If the current char in target is not equal to curr increment count and set curr to the current character in target
    * After the loop ends return the count
#### [Minimum operations to make binary array elements equal to one 2](https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/)
* Explanation:
    * Similar to minimum suffix flips, only here the premise is reversed, you are given some random string and you need to return the minimum number of operations you convert that string to a string containing all zeros
#### [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)
* Explanation:
    * Here you have to just do an inorder traversal and declare and initialize a variable prev to Long.min_value and when the leftmost leaf or the smallest element is reached you have to check if the prev >= root.val 
    * and if it is the return false
    * else prev = root.val
    * basically here we are checking if the inorder traversal of the tree is sorted
#### [Group Anagrams](https://leetcode.com/problems/group-anagrams/)
* Explanation:
    * First learn how to detect anagrams
    * Here I will refer a string in the string array as word
    * Here you will also require a hashmap of <String, List\<String>\>
    * Now we traverse the string array
    * Then to group them what we do is we use an array of size 26 (because 26 alphabets) as a dict
    * After that we record each instance of a char in the word in the string array
    * Then we convert that array to a string
    * Then we use the hashmap and store the word in the string array in a list with the key being the dict string
    * If any other word has the same dict string then it will be added to them same list that is pointed to by the dict string
    * You cannot use the int[] directly as key because then the map will store the pointer to the array as the key and not the content of the array

* Example:
    ```java
    String s1 = "eat";
    String s2 = "tea";
    ```
    dict string of s1 = "10001000000000000001000000"\
    dict string of s2 = "10001000000000000001000000"

#### [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/)
#### [First-Missing-Positive](https://leetcode.com/problems/first-missing-positive/)
#### [Gas Station](https://leetcode.com/problems/gas-station/)
#### [Valid Parenthesis String](https://leetcode.com/problems/gas-station/)
#### [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)
#### [Subarray with k different intergers](https://leetcode.com/problems/subarrays-with-k-different-integers/)
