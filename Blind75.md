# Blind 75
##### This md doc contains the problems that I found very insightful in the [blind 75 list](https://neetcode.io/practice?tab=blind75)

- [Longest common subsequence](#longest-common-subsequence)
- [Jump Game](jump-game)
- [Merge Intervals](#merge-intervals)
- [Group Anagrams](#group-anagrams)
- [Course Schedule](#course-schedule)
- [Minimum window substring](#minimum-window-substring)
- [Serialize and deserialize binary-tree](#serialize-and-deserialize-binary-tree)
- [Reorder List](#reorder-list)
- [Binary Level Order Traversal](#binary-level-order-traversal)
- [Rotate Image](rotate-image)
- [Decode Ways](decode-ways)
- [Search in Rotated Sorted Array](#search-in-rotated-sorted-array)
- [Validate Binary Search Tree](#validate-binary-search-tree)
- [Longest consecutive sequence](#longest-consecutive-sequence)
- [3Sum](#3sum)
- [Product of array except self](#product-of-array-except-self)
- [House Robber](#house-robber)
- [Top k frequent elements](#house-robber)

#### [Longest common subsequence](https://leetcode.com/problems/longest-common-subsequence/)

- Example: text1 = "abcd" text2 = "affceed"
  ```
       a f f c e e d
     a *
     b   \ >
     c   v > *
     d         \ > *
  ```
- The answer is 3
- Explanation:
  - First you will traverse both the strings
  - If text1.charAt(i) == text2.charAt(j) then increment both in next recursion call
  - Else increment one in one call and increment the other in the other call

#### [Jump Game](https://leetcode.com/problems/jump-game/)

- Example: nums = [2,3,1,1,4]
- Explanation:
  - Just look at the problem differently.
  - Instead of solving the problem from index 0 and going forward, go backwards.
  - By backwards I mean start at the 2nd last index beause that last index is our destination and also set a variable prev = nume.length - 1
  - Now if i + nums[i] >= prev then set prev to i
  - When loop is finished if prev == 0 return true else false
  - This approach has O(1) space and O(n) time
- Another approach with O(n) space and O(nk) time, (k is nums[i])
  ```java
  [2, 3, 1, 1, 4] (int)
  [T, T, T, T, T] (boolean)
  ```
  - Same as the above approach but here you initialize a boolean array and traverse backwards
  - The reason it is O(nk) is because inside the loop we have another loop
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
- This approach is just a little easier to understand for new folks

#### [Merge Intervals](https://leetcode.com/problems/merge-intervals/)

- FYI: This is the exact same method used to solve the problem [Insert Intervals](https://leetcode.com/problems/insert-interval)
- Explanation
  - So in this problem you need to visualize the intervals in a number line, like
  ```java
      int[][] interval = [[1, 3], [2, 6], [0, 5]];
      // number line visualization
      //             0 1 2 3 4 5 6
      // interval[0]   |---|
      // interval[1]     |-------|
      // interval[2] |--------|
  ```
  - As you can see, using this visualization it is more helpful for us to understand overlapping intervals
  - So to solve this problem we will sort the intervals array on the basis of the first element in the interval, (like using a heap)
  - Here I am following the heap method
  - Then have 2 variables of type `int[]` prev and cur
  - Then we start popping the elements from heap
  - If the current interval is overlapping with the previous interval, then we update the end value of the previous element
  - If the current element does not overlap the previous element the add prev to a list and set prev to cur, (prev = cur)
  - When the loop ends add prev again to list, this is because notice we are only adding prev to the list and not cur, i.e. there will be one interval left after the loop ends
  - So for the example above it will look something like this
    - `[0, 5]` is polled
    - `[1, 3]` is polled, it overlaps with the prev interval so the new prev = `[0, max(5, 3)]` = `[0, 5]`
    - `[2, 6]` is polled and it overlaps with the prev interval which is `[0, 5]`, therefore we update the prev to `[0, max(5, 6)]`
  - Therefore result list contains the interval `[0, 6]`

#### [Group Anagrams](https://leetcode.com/problems/group-anagrams/)

- Explanation:

  - First learn how to detect anagrams
  - Here I will refer a string in the string array as word
  - Here you will also require a hashmap of <String, List\<String>\>
  - Now we traverse the string array
  - Then to group them what we do is we use an array of size 26 (because 26 alphabets) as a dict
  - After that we record each instance of a char in the word in the string array
  - Then we convert that array to a string
  - Then we use the hashmap and store the word in the string array in a list with the key being the dict string
  - If any other word has the same dict string then it will be added to them same list that is pointed to by the dict string
  - You cannot use the int[] directly as key because then the map will store the pointer to the array as the key and not the content of the array

- Example:
  ```java
  String s1 = "eat";
  String s2 = "tea";
  ```
  dict string of s1 = "10001000000000000001000000"\
  dict string of s2 = "10001000000000000001000000"

#### [Course Schedule](https://leetcode.com/problems/course-schedule/)

- Explanation:
  - First convert edges array to Map<Integer, Set\<Integer>> map
  - Then do a dfs on the map
    - If dfs returns true create a set named visited and add that node to that set, it means that the current course can be completed
    - While doing dfs if you encounter a node that has an empty set in map or a node that exists in the visited set (to optimize) then return true
    - While doing dfs if you encounter a node that is present in the seen set then return false because that means a cicrular dependency
  - While doing dfs on nodes 0 to numCourses if you encounter false even once return false else return true

#### [Minimum window substring](https://leetcode.com/problems/minimum-window-substring/)

- Explanation:
  - What you need to to in this problem is to return a substring in string s which has all the characters in the string t, (in any order)
  - To solve this first create 2 hashmaps one from string s, let's call it "ms" and one from string t, call it "mt"
  - Then you record all the occurrences of all the characters in string t
  - Now you start recording the occurrences for all the characters in string s, but there will be two pointers, since sliding window, l and r
  - While recording if all the characters in map mt are present in map ms with the same or greater occurrences then you store the left pointer and right pointer in an array
  - Now nest a loop that checks while the condition is true decrement occurrences from mt at pointer l and increment pointer l
  - When the nested loop breaks the outer loop resumes and the loop will look for another substring that has all the characters in string t
  - If the condition is true once more then check if the substring length is less than the current result, else let it go
  - return s.substring(res[0], res[1])

#### [Serialize and deserialize binary-tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)

- Explanation

  - One of the easier hard problems, here you need good knowledge about how trees are represented in the form of string and also about doing DFS on trees
  - This problem can be solved using BFS as well, but here we will be using the DFS approach
  - So in the serialize method do a regular DFS on the tree and convert it in a string as you go, but keep this points in minimum-window-substring

    - When the root becomes null then add "n" or something else to the string, don't leave it blank
    - Separate all the nodes with a delimiter like a comma (,)

  - So a tree like this

  ```
    1
  /   \
       3
      /  \
     2
  ```

  - becomes this,

  ```java
      String str = "1,n,3,2,n,n,n"
  ```

  - Now to deserialize what we do is initialize a pointer at the start of the array named idx
    - Create new node root and set to the value (int)str.charAt(idx)
    - Increment idx
    - root.left = recursive call
    - Increment idx
    - root.right = recursive call
    - When str.charAt(idx) == 'n' then return null
    - There is also one more way we can do this, instead on incrementing idx after the root.left recursive call we increment inside the base condition where str.charAt(idx) == 'n' is checked, this was in neetcode video but you can just increment after the root.left recursive call as well
    - BTW here I have represented the serialized data as string but when you deserialize it do convert it to a string array where you split by ",", makes the process a whole lot easier

#### [Reorder List](https://leetcode.com/problems/reorder-list/)

- There are 2 approaches to solve this problem, both use linear time, one uses constant space and other uses linear space
  - O(n) space approach
    - Very simple, just traverse the list and add each node into an arraylist
    - Then initialize 2 pointers i and j, i = 0 and j = list.size() - 1
    - Create a new ListNode dummy
    - dummy = new ListNode(0)
    - Now first attach the node at i and increment i and then attach node at j to dummy and decrement j
    - return dummy.next
  - O(1) space approach
    - Find the mid node in the list (fast and slow pointer approach)
    - Then store mid.next in another variable of type ListNode
    - Now do a dfs on the second half (the variable that stores mid.next) of the list using recursion and initialize a pointer to head
    - Now when you reach the last node assign last node to pointer.next and move the pointer to pointer.next
    - Now the function will return the the function above in the recursive call where the current node is second last node and do the same

#### [Binary Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

```
            (1)
        /         \
      (4)         (7)
    /     \     /
  (2)     (9) (8)

Level order trevarsal =
[[1],[4,7],[2,9,8]]
```

- Explanation:
  - Very simple, initialize a queue of type TreeNode, add root to the queue
  - After that start the loop
  - Inside the loop declare a variable size and initialize it to queue.size()
  - Start a nested loop where the condition is size > 0
  - Add q.poll().val to a list and check if it has any left or right and if it does add those to the queue
  - Decrement the size variable
  - Pseudocode
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

#### [Rotate Image](https://leetcode.com/problems/rotate-image/)

- Example:

  ```
  matrix =

  1  2  3
  4  5  6
  7  8  9
  ```

- Explanation:

  - First do a transpose of matrix in-place, (since in the test cases of this problem number of rows is equal to the number of columns of the matrix)
  - To do the transpose of a matrix in-place refer notes
  - This is what the matrix will look like after the transpose

    ```
    matrix =

    1  4  7
    2  5  8
    3  6  9
    ```

  - Now if you observe carefully you just have to swap the elements at matrix[i][j] with matrix[i]matrix[i].length - j - 1]
  - And that's it

    ```
    matrix =

    7  4  1
    8  5  2
    9  6  3
    ```

#### [Decode Ways](https://leetcode.com/problems/decode-ways/)

- Example: s = "226"
- Explanation:
  - There are basically two branches to the decision tree
  - One where you only consider the current character
  - The other where you consider the current character and the next character only if they do not exceed 26 i.e. 'Z'
  - The base considitons are that if idx >= s.length() return 1
  - And if s.charAt(idx) == '0' return 0 (because the string cannot be decoded if it starts with leading zeros)

#### [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

- Explanation:
  - Create 2 functions, findPivot() and binarySearch()
    ```
    [4,5,6,7,0,1,2]
           p
    [5,6,7,0,1,2,4]
         p
    [5,1,3]
     p
    ```
  - Here is how you find the pivot
    - ```java
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
  - Now you search from (0 to p) and from (p + 1 to nums.length - 1)

#### [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)

- Explanation:
  - Here you have to just do an inorder traversal and declare and initialize a variable prev to Long.min_value and when the leftmost leaf or the smallest element is reached you have to check if the prev >= root.val
  - and if it is the return false
  - else prev = root.val
  - basically here we are checking if the inorder traversal of the tree is sorted

#### [Longest consecutive sequence](https://leetcode.com/problems/longest-consecutive-sequence/)

- Example:
  - ```java
      int[] nums = {100,4,200,1,3,2};
      // res = 4
      // because the longest consecutive sequence is 1, 2, 3, 4
    ```
- Explanation:
  - They are asking you so solve this in O(n)
  - So just add the entire array to a hashset named "set"
  - Now initialize another hashset named "seen"
  - Now traverse "set" and check if there exists +1 of current element or -1 of current element
  - If that is the case do a count++ and keep checking until there does not exist a +1 of current element or -1 of current element
  - Meanwhile keep adding all the elements you've visited in the "seen" set and if the current element is in "seen" continue
  - Return count
