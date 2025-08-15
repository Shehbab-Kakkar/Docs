from random import shuffle

# Shuffle the list
def shuffle_list(mylist):
    shuffle(mylist)
    return mylist

# Ask the player to guess a position
def player_guess():
    guess = ''
    while guess not in ['0', '1', '2']:
        guess = input("Pick a number: 0, 1, or 2: ")
    return int(guess)

# Check if the guess is correct
def check_guess(mylist, guess):
    if mylist[guess] == '0':
        print("Correct! You found the ball!")
    else:
        print("Wrong Guess!")
    print("Shuffled List:", mylist)

# Initial list with the 'ball' hidden in the middle
mylist = ['', '0', '']

# Shuffle the list
mixedup_list = shuffle_list(mylist)

# Get user's guess
guess = player_guess()

# Check if the guess is correct
check_guess(mixedup_list, guess)

