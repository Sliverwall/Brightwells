import sqlite3

# Debug queries
QUERY_DEBUGGING = False

class SQL_Query():

    def __init__(self, name) -> None:
        self.dbName = f"{name}.db"
    # -------- General queries -----------

    def commitData(self,query):
        try:
            # Establish connection and cursor
            conn = sqlite3.connect(self.dbName)
            cursor = conn.cursor()

            cursor.execute(query)

            if QUERY_DEBUGGING:
                print(f"query:{query}\n executed")
            
            # commit changes
            conn.commit()
            conn.close()
        except Exception as e:
            print(e)
        


    def getDataRow(self, query):
        try:
            # Establish connection and cursor
            conn = sqlite3.connect(self.dbName)
            cursor = conn.cursor()

            cursor.execute(query)
            
            if QUERY_DEBUGGING:
                print(f"query:{query}\n executed")
            dataRow = cursor.fetchone()
            
            conn.close()
            # If a result is found, return the password
            if dataRow:
                return dataRow
            else:
                return None
        except Exception as e:
            print(e)
    