# 🧮 Python Program: Check Even Numbers & Square Elements Using Lambda

## 📋 Description

This Python program:

- Takes a list of numbers as input from the user.
- Prints both the string and integer versions of the list.
- Checks whether the list contains any even numbers using the `any()` function.
- Calculates the square of all elements using a `lambda` function and `map()`.

---

## 🚀 How It Works

### 🔢 Input

User is prompted to enter numbers separated by spaces:

Enter the elements: 2 3 4

python
Copy
Edit

### 🧠 Internal Steps

```python
# Step 1: Take input
get_elements = input("Enter the elements: ")

# Step 2: Convert input to list of strings and then to integers
string_list = get_elements.split()
print(string_list)  # ['2', '3', '4']

int_list = list(map(int, string_list))
print(int_list)     # [2, 3, 4]

# Step 3: Check if there's any even number
contains_even = any(num % 2 == 0 for num in int_list)
print("Where the list contain Even or not:")
print(contains_even)  # True

# Step 4: Get square of all elements using lambda and map
squares = list(map(lambda x: x**2, int_list))
print("Give me the square of all Elements of the List")
print(squares)  # [4, 9, 16]
```

---

## 📌 Example Output

```
Enter the elements: 2 3 4
['2', '3', '4']
[2, 3, 4]
Where the list contain Even or not:
True
Give me the square of all Elements of the List
[4, 9, 16]
```

---

## 🔍 Explanation of Key Functions

✅ **map(function, iterable)**  
Applies a function to every element in an iterable.  
Example: `list(map(int, ['1', '2'])) → [1, 2]`

✅ **lambda**  
An anonymous function used for simple operations.  
Example: `lambda x: x**2` squares a number.

✅ **any(iterable)**  
Returns True if any element in the iterable is True.  
Example: `any([False, True, False]) → True`

---
