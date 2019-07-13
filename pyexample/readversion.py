# import json
#
# with open("version") as f:
#     x = json.loads(f.read())
#     print(x)

import tarfile

tar = tarfile.open("/home/chauncey/BMS/bmsStats/app/web.tar.gz","r:gz")
file_names = tar.getnames()
for file_name in file_names:
    if file_name.endswith(".py") and file_name.startswith("cbms/web"):
        tar.extract(file_name,"/home/chauncey/")
    print(file_name)
tar.close()
