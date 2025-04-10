# Extra problems

##### These are some extra problems that are not in any list per say but are very helpful

1. [Longest palindromic subsequence](#longest-palindromic-subsequence)
2. [Minimum Suffix Flips](#minimum-suffix-flips)
3. [Minimum operations to make binary array elements equal to one 2](#minimum-operations-to-make-binary-array-elements-equal-to-one-2)
4. [Single number 3](#single-number-3)
5. [Remove all adjacent duplicates in string 2](#remove-all-adjacent-duplicates-in-string-2)
6. [Sort characters by frequency](#sort-characters-by-frequency)
7. [Find all duplicates in an array](#find-all-duplicates-in-an-array)
8. [Longest ideal subsequence](#longest-ideal-subsequence)
9. [Subarray with k different integers]()
10. [Flip equivalent binary trees]()
11. [Remove duplicate letters]()
12. [Continuous Subarray Sum]()
13. [Maximum operations to reduce X to zero]()
14. [Rotate array]()
15. [Integer to roman]()
16. [Determine if two strings are close]()
17. [Increasing triplet subsequence]()
18. [Path Sum 3]()
19. [Remove all adjacent duplicates in string 2]()
20. [Find peak element]()
21. [Count vowel strings in ranges]()
<!-- - [Remove all adjacent duplicates in string 1]() -->

#### [Longest palindromic subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/)

- Example: s = "hanna"
  ```
       a n n a h
     h v >
     a *
     n   *
     n     *
     a       *
  ```
- The answer is 4
- Explanation:
  - Same as [Longest common subsequence](#longest-common-subsequence) except the second string is just the first string reversed here

#### [Minimum Suffix Flips](https://leetcode.com/problems/minimum-suffix-flips/)

- Explanation:
  - The starting string you have is consisting of all zeros you have to return the minimum number of moves to convert that to the target string given
  - So initialize 2 variables
  ```java
  int count = 0;
  char curr = '0';
  ```
  - Then traverse through the target string
  - If the current char in target is not equal to curr increment count and set curr to the current character in target
  - After the loop ends return the count

#### [Minimum operations to make binary array elements equal to one 2](https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/)

- Explanation:
  - Similar to minimum suffix flips, only here the premise is reversed, you are given some random string and you need to return the minimum number of operations you convert that string to a string containing all zeros

#### [Single number 3](https://leetcode.com/problems/single-number-iii/)

- Prerequisite: single number 1
- Example array : [1,2,1,3,2,5]
- Explanation:

  - In the first problem there was only one unique number, there there are 2
  - So to solve this do an normal XOR operation on the array and store the result in a variable say x.
  - Then find out the differing bit, what I mean by that is this

    - let's say the 2 unique numbers are 3 and 5,
    - 3 ^ 5 = 6
    - ```java
           3 = 011
           5 = 101
           therefore when we perform an xor we get
           0 1 1
           1 0 1
           -----
           1 1 0

           6 = 110
      ```

    - now here you can see that the differing bit is the MSB and the one to the right of it
    - we just care about finding one differing bit, could be anyone
    - the reason for it is that if we have that bit then every number that has a 1 will be separated into one group and every number that has a 0 will be separated into another group

  - So now we have the XORed bit which is 6
  - Now to find the differing bit we will keep performing an and operation between 1 and 6 until we find the differing bit

    - ```java
        // x = XORed bit i.e. 6
        int d = 1;
        while((x & d) == 0){
            d = d << 1;
        }
        /*
        d = 1   d = 2

        0 0 1   0 1 0
        1 1 0   1 1 0
        -----   -----
        0 0 0   0 1 0
        */
        // from the above operations we know that d = 2;
        // now traverse through the entire array
        int[] res = new int[2];
        Arrays.fill(res, 0);
        for(int i : arr){
            if((d & i) == 0){
                res[0] ^= i;
            } else {
                res[1] ^= i;
            }
        }
      ```

    - The reason why this works is because all 3 has 1 in its 2nd to the MSB place and 5 does not
    - So we just find the bit where both 3 and 5 differ and turn it into an int which is variable d
    - Now pay attention to the if condition. If we use that then res[0] will only perform an xor with the half that has 3 and some duplicates and 5 res[1] will perform xor with 5 and the other duplicates
    - Note that all the duplicate numbers will only be part of one group, i.e. if there are 2 occurrences of 2 then both of those occurrences will be XORed in res[0] only
    - One more thing, if you are confused about doing the AND operation to find the differing bit this might help
      - XOR basically means if a bit of the 2 operands are different then the corresponding bit of the result is 1 else 0
      - So to find which bit is 1 we keep doing an AND operation to the XORed bit until the result is not 0

#### [Remove all adjacent duplicates in string 2](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/)

- Explanation:
  - Use a monotonic stack for it
  - But not just any monotonic stack, the type of the stack should be int[]
  - This is to record the occurrence and store the character in int form i.e. ascii value in the stack
- Example:

  - s = "deeedbbcccbdaa", k = 3
  - This is what will happen in the code

    1. First we keep adding the characters and their occurrence in the stack

    2. Now if the next 'e' is added then the occurrences of 'e' will be equal to k, therefore we need to pop out 'e' and add the elements after that
       | i=3 |
       |:---:|
       | e 2 |
       | e 1 |
       | d 1 |

    3. Now if we add the next 'c' the occurrences of 'c' will be equal k, therefore pop out 'c' and add the next characters
       | i=9 |
       |:---:|
       | c 2 |
       | c 1 |
       | b 2 |
       | b 1 |
       | d 2 |
       | d 1 |

    4. Same with 'b'
       |i=10 |
       |:---:|
       | b 2 |
       | b 1 |
       | d 2 |
       | d 1 |

    5. Same with 'd'
       |i=11 |
       |:---:|
       | d 2 |
       | d 1 |

    6. Now what you are left with is
       |i=13 |
       |:---:|
       | a 2 |
       | a 1 |

  - This is your answer

#### [Sort characters by frequency](https://leetcode.com/problems/sort-characters-by-frequency/)

- Explanation:
  - The heap approach is very simple but is is O(nlogn) because sorting
  - The better approach is bucket sort

#### [Find all duplicates in an array](https://leetcode.com/problems/find-all-duplicates-in-an-array/)

- Explanation
  - This is a very unique problem
  - Here you can use a set and all but to solve this in O(1) space there is a trick
  - The questions mentions that all the elements in the array are in the range [1, len(arr)]
  - So with that information what you can do is use the initial array as a hashset
  - Basically if the current element is 2, then arr[2] should be made negative
  - So that if 2 occurs again in the array we can to to arr[2] and see that the value is already negative that implies that 2 already is present in the array
  - But this only works if the occurrence of the duplicate character is not more that 2
  - This technique will also be used in the problem first missing positive
- Example

```java
int[] arr = {1, 4, 3, 3, 2};
// start traversing the array and start making the value -ve

// arr[1 - 1] is 1, so this is what we do
arr[1 - 1] = -1
// arr[4 - 1] is 2, so this is what we do
arr[4 - 1] = -3
// arr[3 - 1] is 3, so this is what we do
arr[3 - 1] = -3
// arr[3 - 1] is -3, that means 3 has already occurred in the array
```

#### [Longest ideal subsequence](https://leetcode.com/problems/longest-ideal-subsequence/)

- Example:
  ```java
      String s1 = "acfgbd"
      String s2 = "acfgbdhijzlm"
  ```
- Explanation:

  - This kinda helped me understand when to use take or skip in top down and when to use for loop
  - So the reason the recursive calls under for loop do not work is because all the characters in the string are compared to the first character, for example take string s1, in that example every character after `a` is going to be compared to `a`, if the longest ideal subsequence starts from any other character our program will fail because our program is only calculating the longest ideal subsequence that starts at index 0. Take string s2 for instance, here the longest ideal subsequence starts at the character `f`, but the for loop approach gives us the wrong answer when we use it on string s2.
  - So to solve this you need to use take or skip method. This is a 2d dp problem, i.e. there are going to be 2 parameters that we'll be tracking during recursive calls. One is the index and the other is the current character.
  - When calling the top down dfs function make the current character as `#` or any other character and in the function write this

  ```java
    if(Math.abs(s.charAt(idx) - ch) <= k || ch == '#'){
            res = Math.max(res, 1 + dfs(s, k, idx + 1, s.charAt(idx)));
    }
    res = Math.max(res, dfs(s, k, idx + 1, ch));
  ```
