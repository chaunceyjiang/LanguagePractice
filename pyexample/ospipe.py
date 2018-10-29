import os,sys
fd =open("dup2.txt","a+")
os.dup2(fd.fileno(),1)
fd.close()
print "sdsdsds"
print "sdsdsdsdddddddddddd"