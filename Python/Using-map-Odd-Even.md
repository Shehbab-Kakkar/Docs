# Python `map()` Function Explained

This document explains the Python built-in `map()` function with clear examples, use cases, and how it is helpful in transforming data in lists or other iterables.

---

## 🔹 What is `map()`?

The `map()` function applies a specified function to every item in an iterable (like a list, tuple, or string), returning a map object (which can be converted to a list, set, etc.).

---

## 📝 Syntax

```python
map(function, iterable)
```

- `function`: The function to apply to each item.
- `iterable`: A sequence, like a list, tuple, or string.

---

## 🔍 Example 1: Basic Usage

```python
nums = ['1', '2', '3']
result = map(int, nums)        # Applies int() to each string
print(list(result))            # Output: [1, 2, 3]
```

**Explanation:**  
- `nums` is a list of strings.
- `map(int, nums)` applies `int()` to each string, converting them to integers.
- `list(result)` converts the map object to a list.

---

## 🔎 Example 2: With a Custom Function

```python
def square(x):
    return x * x

nums = [1, 2, 3]
squares = map(square, nums)    # Applies square() to each number
print(list(squares))           # Output: [1, 4, 9]
```

---

## 🔗 Example 3: With Lambda Functions

```python
nums = [2, 4, 6]
doubled = map(lambda x: x * 2, nums)
print(list(doubled))           # Output: [4, 8, 12]
```

---

## 📦 Example 4: Practical Use in Input Processing

Suppose a user inputs numbers as a string:
```python
get_elements = input("Enter numbers: ")        # e.g., "10 15 20"
elements_list = list(map(int, get_elements.split()))
print(elements_list)                           # Output: [10, 15, 20]
```
**How it works:**
- `.split()` converts the string to `['10', '15', '20']`
- `map(int, ...)` converts each to an integer.
- `list(...)` turns the result into a list of integers.

---

## ⚡ Key Benefits of `map()`
- **Cleaner code:** No need for explicit loops.
- **Faster:** Especially for large datasets.
- **Functional programming style:** Makes code more expressive.

---

## 🏷️ When to Use

- When you want to apply the same transformation to every item in a list, tuple, etc.
- When refactoring loops for cleaner, shorter code.

---

## 📚 References

- [Python Docs: map()](https://docs.python.org/3/library/functions.html#map)

---
