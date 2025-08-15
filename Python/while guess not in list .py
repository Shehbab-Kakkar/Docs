list = ['0','1','2']
def player_guess():
    guess=''
    while guess not in list:
          guess = input("Enter the number: ")
    return int(guess)
print(player_guess())



"""
Enter the number:  0
0
"""
