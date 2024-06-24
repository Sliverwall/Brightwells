
from Monster import *
class Battle():
    def __init__(self, monster_1, monster_2) -> None:
        # get both monsters
        self.monster_1 = Monster(monster_1)
        self.monster_2 = Monster(monster_2)


    # battle method
    def play(self):
        turnCount = 1
        while self.monster_1.stats["health"] > 0 and self.monster_2.stats["health"] > 0: # while both monsters are alive
            print(f"Turn {turnCount}")
            monster_1_first = self.monster_1.stats["speed"] > self.monster_2.stats["speed"]

            # set turn order based on speed for turn 1
            if turnCount == 1:
                if monster_1_first:
                    turn = True
                else: 
                    turn = False
            # Alternative between turns
            if turn:
                # player 1 attacks
                print(f"{self.monster_1.name} ready to attack!")
                monster_1_moves = self.monster_1.current_moves
                choosen_attack = input(f"choose attack: {monster_1_moves[0]}, {monster_1_moves[1]},{monster_1_moves[2]},{monster_1_moves[3]}")
                self.monster_2.damage(int(choosen_attack))
                turn = False
            else:
                # player 2 attacks
                print(f"{self.monster_2.name} ready to attack!")
                monster_2_moves = self.monster_2.current_moves
                choosen_attack = input(f"choose attack: {monster_2_moves[0]}, {monster_2_moves[1]},{monster_2_moves[2]},{monster_2_moves[3]}")
                self.monster_1.damage(int(choosen_attack))
                turn = True
            # next turn
            turnCount += 1
        if self.monster_1.stats["health"] > 0:
            print("Player 1 wins!")
        else:
            print("Player 2 wins!")
    def otherMethod(self):
        pass