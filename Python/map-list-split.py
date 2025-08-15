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
