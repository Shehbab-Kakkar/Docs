#Python Program to Shuffle the Elements of the List
from random import shuffle
origList=[1,2,5,45,21,6,77,100]
newList=shuffle(origList)
print(newList)
print(origList)

"""
None
[5, 2, 100, 1, 77, 21, 45, 6]
"""
