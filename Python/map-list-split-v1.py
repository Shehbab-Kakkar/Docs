#Python Program To Taken Input as the single space seperated elements 2,3,4
#Convert them to the list. Check if any even of the list has Even or not. Return True
#Take the Square of each element of the list
number=input("Enter the elements:")
print(number.split()) #Convert 1 2 3 to String of List ['1','2','3']
print(list(map(int,number.split()))) #String of List to Int of List [2,3,4]
intlist=list(map(int,number.split()))

def check_even(numlist):
    for i in numlist:
       if i % 2 == 0:
         return True
       else:
         pass

def square(x):
     return x*x
print(f'Where the list contain Even or not:')     
print(check_even(intlist))
print(f'Give me the square of all Elements of the List')
print(list(map(square,intlist)))



""""
Enter the elements: 2 3 4
['2', '3', '4']
[2, 3, 4]
Where the list contain Even or not:
True
Give me the square of all Elements of the List
[4, 9, 16]
""""
