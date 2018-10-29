import subprocess
import sys  
# def get_cur_info():  
#     print sys._getframe().f_code.co_filename # 当前文件名，可以通过__file__获得  
#     print sys._getframe().f_code.co_name # 当前函数名  
#     print sys._getframe().f_lineno # 当前行号
# get_cur_info()
# p=subprocess.Popen('rsync -avz -e "ssh -p 22" /root/cncc_up/ root@192.168.1.253:/root/',
# shell=True,stdout=subprocess.PIPE,stderr=subprocess.PIPE)
import re
p=subprocess.Popen("ifconfig",stdout=subprocess.PIPE)
ip_list=set()
ip_list.add("localhost")
reip = re.compile(r'(?<![\.\d])(?:\d{1,3}\.){3}\d{1,3}(?![\.\d])')
for line in p.stdout.readlines():
    for ip in reip.findall(line):
        ip_list.add(ip)
if "127.0.0.1" in ip_list:
    print ("yes")