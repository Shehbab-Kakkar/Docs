# Ask the user for input and convert it into a list of integers
get_elements = input("Enter the elements of the list separated by spaces: ")
elements_list = list(map(int, get_elements.split()))

# Function to check if any even number is present
def check_even_number(numbers):
    for i in numbers:
        if i % 2 == 0:
            return True
    return False

# Call the function and print the result
print(check_even_number(elements_list))


def square(x):
    return x * x


squares = list(map(square, elements_list))
print(squares)   # Output: [1, 4, 9]


```python
# Ask the user for input and convert it into a list of integers
get_elements = input("Enter the elements of the list separated by spaces: ")
# Example input: 1 2 3 4
elements_list = list(map(int, get_elements.split()))
# .split() turns '1 2 3 4' into ['1', '2', '3', '4']
# map(int, ...) turns each string into an integer: [1, 2, 3, 4]
# elements_list now contains: [1, 2, 3, 4]
# Output after this step (if you print elements_list): [1, 2, 3, 4]

# Function to check if any even number is present
def check_even_number(numbers):
    for i in numbers:
        if i % 2 == 0:    # Checks if the number is even
            return True   # Returns True at the first even number found
    return False          # Returns False if no even number is found

# Call the function and print the result
print(check_even_number(elements_list))
# For input [1, 2, 3, 4], output is:
# True
# (Because 2 is even)

def square(x):
    return x * x          # Returns the square of x

squares = list(map(square, elements_list))
# map(square, elements_list) applies square() to each item:
# [1, 2, 3, 4] -> [1, 4, 9, 16]
print(squares)   # Output: [1, 4, 9, 16]

# -----------------------
# Example Input/Output Walkthrough
# User input: 1 2 3 4
# elements_list: [1, 2, 3, 4]
# print(check_even_number(elements_list))  -> True
# print(squares) -> [1, 4, 9, 16]
```
