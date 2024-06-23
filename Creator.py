from database import MONSTER_GAME_DB

class Creator():
    def __init__(self) -> None:
        pass

    # Method to create a monster
    def create_monster(self,monster_templateID,name,level,exp,ability,nature,weight,status):
        
        # insert monster into monster table
        MONSTER_GAME_DB.insertMonster(monster_templateID,name,level,exp,ability,nature,weight,status)

        # get inserted monster's monsterID
        monsterID = MONSTER_GAME_DB.getMaxMonsterID()

        # insert active moves data for created monster
        MONSTER_GAME_DB.insertMonsterActiveMovesData(monsterID)

        # insert active moves data for created monster
        MONSTER_GAME_DB.insertMonsterStatsData(monsterID)
    
    # Method to update a monster's active moves
    def update_monster_active_move(self,monsterID,slot,newMoveID):
        MONSTER_GAME_DB.updateActive_moves(monsterID,slot,newMoveID)

editor = Creator()

