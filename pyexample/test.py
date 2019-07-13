
import pymsql
connection = pymysql.connect(
    host="192.168.1.202",
    port=4000,
    user="root",
    password="",
    charset='utf8mb4',
    database="test")


cur = connection.cursor()
cur.excute("select * from c2019_000000_611_1")


cache = cur.fetchall()
print(len(cache))



