# Neetcode 250

##### This md doc contains the problems that I found very insightful in the [neetcode 250 list](https://neetcode.io/practice?tab=neetcode250). Neetcode 250 is the extension of neetcode 150, i.e. the overlapping problems in both lists which were very important are in the Neetcode150 doc.

1. [Maximum sum circular subarray](#maximum-sum-circular-subarray)
2. [Majority Element 2](#majority-element-2)
3. [Find k closest elements](#find-k-closest-elements)

#### [Maximum sum circular subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/)

- Explanation:
  - Very import thing to remember here, when there is an array and you take out the maximum sum subarray from that array, then what you are left with is the minimum sum circular subarray
  - If the maximum sum subarray is circular then the minimum sum subarray is not
  - To solve this problem apply this principle
  - Calculate both the maximum sum subarray and the minimum sum subarray
  - Also calculate the total
  - Return the greater one of these two expressions
    - total - minimum_sum_subarray
    - maximum_sum_subarray
  - And also if the total and the minimum_sum_subarray are equal then just return the maximum_sum_subarray because if they are equal then that implies the entire array consists of only negative numbers and therefore the maximum_sum_subarray will have the largest negative number

#### [Majority Element 2](https://leetcode.com/problems/majority-element-ii/)

- Explanation:
  - This is one of those questions which tough for the sake of being tough
  - So basically it is asking us to calculate the elements which occur in an array nums more than `len(nums) / 3` times
  - By definition that means that the result can only be either 1 or 2 elements because we can have at most 2 elements in an array that occur more than `len(nums) / 3` times
  - Something like this `[1,1,2,2,3]`, here only 1 and 2 are majority elements because they occur 2 times which is more than 5 / 3.
  - So to solve this you need a hashmap of length 3
  - To achieve that we keep filling the hashmap until the size reaches the size 3
  - When is does traverse the hashmap and decrement the elements by 1, and if the occurrence reaches 0 then remove the element
  - By the end what you are left with in the hashmap are 1 or 2 majority elements of the array
  - But hold on, this only works for like ~70% of the time
  - Because for the array like this `[1,2,3,4,5,6,7,8]` it return 7 and 8, which is not the answer
  - Or for the array `[1,2,1,2,1,2,4,5,6]` it returns 5 and 6, which is not the answer
  - So after the loop ends do something like this
  - ```java
      List<Integer> res = new LinkedList();
      /* the code described above */
      for(int i : nums){
          if(verify(nums, i)){
              res.add(i);
          }
      }
      return res;
    ```
  - Where the method verify is checking if the element actually has more than `len(nums) / 3` occurrences or not

#### [Find k closest elements](https://leetcode.com/problems/find-k-closest-elements/)

- Your initial solution for this was absolute ass
