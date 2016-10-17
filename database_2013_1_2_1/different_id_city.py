# 统计数据库2013_1_2_1.db中, 相同的city字段下，有多少个不同的id_str值
# author lee
# 2016年08月07日16:05:44
# Linux python3

import os
import sqlite3

def getCityCount(dbname, filename):
	if os.path.exists(filename):
		os.remove(filename)

	f = open(filename, "a")
	conn = sqlite3.connect(dbname)
	cursor = conn.cursor()
	sql_select_table = "select city, count(id_str) as different_id from (select distinct city, id_str from table1) group by city"
	cursor.execute(sql_select_table)
	rows = cursor.fetchall()
	print("Fetchall complete!")
	own_counter = 0 # 计数器，展示进度
	for row in rows:
		f.write(str(row[0]))
		f.write("=")
		f.write(str(row[1]))
		f.write("\n")
		if own_counter > 50000:
			own_counter = 0
			print("50000")
		own_counter = 0

	f.close()
	cursor.close()
	conn.close()
	print("Get city count success!")

def main():
	file_path = "/home/lee/source/cgao_file/get_city_citycount/"
	filename = file_path + "city_citycount.txt"
	db_path = "/home/lee/source/cgao_database/"
	dbname = db_path + "2013_1_2_1.db"
	getCityCount(dbname, filename)

if __name__ == '__main__':
	main()