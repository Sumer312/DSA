# General Notes
* To identify a palindrome have 2 pointers from the middle, one goes towards the start and the other goes towards the end.
* When dealing with a decison tree problem you do tend to get confused between (for loop) and (take or not take method) here's some points to help you
    * usually use for loop method in subsequence problem because if there is an array say [1,2,3] and if you use this method the decision tree will look something like this
        ```
                        [_]   
                /        |        \ 
              [1]       [2]       [3]
            /     \      |
        [1,2]     [1,3] [2,3]
        ```
    
    * now if you use the (take or not take) method your tree will look something like this
        ```
                               [_]  
                    /                     \
                  [_]                     [1]
                /     \               /         \
            [_]        [2]          [1]        [1,2]
           /   \      /   \        /   \      /     \    
         [_]   [3]  [2]   [2,3] [1]  [1,3] [1,2]   [1,2,3]
        ```
* A monotonoic stack is a stack with all increasing or all decreasing elements
* To solve an anagram problem you need a map
    * first record the occursences of each character in the first string
    * then substract those occursences from second string
    * if map contains any non zero value return false else true
    * DO NOT DOUBLE THE VALUES WHILE PARSING THE SECOND MAP
    * it will not work on anagrams that have duplicates

* To take out transpose of a matrix in-place do this (transpose of am matrix in-place can only be possible if matrix.length == matrix[0].length)
    * Example:
        ``` 
        1  2  3
        4  5  6
        7  8  9
        ```
    * Now before you start traversing the matrix create a varible col and initiliaze it to 0.
    * Now the first for loop will run from 0 to matrix.length - 1
    * The second for loop will run from col to matrix[i].length - 1 and col will increment after every iteration of the first for loop
    * This is how the process looks like
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
### Topological Sort
* This algorithm sorts a graph on the basis of some creiteria
* Here is how it goes
    * First convert the **(int[][] edges)** to an Adj List mostly of type **(Map\<T, Set\<T>>)** (T = generic)


