# 🎯 Python Mini Game: Guess the Ball

This is a simple "Guess the Ball" game written in Python.  
It demonstrates Python basics, randomization, user input handling, and list operations.

---

## ✅ What This Program Does

- The program hides a ball (represented by `'0'`) in one of three positions in a list.
- It then shuffles the list randomly, so the ball could be in position 0, 1, or 2.
- The user is asked to guess the position of the ball.
- After the user enters their guess, the program checks whether the guess is correct or not and then reveals the shuffled list.

---

### 🧠 Example

Original list: `['', '0', '']` → ball is in the middle.

After shuffle: `['0', '', '']`

User guesses: `1`

Program says: `"Wrong Guess!"` and shows the shuffled list.

---

## 🎯 Purpose / Goal of the Program

The main purpose of this program is to demonstrate:

- Basic Python syntax and structure
- Use of lists
- Use of randomization (`shuffle`)
- Use of functions
- Getting and validating user input
- Simple conditional logic (`if` statements)

---

## 📚 Subject / Topic of the Program

This program falls under:

- **Python Basics / Fundamentals**
  - Variables, input/output, functions, loops, conditionals
- **Randomization**
  - Using the random module to shuffle data
- **List Operations**
  - Indexing, modifying, and accessing list elements
- **Control Flow**
  - `if-else`, `while` loop for validation

---

## 📝 Sample Code

```python
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
```

---

## 🆕 Want to Improve?

You could:

- Replace `'0'` with `'🎱'` or `'Ball'` for clarity.
- Add a replay loop to let the user play again.
- Keep score of correct guesses.

Let me know if you'd like help adding any of those!

---

## 🔍 If you're labeling this for school or GitHub

You could title it something like:

- "Python Mini Game: Guess the Ball"
- "Beginner Python Project using Lists and Random Module"
- "Python Exercise on List Indexing and User Input"

---
