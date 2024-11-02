# General Notes

- To identify a palindrome have 2 pointers from the middle, one goes towards the start and the other goes towards the end.
- When dealing with a decision tree problem you do tend to get confused between (for loop) and (take or not take method) here's some points to help you
  - usually use for loop method in subsequence problem because if there is an array say [1,2,3] and if you use this method the decision tree will look something like this
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
      // what we are basically doing is this
      // x = 2 ^ 3 ^ 4 ^ 5 ^ 4 ^ 3 ^ 2
      // which is basically this x = 2 ^ 2 ^ 3 ^ 3 ^ 4 ^ 4 ^ 2
      // whic implies x = 5
      return x;
    ```

### Topological Sort

- This algorithm sorts a graph on the basis of some criteria
- Here is how it goes
  - First convert the **(int[][] edges)** to an Adj List mostly of type **(Map\<T, Set\<T>>)** (T = generic)
