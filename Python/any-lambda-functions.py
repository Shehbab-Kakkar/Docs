get_input = input("Enter Elements: ")
result = list(map(int,get_input.split()))
statusOfEvenInList = any(x % 2 == 0 for x in result )
print(f'The status of Even in the list is {statusOfEvenInList}')
statusOfSquareOfListElements = list(map(lambda x: x**2, result))
print(f'Print Square of the list elements \n {statusOfSquareOfListElements}')

"""
Enter Elements:  1 2 3
The status of Even in the list is True
Print Square of the list elements 
 [1, 4, 9]
"""
