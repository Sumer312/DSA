# Blind 75

##### This md doc contains the problems that I found very insightful in the [blind 75 list](https://neetcode.io/practice?tab=blind75)

1. [Longest common subsequence](#longest-common-subsequence)
2. [Jump Game](#jump-game)
3. [Merge Intervals](#merge-intervals)
4. [Non overlapping intervals](#non-overlapping-intervals)
5. [Group Anagrams](#group-anagrams)
6. [Course Schedule](#course-schedule)
7. [Minimum window substring](#minimum-window-substring)
8. [Serialize and deserialize binary-tree](#serialize-and-deserialize-binary-tree)
9. [Reorder List](#reorder-list)
10. [Binary Level Order Traversal](#binary-level-order-traversal)
11. [Rotate Image](#rotate-image)
12. [Decode Ways](#decode-ways)
13. [Search in Rotated Sorted Array](#search-in-rotated-sorted-array)
14. [Validate Binary Search Tree](#validate-binary-search-tree)
15. [Longest consecutive sequence](#longest-consecutive-sequence)
16. [3Sum](#3sum)
17. [Product of array except self](#product-of-array-except-self)
18. [Set matrix zeros](#set-matrix-zeros)
19. [Number of 1 bits](#number-of-1-bits)
20. [Sum of two integers](#sum-of-two-integers)
21. [Clone Graph](#clone-graph)
22. [Maximum Product subarray](#maximum-product-subarray)
23. [Binary tree maximum path sum](#binary-tree-maximum-path-sum)
24. [Meeting rooms 2](#meeting-rooms-2)
25. [Implement trie](#implement-trie)
26. [Pacific atlantic water flow](#pacific-atlantic-water-flow)
27. [Word search 2](#word-search-2)
28. [Graph valid tree](#graph-valid-tree)
29. [Construct binary tree from preorder and inorder traversal](#construct-binary-tree-from-preorder-and-inorder-traversal)
30. [Spiral matrix](#spiral-matrix)
31. [House Robber](#house-robber)
32. [Top k frequent elements](#top-k-frequent-elements)
33. [Container with most water](#container-with-most-water)
34. [Longest Repeating Character Replacement](#longest-repeating-character-replacement)
35. [Longest increasing subsequence](#longest-increasing-subsequence)
36. [Kth smallest element in BST]()
37. [Lowest common ancestor of a BST]()
38. [Find Median From Data Stream]()

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

#### [3Sum](https://leetcode.com/problems/3sum/)

- Prerequisites: [Two sum 2](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)
- Explanation:
  - So basically here we sort the input array
  - Then we run a loop from 0 to the second last element of the array (I'll explain why till the second last element)
  - Now we initialize two pointers j = i + 1 and k = nums.length - 1
  - Now for the sum of all three numbers to be zero `(nums[i], nums[j], nums[k])`, sum of two elements should be equal to the negative of the third element
  - Check if the sum of `nums[j] + nums[k]` is equal to `-1 * nums[i]`, if it is then add it to the list and move the two pointers in their respective directions, the reason why we are not breaking here is because there might me more elements whose sum might be equal to `-1 * nums[i]`, for example let's take this array `[2, 3, 0, -3, -2]`, here you have to move the pointers in their respective directions because if you don't do that then the result will be `[0, 2, -2]`, which is wrong.
  - If `sum > -1 * nums[i]` then k-- else j++
  - Now why the outer loop runs from 0 to the second last element is because k is starting at the last element, therefore when i = nums.length - 1 and k = nums.length - 1, it will lead to one element being taken twice. But by only making it run till the second last element and initializing j to i + 1, we are making sure that neither j nor k will be equal to i, this is because j will never be equal to i since it starts at i + 1, and the condition inside the inner while loop is `j < k`, i.e. before k could reach i the inner loop will terminate.
  - Instead of using a `List<List<Integer>>`, use a `HashSet<List<Integer>`, because you will encounter duplicates if you use the first type to store the answers
  - Then convert that hashset to list and return the list

#### [Product of array except self](https://leetcode.com/problems/product-of-array-except-self/)

- Example:
  - ```java
      int[] nums = {2, 3, 4};
    ```
- Explanation:
  - It's can be a little tricky to understand but bear with it, let is sink in properly
  - So to solve this without dividing by zero you can declare 2 arrays named left and right with length nums.length + 1
  - Array left will contain the cumulative product of the array starting from the left
  - Array right will contain the cumulative product of the array starting from the right
  - In array left the element at index 0 is set to 1 and in array right the element the last index is set to 1, because to calculate the cumulative product we must multiply the current element in nums array with the previous element in either the left or right array, depending on which side you are traversing.
  - Considering the above example we get
    ```java
        int[] left = new int[nums.length + 1];
        int[] left = new int[nums.length + 1];
        // ... calculating the cumulative product
        // left = [1, 2, 6, 24]
        // right = [24, 12, 4, 1]
    ```
    - Now declare another result array of size nums.length and to fill that array with the product of array except self element just run a loop from 0 to (i < nums.length) and do this
    ```java
        res[i] = left[i] * right[i + 1];
    ```
    - Think of it like this, the last element of left and the first element of right is the product of all the elements in the array so we ignore those, and if we have to calculate the product of all elements except the current one then we look towards the element at the same index in left array (because it contains the product of all the elements before the current element), and the element at current index + 1 in the right array (because it contains the product all elements after current element).
    - The above approach was the linear space approach, to achieve constant space or O(1) space, we just need to modify the code that we have used for the linear approach.
    - So if you observe carefully, we just need the element at index i from the left array and the element of index i + 1 from the right array to calculate the result.
    - So using that knowledge we can calculate the product and store it in result array at the same time, of course we would have to do it twice because once from the left and next from the right.
    - So just declare the result array of length nums.length, and make `res[0] = 1`
    - Then run a loop from 1 to (i < nums.length) and do this
    ```java
        res[i] = res[i - 1] * nums[i];
    ```
    - Taking the above example into account the res array will become `[1, 2, 6]`
    - Now take the last element from the nums array i.e. 4 and assign it to a variable, say tmp.
    - Now run a loop from (i = nums.length - 2) to (i >= 0), and do this
    ```java
        res[i] *= tmp;
        tmp *= nums[i];
    ```
    - So here we are basically building the right product array and at the same time modifying res.
    - After the loop is done res becomes `[12, 8, 6]`, which is correct
    - If you don't understand the constant space algorithm, pay close attention to the linear one, and then try to understand this, you'll get it.

#### [Set matrix zeros](https://leetcode.com/problems/set-matrix-zeroes/)

- Explanation:

  - The brute force approach to solve this is straightforward, you just create another matrix and set zeros according to the matrix given in the question.
  - Slightly more optimized approach will be to have 2 sets, one to store the rows which will be set to zero and the other to store the columns that will be set to zero. Here you just traverse the matrix and when you encounter that an element is zero store idx i in row set and idx j in column set. Then traverse the sets and make all the rows and columns zero accordingly.
  - To understand the most optimized approach first understand what we are doing in the sets solution, because here we will be doing something similar but we will be using constant space.
  - So instead of taking 2 sets we can use the first column and the first row in the matrix as markers, which will indicate if the row or column should be set to zero or not
  - I will be calling the in place set as markers here

  ```java
  int[][] matrix = {
      {0,1,2,0},
      {3,0,5,2},
      {1,3,1,5}
  };
  [[0,0,0,0],[0,0,0,0],[0,0,1,0]]
  /*
   so here we will use the row {0, 1, 2, 0}, and the column {0, 3, 1} as marker
   so whenever matrix[i][j] == 0, then we set matrix[i][0] = 0 and matrix[0][j] = 0, for instance matrix[1][1] is zero in the above matrix.
   Therefore, we will set matrix[1][0] to 0 and matrix[0][1] to 0
   hence the matrix will become
   {0,0,2,0},
   {0,0,5,2},
   {1,3,1,5}
  */
  ```

  - But there is a problem, as if we use the entire first row to mark the indices of rows that are to be set to zero we will end up using the entire row, which leaves only 2 elements for the column, i.e. there will be a conflict at index `[0][0]`, the element at the index overlaps with both column and row markers.
  - To solve this we will be declaring another variable, say `int colMarker` and set it to one.
  - Now if any element at 0th row is 0, then will will set colMarker to zero.

  ```java
      int colMarker = 1
      int[][] matrix2 = {
          {1,0,2},
          {3,2,5},
          {1,3,1}
      };
      // after traversal colMarker will be set to zero because matrix2[0][1] is 0

  ```

  - Now when we starting setting elements to zero start i from 1 and j from 1, because we are using the row at index `[i][0]` and the column at index `[j][0]` as markers, so setting then first will mess up our marker.
  - After all the elements excepts the markers are set, we will be coming to the marker elements.
  - Check the value at `matrix[0][0]`, if that is zero then make the entire 0th column, i.e. `matrix[0][j]` as zero.
  - Check the value colMarker, if that is zero then make the entire 0th row, i.e. `matrix[i][0]` as zero.
  - The reason it is called colMarker is because it is being used for column, something like this

  ```java
    [a] [1, 2, 3]
        [b]
        [c]
    // here a is the colMarker
    // You could also do this if you want
    [1] [a, 2, 3]
        [b]
        [c]
  ```

#### [Number of 1 bits](https://leetcode.com/problems/number-of-1-bits/)

- Explanation:
  - so the optimized approach is basically that the number of times the loop iterates is the number of 1's that are present in the binary representation of the number.
  - we can achieve this by this forumla
  ```java
      n = n & (n - 1)
  ```
  - run the loop till n > 0
  - lets take the example of number 11
  ```java
      11 = 1011
      // 1st iter
      11 & 10 = 1011 & 1010 = 1010 = 10
      // 2st iter
      10 & 9 = 1010 & 0111 = 0010 = 2
      // 3st iter
      2 & 1 = 010 & 001 = 0
  ```
  - As you can see the loop only ran from 3 iterations and the number 11 has three 1's.
  - Same logic is used in this problem [Counting bits](https://leetcode.com/problems/counting-bits/)

#### [Sum of two integers](https://leetcode.com/problems/sum-of-two-integers/)

- Explanation:
- If we are not to use + or - operator then this is how you add 2 integers
- First keep in mind in addition we have 2 elements, one adding itself, and the other one is to add the carry
- By using the XOR operator we can add the numbers without the carry, i.e.

  ```
      100 (4) + 110 (6) = 1010 (10)
      100 (4) ^ 110 (6) = 010 (2)
  ```

  - In the above example you can see that 2 is not the answer to 4 + 6, but observe care fully, disregarding the MSB, we basically have the result, the only thing left is to add the carry
  - To calculate the carry we can & the numbers

  ```
      100 & 110 = 100
      100 << 1 = 1000 (8)

      1000 ^ 0010 = 1010 (10)
  ```

  - So we just have to and the 2 numbers and shift the result by one place to the left, we do the shifting because the carry is added to the next set of digits, even in normal addition.
  - Here's an example, the numbers are 2 and 3

  ```
     010 ^ 011 = 001
     (010 & 011) << 1 = 0100

     0001 ^ 0100 = 0101
     (0001 & 0100) << 1 = 0000

      herefore ans = 0101 = 5
  ```

  - The code is very short, this is the entire code btw.

  ```java
    int xor = a ^ b;
    int and = (a & b) << 1;
    while(and != 0){
        int tmp = (xor & and) << 1;
        xor = xor ^ and;
        and = tmp;
    }
    return xor;
  ```

  - Don't do `and > 0`, because it will give issues with negative numbers.

#### [Non overlapping intervals](https://leetcode.com/problems/non-overlapping-intervals/)

- Example:
  ```java
      int[][] interval = [[1, 2], [1, 3], [2, 3], [3, 4]];
      // number line visualization
      //             0  1  2  3  4
      // interval[0]    |--|
      // interval[1]    |-----|
      // interval[2]       |--|
      // interval[3]          |--|
  ```
- Explanation:
  - So to solve this problem you first sort the intervals based on the first element, in this approach I use a heap, and initialize a variable named count to 0.
  - Then pop the first element out and assign it to a variable prev
  - Then enter a while loop where start popping the elements
  - Do include this at the end of the loop `prev = cur` where `cur = pq.poll()`
  - If the prev interval and the cur interval overlap then you have to ignore the overlapping interval and increment the `count` variable.
  - Now how do you know out of prev and cur which interval is overlapping, well you know that by looking the end of each interval.
  - Consider 2 intervals `[1, 100]`, and `[1, 26]`
  - Now `prev` could be `[1, 100]` or `[1, 26]`, and the same goes for `cur`, so you cannot just ignore either prev or cur, and then increment the count and move forward, that approach does not work.
  - To properly ignore the overlapping element just compare the ending the of the 2 intervals, and ignore the one where the end is the greatest which in this case is `[1, 100]`
  - After the loop ends return the count
  - It'll look something like this, (this is basically the entire code)
  ```java
    int[] prev = pq.poll();
    int[] cur = new int[2];
    while(!pq.isEmpty()){
        cur = pq.poll();
        if(prev[1] > cur[0]){ // checks for overlap
            count++;
            if(prev[1] > cur[1]){ // checks for which interval to ignore based of the end of each interval
                prev = cur;
            }
            continue;
        }
        prev = cur;
    }
  ```

#### [Clone Graph](https://leetcode.com/problems/clone-graph/)

- Explanation:

  - The first thought that comes to mind to solve this problem is to do a dfs with a hashset
  - But here's the reason why that does not work, take this graph for instance

  ```
  (1)-------(2)
   |         |
   |         |
  (3)-------(4)
  ```

  - Here if you start at one, you create a new node for 1 and you put old 1 in the seen set, the go to 2,

    - then you create a new node for 2 and put 2 in the seen set then go to 4,
    - then you create a new node for 4 put 4 in the seen set then go to 3,
    - you cannot go to 1 from 3 because 1 is already in the seen set so you return back to the call where the current node is 1
    - now when you try to go to 3 from 1, you would not be able to because 3 is already in the seen set
    - therefore new 1 only has new 2 as it's neighbor, because a new node for 3 was never created.

  - So to solve this we use a HashMap and we if `map.containsKey(currentNode)` then we `return map.get(currentNode)`, we are gonna map old one to new one
    ```
        1 -> new Node(1)
    ```
    - then we go to 2 and do the same
    - and same for 4 and 3
    - This is what our map looks like rn btw
    ```
        1 -> new Node(1)
        2 -> new Node(2)
        4 -> new Node(4)
        3 -> new Node(3)
    ```
    - Now when the call stack finally returns to 1 and it goes to 3, then we hit the condition `map.containsKey(3)`, and therefore we return `map.get(3)`, which is a new node with the value 3.
    - and that new node with the value 3 will be added as a neighbor to the new node with the value 1.
  - Therefore the entire graph is cloned.

#### [Maximum Product subarray](https://leetcode.com/problems/maximum-product-subarray/)

- Explanation:
  - This is really not as hard as it looks, and I still don't know why this is classified as a dynamic programming problem
  - But to solve this we just use a modified version of Kadane's algorithm.
  - The big problem with taking out the maximum product subarray is the fact that the array has -ve numbers
  - Take this array for instance `[-1, -2, -3]` here the result is `6` because the product of `[-2, -3]` is 6.
  - So at every step we are gonna calculate the current maximum and the minimum, and the res will be `Math.max(res, current_max)`
  - So to assign the current max and min at every index of the array we do this.
    - At every the max can be either `arr[i]`, `arr[i] * current_max` or `arr[i] * current_min`
    - for example, `arr[i] = 4`, `current_max = -3, current_min = -3`, then the current_max will become 4
    - for example, `arr[i] = -4`, `current_max = 3, current_min = -3`, then the current_max will become `arr[i] * current_min`, i.e. 12
    - for example, `arr[i] = -1`, `current_max = 3, current_min = -3`, then the current_max will remain 3
    - Same for the current minimum as well
  - After calculating the current_max and current_min for the current index in the array you calculate res which will be `Math.max(res, current_max)`
  - The when the loop ends return res

#### [Binary tree maximum path sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/)

- Explanation:

  - To solve this problem we must understand the concept of split, or diversion in a binary tree.
  - So a path in a binary tree can at most have only one split, or one diversion.

  ```
      Valid path          Invalid path
         /\                   /\
        /  \                 / /\
        \   \               /\
            /
  ```

  - In the above diagram the first one is valid because it only has one split at the root, the second one has 3 splits. The reason atmost one split is allowed is because if you start from one end you can reach the other end, and the path does not diverge.
  - So to solve this we are kinda doing a version of Kadane's algorithm. By that I mean at every node we will either be taking the sum of the split (taking the sum of the root, the left subtree, and the right subtree) if it is greater than the sum where we are not splitting.
  - We can do that either by using a global variable or a local variable
  - This is the entire code using a global variable

  ```java
      private static int res;
      private static int dfs(TreeNode root){
          if(root == null){
              return 0;
          }
          int left = Math.max(dfs(root.left), 0);
          int right = Math.max(dfs(root.right), 0);
          //with the split
          res = Math.max(res, root.val + left + right);
          return root.val + Math.max(left, right);
      }
      // In the main function return dfs(root)
  ```

  - This is the entire code using local variables

  ```java
    private static int[] dfs(TreeNode root){
        if(root == null){
            return new int[]{0, Integer.MIN_VALUE};
        }
        int[] left = dfs(root.left);
        int[] right = dfs(root.right);

        left[0] = Math.max(left[0], 0);
        right[0] = Math.max(right[0], 0);

        int skipSplit = root.val + Math.max(left[0], right[0]);
        int takeSplit = Math.max(Math.max(left[1], right[1]), root.val + left[0] + right[0]);
        return new int[]{skipSplit, takeSplit};
    }
  // In the main function return dfs(root)[1]
  ```

  - The approach using global variable is more similar to Kadane's algorithm.
  - The approach using local variables, we are basically returning `int[2]` where the 0th index has the sum without the split, and the 1st index has the maximum path sum
  - And the reason we are doing `Math.max(dfs(root.left), 0)` or `Math.max(left[0], 0)` is because we also have the choice to not add the sum of a subtree, the left subtree in this case, we do this when adding the maximum sum from a subtree will decrease the value from our global maximum.

#### [Meeting rooms 2](https://neetcode.io/problems/meeting-schedule-ii)

- Explanation:
  - The question is asking us to find out how many days it takes to conduct all the meetings without conflicts, or you can see it this way, how many meetings are conducted together, because if 2 meetings are clashing then the number of days it will take to conduct 2 meetings without any conflicts is 2. If the number of meetings are 5, and only 2 meetings are clashing in it, then your answer will also be 2.
  - So we can say that the result will the maximum number of meeting that clashed, or the maximum number of meetings that were going on at one point of time.
  - So to solve this, make a deep copy of the list, sort is based on the start time of the meeting
  - Make another deep copy of the list and sort it based on the end time of the meeting.
  - Start a loop where you have 2 pointers, one at the start of the `start time sorted list` which is '`i`, and one place the `end time sorted list` which is `j`, and also declare 2 variables, `count` and `res`.
  - Now keep incrementing count if `start.get(i).start < end.get(j).end`, and also increment `i`, here `count` basically means how many meetings are current going on.
  - If `end.get(j).end <= start.get(i).start`, decrement `count` and increment `j`.
  - And while you do that also keep track the maximum value reached by `count` and store it in `res`.
  - Return `res`.

#### [Implement trie](https://leetcode.com/problems/implement-trie-prefix-tree/)

- Example
  - The words in the trie are "bad", "app", "apple", "apps"
  ```
  ^ means isEnd
        root
       /    \
      a      b
      |      |
      p      a
      |      |
      p^     d^
     / \
   l^   s^
   |
   e^
  ```
  ```
                                                       root
                                                      (children = {a -> *a, b -> b*}, isEnd = false)
                                           *a         /                                    *b     \
                                          (children = {p -> *p}, isEnd = false)           (children = {a -> *a}, isEnd = false)
                                       *p          /                                       *a       \
                                      (children = {p -> *p}, isEnd = false)               (children = {d -> *d}, isEnd = false)
               *p                    /                                                              *d  \
              (children = {l -> *l, s -> *s}, isEnd = true)                                         (children = {}, isEnd = true)
           *l                /                      *s   \
          (children = {e -> *e}, isEnd = false)    (children = {}, isEnd = true)
      *e             /
     (children = {}, isEnd = true)
  ```
- Explanation
  - The 1st figure shows a very understandable view to a Trie and the second one is actually what is happenning, If you are confused by the second one the just follow the nodes i.e Node of `a` is `*a`, so `*a, *p, *p` make app and therefore we have isEnd set to true at the last `*p`
  - In the 2nd figure, the root node is of custom type `Node` which contains a `HashMap<Character, Node>` and a `boolean` variable.
  - The children of root node contains `a` and `b`, the map.get(character) will give you the corresponding Node.
  - The children of Node `a` contains `p`, and the children of that contains `p` and so on.
  - If the character that the map is pointing to is the last character of the word then the `isEnd` variable is set to true;
  - Also while adding if you are using the recursion and substring approach and you reach the point where the `word.isEmpty()`, the just set the `isEnd` attribute of current node to true.

#### [Pacific atlantic water flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)

- Example:

  ```
        Pacific
      ----------------------
   P  | 1   2   2   3   5* | A
   a  |                    | t
   c  | 3   2   3   4*  4* | l
   i  |                    | a
   f  | 2   4*  5*  3   1  | n
   i  |                    | t
   c  | 6*  7*  1   4*  5* | i
      |                    | c
      | 5*  1   1   2   4* |
      ----------------------
        Atlantic
  ```

- Explanation:
  - In the above diagram the \* represents the cells that can flow water to both the pacific and the atlantic ocean.
  - Now to solve this problem we're gonna need to do some reverse thinking, instead of visiting cells that were less than or equal to the current cell from some point in the center we're gonna start from the borders and visit the cells that are greater than the current cell.
  - When we are visiting greater cells we are going to add them to a HashSet.
  - If we are starting from the pacific border then we create a set named `pacific` and add the cells greater than our starting cell to the `pacific hashset`
  - Same goes for `atlantic`
  - Then we just check for the cells that occur in both of the hashsets, and when we find them we put them in a list and return them.

#### [Word search 2](https://leetcode.com/problems/word-search-ii/)

- Explanation:
  - Now if you would think carefully you will come to the conclusion that this is a trie problem because if we were to check every character in every word in the board our time complexity would have been `O(m.n.k)`, where `m = row.length`, `n = col.length`, and `k = word.length`, instead we could you a trie where the lookup is `O(log(k))`, (the time complexity for that I have given for the problem is completely wrong because I am ignoring recursive calls).
  - First add all the words in the trie.
  - Then for each character in the board call the helper function
  - But if you implement a regular trie and have the below code in the base condition of the helper function then,
    ```java
    if(!trie.search(current_word) || ... || ...){
        return;
    }
    ```
  - You would get a TLE
  - Instead just don't implement a search method and declare class `TrieNode`, outside the `Trie`, class
  - Now in the base condition of the helper function you just check
    ```java
    if(!current_node.children.containsKey(board[r][c]) || ... || ...){
        return;
    }
    current_node = current_node.children.get(board[r][c]);
    ```

#### [Graph valid tree](https://neetcode.io/problems/valid-tree)

- Explanation:
  - If a graph is a valid tree then when we do a DFS on it we will never encounter a node twice
  - We can use this information to solve this problem
  - First build an adj list, and it should be bidirectional, because there are not directions specified
  - Then when we do a regular DFS on the graph with a seen hashset
  - But have both the `current node` and the `parent node` in the parameters of function and before making the recursive call check if the next node which we're about to go is is not parent
  - If that check is not present then recursive call will be made where our `current node` we be the previous call's `parent node`, since our adj list is bidirectional, and it will return false.
  - The same concept of using parent node is applied in [number of components in an undirected graph](https://neetcode.io/problems/count-connected-components).

#### [Construct binary tree from preorder and inorder traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

- Example:

  ```
  Tree
        1     preorder = [1, 2, 4, 5, 3, 6, 7]
      /   \   inorder = [4, 2, 5, 1, 6, 3, 7]
     2     3
    / \   / \
   4   5 6   7

   Subtree
      2    preorder = [2, 4, 5]
     / \   inorder = [4, 2, 5]
    4   5
  ```

- Explanation:
  - Consider the subtree in the example above, by looking at the traversals, we can say that if we take a element in preorder traversal and locate it in inorder traversal then all the elements towards the left of it form the left subtree and all the elements towards the right of it form the right subtree.
  - We can apply the same logic on the whole tree. So if we apply this logic to the full tree.
  - In the code we will be using a `HashMap` to map the inorder values to its indices, and we declare and initialize a `idx` variable to 0
  - Then inside the helper function we will create a new `TreeNode` with the value `preorder[idx]`.
  - Then we will get the index of `preorder[idx]` from the `HashMap` for inorder values.
  - Then we do
  ```java
      root.left = helper(map, preorder, inorder, inorder_start, inorder_index_from_map - 1);
      root.right = helper(map, preorder, inorder, inorder_index_from_map + 1, inorder_end);
  ```
  - And the base conditions will be
  ```java
      if(idx >= preorder.length || inorder_start > inorder_end){
          return null;
      }
  ```

#### [Spiral matrix](https://leetcode.com/problems/spiral-matrix/)

- Example:

  ```
     (l)   (r)
  (t) 1  2  3
      4  5  6
  (b) 7  8  9

  l = left
  r = right
  t = top
  b = bottom
  ```

- Explanation:
  - So here you declare four variables and initialize it to the positions shown in the examples.
  - Then start this loop `while(true)`
  - Then inside the while loop run a for loop from left to right and add while traversing add the elements to the list.
  - Then increment top, because the top row has been added to the list.
  - Then check for this condition
  ```java
    boolean flag = top <= bottom && left <= right;
    if(!flag){
        break;
    }
  ```
  - As show in the code above if the condition is false then break the outer `while` loop.
  - Now the start another for loop from `top` to `bottom` and add the variables to the list.
  - Now increment right.
  - check the condition
  - Same goes on for the bottom row and the left most column.
