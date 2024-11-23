# Neetcode 150

##### This md doc contains the problems that I found very insightful in the [neetcode 150 list](https://neetcode.io/practice?tab=neetcode150). Neetcode 150 is the extension of blind 75, i.e. the overlapping problems in both lists which were very important are in the Blind75 doc.

1. [Edit Distance](#edit-distance)
2. [Last Stone Weight](#last-stone-weight)
3. [Distinct Subsequences](#distinct-subsequences)
4. [Rotting Oranges](#rotting-oranges)
5. [Koko eating bananas](#koko-eating-bananas)
6. [Binary Tree Right Side View](#binary-tree-right-side-view)
7. [Palindrome Partitioning](#palindrome-partitioning)
8. [Hand of Straights](#hand-of-straights)
9. [Permutations](#permutations)
10. [Generate parentheses](#generate-parentheses)
11. [Jump Game 2](#jump-game-2)
12. [Trapping Rain Water](#trapping-rain-water)
13. [Sliding Window Maximum]()
14. [Task Scheduler]()
15. [Gas Station]()
16. [Partition Labels]()
17. [Coin change 2]()
18. [Valid Parenthesis String]()
19. [First-Missing-Positive]()
20. [Reverse Integer]()

#### [Edit Distance](https://leetcode.com/problems/edit-distance/)

- Example: word1 = "horse", word2 = "ros"
- Explanation:
  - Similar to the two questions above.
  - Now first traverse the two strings using recursion
  - If the characters at indices are equal increment both the indices
  - Else there are 3 paths you can take
    - First you replace where you do a 1 + dfs(word1, word2, i + 1, j + 1)
    - Second you delete where you do a 1 + dfs(word1, word2, i + 1, j)
    - Third you insert where you do a 1 + dfs(word1, word2, i, j + 1)
  - If (i >= w1.length() && j >= w2.length()) { return 0; }

#### [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)

- Explanation:
  - Easy problem, just use a max heap
  - In the loop `while(!pq.isEmpty())` pop 2 elements from the heap
  - If after one element is popped the heap becomes empty just return the popped value
  - Else return 0 after loop

#### [Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/)

- Prerequisite:
  - longest-common-subsequence
- Explanation:
  - One of the easier hard problems
  - So the recursive + memo approach is basically you start with a pointer at both strings
  - ```java
      if(s.charAt(i) == t.charAt(j)){
          dfs(s, t, i + 1, j + 1);
      }
      dfs(s, t, i + 1, j);
    ```
  - This is basically the recursive part
  - The base condition is that if j == t.length() then return true else return false.
  - The reason for it is that, if you observe we are only incrementing j if we have found that a char at i in s is equal to a char at j in t
  - The 2 variables to memoize are i and j

#### [Rotting Oranges](https://leetcode.com/problems/rotting-oranges/)

- Explanation:
  - This is basically BFS on a graph
  - Here you use a queue
  - In that queue store the position of all the rotten oranges and also calculate the number of fresh oranges
  - Now have this condition in you while loop
  - ```java
      while(!queue.isEmpty() && fresh > 0)
    ```
  - Then record the size of the queue (basically what you do in level order tree traversal)
  - Then nest another while loop with this condition
  - ```java
      while(size > 0)
    ```
  - Now pop element from queue
  - Explore all the directions that are valid and if the orange is fresh make it rotten, decrement fresh and add the position of the orange in the queue, since it in now rotten
  - Also decrement size at the end of the nested loop
  - When the nested loop completed increment the res (res here is the time taken to rot all the oranges)
  - return -1 is fresh is not 0 else return res

#### [Koko eating bananas](https://leetcode.com/problems/koko-eating-bananas/)

- Explanation
  - So the brute force way is very simple
    1. Calculate the max value
    2. Then run a for loop from 1 to the max value
    3. In that for loop see with the current value of 'i' the hours it would take to finish the array
    4. You do that by nesting another loop inside the current loop where you traverse the array
    5. Now if the hours you have calculated with 'i' is equal to the hours specified in the params then just return 'i', if multiple 'i's give the same number of hours then return the minimum 'i'
  - Better approach
    1. Calculate the max value
    2. Here instead of running a for loop from 1 to max, observe that this is not an array, it's just a range from 1 to max which means it is sorted which intern means that binary search can be used here
    3. Therefore use binary search for it

#### [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)

- First off read the question properly and if you do not understand come back to it later, honestly this was such an easy question that I had just ignored for the longest amount of time because I did not put effort into understanding the question
- Explanation:

  - The question is just asking you to give the right side view of the tree i.e the right most non null elements
  - The solution is simple, just do a level order traversal and then add the right most elements of the tree, i.e. the last element of every level to a list

    ```
                    (1)
                /         \
              (4)         (7)
            /     \     /
          (2)     (9) (8)

      Level order traversal =
      [[1],[4,7],[2,9,8]]

      Right Side View =
      [1,7,8]
    ```

#### [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)

- Example s = "abab"

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

- Explanation:
  - Backtracking problem, here we are going to have 2 functions, one is gonna be the recursive function and the other function will check if a string is a palindrome or not
  - The palindrome function is easy to implement
  - The recursive function's parameters are (String s, int idx)
  - Start a for loop where i = idx and inside the loop create a substring from idx to i
  ```java
  String sub = s.substring(idx, i)
  ```
  - If sub is a palindrome add sub to list and do a recursive call to the function and idx = i + 1
  - In the base condition you have
  ```java
  if(idx >= s.length()){
       res.add(new LinkedList(list));
       return;
  }
  ```

#### [Hand of Straights](https://leetcode.com/problems/hand-of-straights/)

- Explanation:

  - Here you have to put all the unique digits of the array in a min heap and you also have to record the occurrences of all the digits using a map or an array
  - Now the idea is you peek the top element from the heap and then you run a loop from 0 to groupSize (given in params)
  - Now you pop one digit out of the heap, then you check if then polled digit is equal to the peeked digit + i and if it is not return false
  - Something like this

    ```java
      while(!pq.isEmpty()){
          int peek = pq.peek();
          for(int i = 0; i < groupSize; i++){
              int poll = pq.poll();
              if(poll != peek + i){
                  return false;
              }
          }
      }

    ```

  - Note: You will now get the solution if you yank and paste this code
  - Then you check if the digit has more than one occurrence, and if it does then just put it in a set or something and decrement its occurrence in the occurrence map
  - When the for loop ends add the elements back into the heap that are there in the set where you recorded the digits which had more than one occurrence
  - If everything executes and the while loop is exited return true

#### [Permutations](https://leetcode.com/problems/permutations/)

- Explanation:
  - Easy backtracking problem, traverse through the array
  - Add the current element to a list and do a recursive call
  - In the next call the loop should start from the beginning ( here there is no idx in the parameter )
  - To make sure that there are no duplicates add this check in the loop
  ```java
      if(list.contains(nums[i])){
          continue;
      }
  ```
  - One optimization you can make is use a hashset instead of a list because the problem specifies that there are no duplicates in the array

#### [Generate parentheses](https://leetcode.com/problems/generate-parentheses/)

- Explanation:
  - Here you need to keep track of one thing, and this is a closed parentheses can only be added to the string if an open one already exists
  - So to do that we keep track of two variables 'o' and 'c'
  - Add an open parentheses to the string and decrement o
  - And if c > o i.e. there are more open parentheses in the current string than closed ones then you can add a closed parentheses and decrement c
  - Once both o and c are 0 then add the current string to the result list of type string.
