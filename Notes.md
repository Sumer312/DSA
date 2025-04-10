# General Notes

- To identify a palindrome have 2 pointers from the middle, one goes towards the start and the other goes towards the end.
- When dealing with a decision tree problem you do tend to get confused between (for loop) and (take or not take method) here's some points to help you
  - usually use for loop method in subsequence problem because if there is an array say `[1,2,3]` and if you use this method the decision tree will look something like this
    ```
                    [_]
            /        |        \
          [1]       [2]       [3]
        /     \      |
    [1,2]     [1,3] [2,3]
    ```
  - now if you use the (take or not take) method your tree will look something like this
    ```
                           [_]
                /                     \
              [_]                     [1]
            /     \               /         \
        [_]        [2]          [1]        [1,2]
       /   \      /   \        /   \      /     \
     [_]   [3]  [2]   [2,3] [1]  [1,3] [1,2]   [1,2,3]
    ```
  - Please refer to my solution of [Longest ideal subsequence](#longest-ideal-subsequence), in the ExtraProblems md doc, to get more clarity on when to use what.
- A monotonic stack is a stack with all increasing or all decreasing elements
- To solve an anagram problem you need a map

  - First record the occurrences of each character in the first string
  - Then subtract those occurrences from second string
  - If map contains any non zero value return false else true
  - DO NOT DOUBLE THE VALUES WHILE PARSING THE SECOND MAP
  - It will not work on anagrams that have duplicates

- To take out transpose of a matrix in-place do this (transpose of am matrix in-place can only be possible if matrix.length == matrix[0].length)

  - Example:
    ```
    1  2  3
    4  5  6
    7  8  9
    ```
  - Now before you start traversing the matrix create a variable col and initialize it to 0.
  - Now the first for loop will run from 0 to matrix.length - 1
  - The second for loop will run from col to matrix[i].length - 1 and col will increment after every iteration of the first for loop
  - This is how the process looks like

    ```
    i = 0, col = 0, j = col

    1  4  7
    2  5  6
    3  8  9

    i = 1, col = 1, j = col

    1  4  7
    2  5  8
    3  6  9

    i = 2, col = 2, j = col

    1  4  7
    2  5  8
    3  6  9
    ```

# Concepts

### XOR

- XOR stands for "exclusive or" and here is the truth table
- |  a  |  b  | XOR |
  | :-: | :-: | :-: |
  |  0  |  0  |  0  |
  |  0  |  1  |  1  |
  |  1  |  0  |  1  |
  |  1  |  1  |  0  |
- This can be used to find single number in an array which one unique number and rest duplicate numbers
- Example:
  - ```java
      int[] arr = {2,3,4,5,4,3,2};
      int x = 0;
      for(int i : arr){
          x ^= i;
      }
      /* what we are basically doing is this
         x = 2 ^ 3 ^ 4 ^ 5 ^ 4 ^ 3 ^ 2
         which is basically this x = 2 ^ 2 ^ 3 ^ 3 ^ 4 ^ 4 ^ 2
         which implies x = 5 */
      return x;
    ```

### Counting sort

- If the maximum element of the array is not super large, like 500 or something then we can use this sorting algorithm, it is O(n) btw.
- So, let's say that the max of the array is 10, what we do here is that we crate an array named count of size 10, the we start incrementing the elements of count at index arr[i]
- Then traverse the count array and if the value is > 0 then add the element count[i] number of times to the new array
- There cannot be negative elements in the array btw.
- Example:

```java
int[] arr = {3, 2, 5, 1};
int max = 0;
for(int i : arr){
    max = Math.max(i, max);
}
int[] count = new int[max];
for(int i : arr){
    count[i]++;
}
int[] res = new int[arr.length];
int idx = 0;
for(int i = 0; i < count.length; i++){
    int tmp = count[i];
    while(tmp > 0){
        res[idx++] = i;
        tmp--;
    }
}
// res = {1, 2, 3, 5}
```

### Bucket sort

- This is mostly gonna be used to sort the array based on the frequency of the element
- Let's say you want to sort this array by frequency `[1, 2, 2, 4, 4, 3, 3, 3, 3]`
- The result should be `[3, 3, 3, 3, 4, 4, 2, 2, 1]`
- So to solve this using bucket sort what we do is we create a new array.
- And in that array at index 1 we add 1, because 1 has a frequency of 1;
- In the index 2 we add 2 and 4, because both the elements have frequencies of 2 and so on
- So this is what our bucket array will look like

```
[0,  1,    2,    3,  4] //indices
[-, [1], [2, 4], -, [3]] //values
```

- So if we traverse from the back we can sort from the most frequent and if we traverse from the front we can sort based on the least frequent
- The time complexity of the algorithm is O(n)

### Heap sort

- Just sort using a heap

### Topological Sort

- This algorithm sorts a graph on the basis of some criteria
- Here is how it goes
  - First convert the **(int[][] edges)** to an Adj List mostly of type **(Map\<T, Set\<T>>)** (T = generic)

# Identify Patterns

- Algo Monster has a flowchart
- There is also a list on seanprashad leetcode patterns
- If input array is sorted then
  - Binary search
  - Two pointers
- If asked for all permutations/subsets then

  - Backtracking

- If given a tree then

  - DFS
  - BFS

- If given a graph then

  - DFS
  - BFS

- If given a linked list then

  - Two pointers

- If recursion is banned then

  - Stack

- If must solve in-place then

  - Swap corresponding values
  - Store one or more different values in the same pointer

- If asked for maximum/minimum subarray/subset/options then

  - Dynamic programming

- If asked for top/least K items then

  - Heap
  - QuickSelect

- If asked for common strings then

  - Map
  - Trie

- Else
  - Map/Set for O(1) time & O(n) space
  - Sort input for O(nlogn) time and O(1) space
