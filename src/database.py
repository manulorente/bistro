#Import libraries
import sqlite3

db_conector = sqlite3.connect('examplebbdd.db')

db_cursor = db_conector.cursor()

item = 1
db_cursor.execute("INSERT INTO producto values((?))", (item,))
db_conector.commit()

db_conector.close()