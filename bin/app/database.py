#!/usr/bin/env python3
import sqlite3
from sqlite3 import Error
	
def create_connection(db_file):
    """ create a database connection to a SQLite database """
    conn = None
    try:
        conn = sqlite3.connect(db_file)
        print(sqlite3.version)
    except Error as e:
        print(e)
    finally:
        if conn:
           conn.close()
    return conn

def create_table(conn, create_table_sql):
    """ create a table from the create_table_sql statement
    :param conn: Connection object
    :param create_table_sql: a CREATE TABLE statement
    :return:
    """
    try:
        c = conn.cursor()
        c.execute(create_table_sql)
    except Error as e:
        print(e)
		
if __name__ == '__main__':
    database = r"db//restaurante_a.db"
    sql_create_table= """CREATE TABLE IF NOT EXISTS menu (
                                        id integer PRIMARY KEY,
                                        product text NOT NULL,
                                        subproduct text,
                                        price integer);"""
	
    # create a database connection
    conn = create_connection(database)

    # create tables
    if conn is not None:
        # create projects table
        create_table(conn, sql_create_table)
    else:
        print("Error! cannot create the database connection.")