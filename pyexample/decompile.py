import os

for root,dirs,files in os.walk("/home/chauncey/BMS/nmsMaser"):
    for name in files:
        os.chdir(root)
        name=str(name)
        if name.endswith("pyc"):
            os.remove(name)
        if name.endswith("pyc_dis"):
            os.rename(name,name.split(".")[0]+".py")