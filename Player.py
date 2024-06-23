
from database import MONSTER_GAME_DB
class Player():
    def __init__(self,id) -> None:
        self.playerID = id
        self.inventory = self.getInvetory()
        self.inventoryLabel = self.getInvetoryLabel()


    def getInvetory(self):
        result = []
        inventoryDataRow = MONSTER_GAME_DB.getInventoryData(self.playerID)
        for objectID in inventoryDataRow:
            if objectID:
                result.append(objectID)
        return result
    
    def getInvetoryLabel(self):
        result = []
        inventoryDataRow = self.inventory
        for objectID in inventoryDataRow:
            if objectID:
                result.append(MONSTER_GAME_DB.getMonsterName(objectID))
        return result


