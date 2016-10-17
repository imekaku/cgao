import sqlite3

def create_database(dbname):
	conn = sqlite3.connect(dbname)
	cursor = conn.cursor()
	sql_create_table = "create table test(id int, name test)"
	cursor.execute(sql_create_table)
	conn.commit()
	cursor.close()
	conn.close()
	print("Create Database Success!")

def main():
	database_path = "/home/lee/source/cgao_database/"
	dbname = database_path + "test.db"
	create_database(dbname)

if __name__ == '__main__':
	main()