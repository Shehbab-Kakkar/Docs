#Get the Element of the List and Print List with the Even Elements
get_input = input("Enter Elements: ")
result = list(map(int,get_input.split()))
even_list=[]
for i in result:
   if i % 2 == 0:
      even_list.append(i)
   else:
      pass

print(even_list)    
