txt='adsasadmasdxyjtax'
s=set()
maxLen=0
maxTxt=[]
txtLen=len(txt)
for i in range(txtLen):
    for j in range(txtLen):
        if txt[i]==txt[j]:
            for m in txt[i:j+1]:
                s.add(m)
            if j-i==len(s) and j-i>=maxLen:
                maxLen=j-i
                maxTxt.append(txt[i:j+1])
        s.clear()
print(list(filter(lambda x:len(x)==maxLen+1,maxTxt)))
