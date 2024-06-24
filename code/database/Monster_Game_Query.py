from database.SQL_Query import SQL_Query

class Monster_Game_Query(SQL_Query):

    def __init__(self, name) -> None:
        self.dbName = f"{name}.db"
    
    # ----------------- monsters_template table -------------------------
    def getMonsterTemplateData(self, monster_templateID):
        query = f'''
                       SELECT name,ability_1,ability_2,minimal_weight,max_weight
                       FROM monsters_template
                       WHERE monster_templateID = '{monster_templateID}'
                       '''
        execute_query = self.getDataRow(query)
        return execute_query
    
    # ----------------- monster table -------------------------

    def insertMonster(self,monster_templateID,name,level,exp, ability,nature,weight, status):
        query = f'''
                    INSERT INTO monsters(monster_templateID, name, level,exp,ability,nature,weight,status)
                    VALUES ({monster_templateID}, {name}, {level}, {exp}, {ability}, {nature}, {weight},{status});
                       '''
        
        self.commitData(query)


    def getMonsterData(self, monsterID):
        query = f'''
                       SELECT name, level,exp,ability,nature,weight,status
                       FROM monsters
                       WHERE monsterID = '{monsterID}'
                       '''
        execute_query = self.getDataRow(query)
        return execute_query
    
    def getMonsterName(self, monsterID):
        query = f'''
                       SELECT name
                       FROM monsters
                       WHERE monsterID = '{monsterID}'
                       '''
        execute_query = self.getDataRow(query)[0]
        return execute_query

    def getMonsterTemplateID(self, monsterID):
        query = f'''
                       SELECT monster_templateID
                       FROM monsters
                       WHERE monsterID = '{monsterID}'
                       '''
        execute_query = self.getDataRow(query)[0]
        return execute_query
    
    def getMaxMonsterID(self):
        query = f'''
                       SELECT MAX(monsterID)
                       FROM monsters
                       '''
        execute_query = self.getDataRow(query)[0]
        return execute_query

    # ----------------- stats_template table -------------------------

    def getStatTemplateData(self, monster_TemplateID):
        query = f'''
                       SELECT health,defence,strength,range,magic,prayer,speed
                       FROM stats_template
                       WHERE monster_TemplateID = '{monster_TemplateID}'
                       '''
        execute_query = self.getDataRow(query)
        return execute_query

    # ----------------- stats table -------------------------

    def insertMonsterStatsData(self, monsterID):

        # get monster_templateID of monster
        monster_TemplateID = self.getMonsterTemplateID(monsterID)

        # get the base stats from of the choosen monsterTemplate
        baseStats = self.getStatTemplateData(monster_TemplateID)

        # query to instert base stats into stats table for monster
        query = f'''
                    INSERT INTO stats (monsterID, health, defence,strength,range,magic,prayer,speed)
                    VALUES ({monsterID}, {baseStats[0]},{baseStats[1]}, {baseStats[2]},{baseStats[3]},{baseStats[4]},{baseStats[5]}, {baseStats[6]});
                       '''

        # commit data to db
        self.commitData(query)



    def getStatData(self, monsterID):
        query = f'''
                       SELECT health,defence,strength,range,magic,prayer,speed
                       FROM stats
                       WHERE monsterID = '{monsterID}'
                       '''
        execute_query = self.getDataRow(query)
        return execute_query
    

    # ----------------- moves_template table -------------------------
    def getMoveID(self, name):
        query = f'''
                       SELECT moveID
                       FROM moves_template
                       WHERE name = '{name}'
                       '''
        execute_query = self.getDataRow(query)[0]
        return execute_query

    def getMoveData(self, moveID):
        query = f'''
                       SELECT name,type,power,chance,effect
                       FROM moves_template
                       WHERE move_templateID = '{moveID}'
                       '''
        execute_query = self.getDataRow(query)
        return execute_query
    
    
    def getMoveName(self,moveID):
        query = f'''
                       SELECT name
                       FROM moves_template
                       WHERE move_templateID = '{moveID}'
                       '''
        execute_query = self.getDataRow(query)
        return execute_query

    # ----------------- active_moves_template table -------------------------
    def getActive_Moves_TemplateData(self, monster_TemplateID):
        query = f'''
                       SELECT *
                       FROM active_moves_template
                       WHERE monster_TemplateID = '{monster_TemplateID}'
                       '''
        # do not include template_ids
        execute_query = self.getDataRow(query)[2:]
        return execute_query

    # ----------------- active_moves table -------------------------
    def insertMonsterActiveMovesData(self, monsterID):

        # get monster_templateID of monster
        monster_TemplateID = self.getMonsterTemplateID(monsterID)

        # get the base stats from of the choosen monsterTemplate
        possibleMoves = self.getActive_Moves_TemplateData(monster_TemplateID)

        # query to instert base acitve_moves, defaulting to first 4 for now ****
        query = f'''
                    INSERT INTO active_moves (monsterID, slot1,slot2,slot3,slot4)
                    VALUES ({monsterID}, {possibleMoves[0]}, {possibleMoves[1]}, {possibleMoves[2]},{possibleMoves[3]});
                       '''

        # commit data to db
        self.commitData(query)


    def getActive_Moves(self, monsterID):
        query = f'''
                       SELECT slot1, slot2, slot3, slot4
                       FROM active_moves
                       WHERE monsterID = '{monsterID}'
                       '''
        execute_query = self.getDataRow(query)
        return execute_query
    

    def updateActive_moves(self,monsterID,slot,newMoveID):
        # query to instert base acitve_moves, defaulting to first 4 for now ****
        query = f'''
                    UPDATE active_moves
                    SET {slot} = {newMoveID}
                    WHERE monsterID = {monsterID};
                       '''

        # commit data to db
        self.commitData(query)
    # ----------------- player table -------------------------



    # ----------------- inventory table -------------------------
    def getInventoryData(self,playerID):
        query = f'''
                       SELECT *
                       FROM inventories
                       WHERE playerID = '{playerID}'
                       '''
        # do not include player/inventory ids
        execute_query = self.getDataRow(query)[2:]
        return execute_query
    
