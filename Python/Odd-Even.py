#Program to check if the number is even or not
def even_check(number):
      result = number % 2 == 0
      if result:
        return 'Number is Even'
      else:
        return 'Number is Odd'
      #return result

number=int(input("Enter the Number: "))
print(even_check(number))      
