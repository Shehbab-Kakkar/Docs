import random
import time

class Character:
    def __init__(self, name, health, attack_power, is_zombie=False):
        self.name = name
        self.health = health
        self.attack_power = attack_power
        self.is_zombie = is_zombie

    def attack(self, target):
        damage = random.randint(1, self.attack_power)
        target.health -= damage
        print(f"{self.name} attacks {target.name} for {damage} damage!")

    def is_alive(self):
        return self.health > 0

    def status(self):
        status = "Zombie" if self.is_zombie else "Alive"
        return f"{self.name} - HP: {self.health} - Status: {status}"

def battle(player, enemy):
    print("\n🧟 The battle begins!\n")
    while player.is_alive() and enemy.is_alive():
        print(player.status())
        print(enemy.status())
        print("\nChoose an action:")
        print("1. Attack")
        print("2. Run")

        choice = input(">> ")
        if choice == "1":
            player.attack(enemy)
            if enemy.is_alive():
                enemy.attack(player)
        elif choice == "2":
            print("You run away! Game Over.")
            return
        else:
            print("Invalid choice. Try again.")

        time.sleep(1)
        print("\n" + "-"*30 + "\n")

    if player.is_alive():
        print(f"\n🎉 You defeated the {enemy.name}!")
    else:
        print("\n💀 You have been defeated by the Zombie Ogre...")

def main():
    print("=== Welcome to Zombie Ogre Battle! ===")
    name = input("Enter your hero's name: ")
    player = Character(name=name, health=100, attack_power=20)
    zombie_ogre = Character(name="Zombie Ogre", health=80, attack_power=15, is_zombie=True)

    battle(player, zombie_ogre)

if __name__ == "__main__":
    main()
