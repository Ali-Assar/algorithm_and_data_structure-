# Recursion and Factorial

## Introduction to Recursion

**Recursion** is a method in computer science where a function calls itself in order to solve a problem. A recursive function typically has two parts:
1. **Base Case**: The condition under which the function stops calling itself.
2. **Recursive Case**: The part where the function calls itself with a simpler or smaller input.

The power of recursion lies in the fact that it can break complex problems into smaller, more manageable sub-problems. A classic example to illustrate recursion is the **factorial** function.

### What is a Factorial?

The factorial of a positive integer `n`, denoted by `n!`, is the product of all positive integers less than or equal to `n`. 

\[
n! = n \times (n-1) \times (n-2) \times \dots \times 1
\]

For example:

\[
5! = 5 \times 4 \times 3 \times 2 \times 1 = 120
\]

### Recursive Factorial Function

To compute the factorial using recursion, we break down the problem as follows:

- **Base Case**: When `n = 1`, the factorial of 1 is simply 1. This is the condition where the recursion will stop.
- **Recursive Case**: When `n > 1`, the factorial of `n` can be expressed as:

\[
n! = n \times (n-1)!
\]

This leads to the function calling itself with `n-1` until it reaches the base case.

### How the Recursive Function Works

Let's compute the factorial of 5 recursively:

1. **factorial(5)**: The function is called with `n = 5`. Since `n > 1`, it returns `5 * factorial(4)`.
2. **factorial(4)**: The function is called again with `n = 4`. It returns `4 * factorial(3)`.
3. **factorial(3)**: Now the function is called with `n = 3`, and it returns `3 * factorial(2)`.
4. **factorial(2)**: The function is called with `n = 2` and returns `2 * factorial(1)`.
5. **factorial(1)**: The base case is reached, so the function returns `1`.

Now, the recursion "unwinds" and calculates the final result:

- **factorial(1)** returns `1`
- **factorial(2)** returns `2 * 1 = 2`
- **factorial(3)** returns `3 * 2 = 6`
- **factorial(4)** returns `4 * 6 = 24`
- **factorial(5)** returns `5 * 24 = 120`

Thus, `factorial(5) = 120`.

### Time Complexity

The time complexity of the recursive factorial function is **O(n)**, where `n` is the input value. The function makes `n` recursive calls, each taking constant time.

### Space Complexity

The space complexity is also **O(n)**, due to the stack frames created by the recursive calls. Each call to the function is placed on the call stack until the base case is reached, after which the stack unwinds.

### Visualizing the Recursive Call Stack

Here's a visual representation of how the function works when calculating `factorial(5)`:

```
factorial(5)
|
|--> factorial(4)
      |
      |--> factorial(3)
            |
            |--> factorial(2)
                  |
                  |--> factorial(1)
                        |
                        |--> Base Case: return 1
```

As each function call returns, the values are multiplied and passed back up the stack, ultimately calculating `factorial(5) = 120`.

---

### Conclusion

Recursion is a powerful tool, especially for problems that can be broken down into smaller, repetitive sub-problems like factorials. However, it comes with trade-offs in terms of space complexity due to the stack, which should be considered for deep recursive calls.
