import os
for files in os.listdir("."):
    print files
    files=str(files)
    if "_" in files:
        os.rename(files,files.split("_")[0]+".nts.pack.sz")