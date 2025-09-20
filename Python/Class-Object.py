# Define the Dog class
class Dog:
    # A class attribute, shared by all Dog objects
    species = "Canis familiaris"

    # The constructor method to initialize instance attributes
    def __init__(self, name, age):
        self.name = name
        self.age = age

    # An instance method to represent a behavior
    def bark(self):
        return f"{self.name} says woof!"

# Create two objects (instances) of the Dog class
dog1 = Dog("Buddy", 3)
dog2 = Dog("Lucy", 5)

# Access and print attributes of each object
print(f"{dog1.name} is a {dog1.species} who is {dog1.age} years old.")
print(f"{dog2.name} is a {dog2.species} who is {dog2.age} years old.")

# Call the methods of each object
print(dog1.bark())
print(dog2.bark())
