from database import MONSTER_GAME_DB
class Monster():

    def __init__(self, monsterID) -> None:

        # load Monster's bio info in 
        self.bioData = MONSTER_GAME_DB.getMonsterData(monsterID)
        self.name = self.bioData[0]
        self.level = self.bioData[1]
        self.exp = self.bioData[2]
        self.ability = self.bioData[3]
        self.nature = self.bioData[4]
        self.weight = self.bioData[5]
        self.status = self.bioData[6]

        # Monster Stats
        self.statsDataRow = MONSTER_GAME_DB.getStatData(monsterID)
        self.stats =   {
                        "health": self.statsDataRow[0],
                        "defence": self.statsDataRow[1],
                        "strength": self.statsDataRow[2],
                        "range": self.statsDataRow[3],
                        "magic": self.statsDataRow[4],
                        "prayer": self.statsDataRow[5],
                        "speed": self.statsDataRow[6]
                        }
        
        # stat derived status
        self.alive == self.stats["health"] > 0 

        # contains moveIDS of current moves
        self.current_moveIDs = MONSTER_GAME_DB.getActive_Moves(monsterID)
        self.current_moves = self.getCurrentMoves()
        
    
    def getStat(self, stat) :
        return self.stats[stat]

    def getCurrentMoves(self):
        result = []
        for i in self.current_moveIDs:

            result.append(MONSTER_GAME_DB.getMoveName(i))
        return result



    def damage(self,moveID):
        # check attackData
        attackData = MONSTER_GAME_DB.getMoveData(moveID)
        attackName = attackData[0]
        attack_power = attackData[2] # base power of attack

        # Calculate damage
        incomingDamage = (attack_power) 
        self.stats["health"] = self.stats["health"] - incomingDamage
        # check if monster is dead

        if self.stats["health"] < 0: # check if overkill
            self.stats["health"] = 0
        if self.stats["health"] == 0:
            self.alive = False
        
        print(f"{self.name} endured {attackName} for {incomingDamage}")

        

            

        




    


    