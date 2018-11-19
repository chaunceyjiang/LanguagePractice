import time
def read_file(filename):
    
    with open(filename,"r") as f:
        while True:
            yield f.readline()
cache = {}
records = []
count = 0
for line in read_file('/home/chauncey/Desktop/1825566008_trac'):
    try:
        if line != '\n':
            record = eval(line)
            if record[0]['transaction_id']=="2018102618493216":
                records.append(record)
    except:
        break
        # ms_key = record[0]['ms_key'] 
        # # if tuple(ms_key) not in cache and ms_key[2] =='M':
        # if ms_key[2] =='M':
        #     count+=1
        #     if ms_key[1] == '2018102618493216':
        #         records.append(record)
            # cache[tuple(ms_key)] = record
        # else:
        #     print record
        #     print cache[tuple(ms_key)]
        #     time.sleep(60)

    # print count
print len(records)
print records
#39701