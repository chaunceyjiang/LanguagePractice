import json
import redis
# with open("twitterdata.json", "w+") as twitter_data_file:
# json.dump({"a":{"b":1,"c":["a","d"]}}, twitter_data_file, indent=4,
# sort_keys=True)
r = redis.Redis(host="192.168.1.34", db=2,decode_responses=True)
with open("twitterdata.json", "r") as twitter_data_file:
    r.set("y",twitter_data_file.read())
    # x=json.load(twitter_data_file)
    # print(x["KVXPATH_PROTOCOLS"][0])
    # y = json.dumps(x)

# print(r.get("t"))
x=json.loads(r.get("y"))
print(x["KVXPATH_PROTOCOLS"][0])
# r.set("t",y)
x = {
    "XPATH_PROTOCOLS": [
        [90, "XML", "eXtensible Markup Language"],
        [603001, "WTC XML", "Weblogic Tuxedo Connector with XML USERDATA block"],
        [701, "X XML", "eXtensible Markup Language with one nested support"],
        [702, "XX XML", "eXtensible Markup Language with tow layer nested support"],
        [703, "XXX XML", "eXtensible Markup Language with three layer nested support"],
        [704, "XXXX XML", "eXtensible Markup Language with four layer nested support"],
        [201300001, "Padding XML", "Auto search to export padding and xml"],
        [93, "CDATA XML", "eXtensible Markup Language with one nested support"],
        [41, "SMTP", "Simple Mail Transfer Protocol"],
        [83, "http mime", "Http Multipurpose Internet Mail Extensions"],
        [96, "SEARCH XML", "seach xml rootNode support"],
        [97, "MULTI XML", "Multiple eXtensible Markup Language"],
        [99, "MULTI CDATA XML", "Multiple eXtensible Markup Language with one nested support"], ],

    "KVXPATH_PROTOCOLS": [
        [92, "KEYVALUE XML", "Key-Value base on eXtensible Markup Language"],
        [603002, "WTC KEYVALUE XML", "Weblogic Tuxedo Connector with Key-Value XML USERDATA block"],
        [94, "CDATA KEYVALUE XML",
         "Key-Value base on eXtensible Markup Language with one layer nested support"],
        [711, "X KEYVALUE XML",
         "Key-Value base on eXtensible Markup Language with one layer nested support"],
        [712, "XX KEYVALUE XML",
         "Key-Value base on eXtensible Markup Language with tow layer nested support"],
        [713, "XXX KEYVALUE XML",
         "Key-Value base on eXtensible Markup Language with three layer nested support"],
        [714, "XXXX KEYVALUE XML",
         "Key-Value base on eXtensible Markup Language with four layer nested support"],
        [96, "SEARCH XML", "seach xml rootNode support"],
        [1010, "SEARCH KEYVALUE XML", "merge seach xml and keyvalue xml"],
        [98, "MULTI KEYVALUE XML", "Multiple Key-Value base on eXtensible Markup Language"],
        [100, "MULTI CDATA KEYVALUE XML",
         "Multiple Key-Value base on eXtensible Markup Language with one layer nested support"],
    ],

    "XTREE_PROTOCOLS": [
        [95, "XMLT", "XML Tree Implementation"],
        [731, "X XMLT", "eXtensible Markup Language with one nested support"],
        [732, "XX XMLT", "eXtensible Markup Language with tow layer nested support"],
        [733, "XXX XMLT", "eXtensible Markup Language with three layer nested support"],
        [734, "XXXX XMLT", "eXtensible Markup Language with four layer nested support"],
    ],
    "JSON_PROTOCOLS": [
        [111, "JSon", "JavaScript Object Notation"]],

    "NAME_PROTOCOLS": [
        [114, "HTML", "Hypertext Markup Language"],
        [120, "JSOUP", "Selector Html"],
        [83, "http mime", "Http Multipurpose Internet Mail Extensions"],
        [80, "HTTP", "HyperText Transfer Protocol"],
        [82, "HTTP Data", "Export HTTP Data to file"],
        [81, "HTTP Query", "HTTP URI Parameter"],
        [84, "multipart", "Multipart"],
        [85, "LB HTTP", "Link Break HTTP"],
        [122, "NVC HTTP", "No Version Check HTTP"],
        [124, "NUC HTTP", "No URL Check HTTP"],
        [1008, "SEARCH PATTERN", "search pattern decoder"],
        [101, "Image GIF", "Compuserve GIF"],
        [102, "Image JPEG", "JPEG File Interchange Format"],
        [103, "Image PNG", "Portable Network Graphics"],
        [104, "Image BMP", "Bitmap format"],
        [113, "CSS", "Cascading Style Sheets"],
        [115, "KEYVALUE PAIR", "Key-Value pair protocol"],
        [116, "wraplength", "Wrap Length Protocol"],
        [121, "length begin", "Length Begin Protocol"],
        [123, "skip wraplength", "Skip Wrap Length Protocol"],
        [117, "fix buffer", "Fix Buffer Protocol"],
        [118, "chunk buffer", "Chunk Buffer Protocol"],
        [119, "encoding buffer", "Encoding Buffer Protocol"],
        [42, "POP3", "Post Office Protocol 3"],
        [43, "FTP", "File Transfer Protocol"],
        [44, "FTP Data", "FTP Data"],
        [45, "TELNET", "TELecommunication NETwork (Telnet)"],
        [46, "DNS", "Domain Name System"],
        [61, "IGMP", "Internet Group Management Protocol"],
        [62, "GGP", "Gateway-to-Gateway Protocol"],
        [63, "BGP", "Border Gateway Protocol"],
        [64, "IGRP", "Interior Gateway Routing Protocol"],
        [65, "GRE", "Generic Routing Encapsulation"],
        [221, "SMB", "Server Message Blocks(CIFS)"],
        [222, "SMB2", "Server Message Blocks(CIFS) v2"],
        [301, "RMI", "Java RMI"],
        [302, "GIOP", "General Inter-ORB Protocol"],
        [303, "EJB", "Enterprise JavaBean"],
        [401, "SIP", "Session Initial Protocol"],
        [402, "RTP", "Real-time Transfer Protocol"],
        [403, "RTCP", "Real-time Transfer Control Protocol"],
        [501, "H.323", "H.323 ITU Protocol"],
        [1009, "JOBJ", "Java Object Protocol"],
        [20171115, "CMT", "{P::tag:val:tag:val...}"],
        [600001, "WMQ", "WebSphere MQ"],
        [600014, "WMQ BUFFER", "WebSphere MQ"],
        [20181121, "WMQLen", "WebSphere MQ With Length"],
        [600002, "T3", "Weblogic T3"],
        [600012, "WL RM", "Weblogic Remote"],
        [600003, "WTC", "Weblogic Tuxedo Connector"],
        [603000, "WTC Detector", "To detect Weblogic Tuxedo Connector"],
        [600004, "CPIC", "IBM Common Program Interface for Communication"],
        [600005, "CICS", "IBM Customer Information Control System"],
        [600006, "SAP RFC", "SAP Remote Function Call Protocol"],
        [600007, "CTG", "CICS Transaction Gateway"],
        [600008, "CICS Buffer", "Buffer for multiple CICS messages"],
        [600009, "DCE RPC", "Distributed Computing Environment Remote Procedure Calls"],
        [600010, "spark", "BES SPARK"],
        [600011, "tong easy buffer", "Tong Tech Transaction Middleware"],
        [600013, "AJP", "Apache JServ Protocol"],
        [600015, "JOLT", "Oracle Jolt Protocol"],
        [600016, "TIBCO EAI", "TIBCO Enterprise Application Integration Protocol"],
        [600017, "Dubbo", "Alibaba Dubbo Distributed Service Framework Protocol"],
        [108001, "ABOC NET ESSENTIAL", "Agricultural Bank of China Net Essential Protocol"],
        [181001, "BEA POC", "Bank of East Asia POC Protocol"],
        [180201, "LJB SOP", "LJB SOP"],
        [181502, "BON SOP KEY", "Bank of NANJING ESB Encrypt-Decrypt Key Protocol"],
        [181501, "BON SOP", "Bank of NANJING ESB Protocol"],
        [181101, "BQD ESB", "Bank of Qingdao ESB Protocol"],
        [181801, "BOWF ZJYW", "BOWF ZJYW"],
        [181802, "BOWF OUTER SI",
         "Bank of WeiFang Intermediate Business to Outer Social Insurance Protocol"],
        [181803, "BOWF OUTER UNICOM",
         "Bank of WeiFang Intermediate Business to Outer Unicom Protocol"],
        [180601, "RCB CARRAY", "WH RC Bank CARRAY"],
        [182001, "WHRCB CARRAY", "WH RC Bank CARRAY"],
        [180602, "RCB ATM", "WH RC Bank ATM"],
        [182002, "WHRCB ATM", "WH RC Bank ATM"],
        [181901, "BOWH CDM", "Bank of Wuhan CDM Protocol"],
        [182003, "WHRCB CDM", "Bank of Wuhan CDM Protocol"],
        [182101, "BOYK CORE", "Bank of YingKou Core Protocol"],
        [180801, "BJB CRD CORE", "alias of BOB CRD CORE"],
        [180808, "BOB CRD CORE", "BOB CRD CORE"],
        [180802, "BOB XML", "BOB XML"],
        [180803, "BJB ATM POS", "alias of BOB ATM POS"],
        [180809, "BOB ATM POS", "BOB ATM POS"],
        [180804, "BJB CRD ESB", "alias of BJB CRD ESB"],
        [180807, "BOB CRD ESB", "BOB CRD ESB"],
        [180805, "BOB CBOD", "Bank of Beijing CBOD protocol."],
        [180806, "BOB XM", "Bank of Beijing XOR middleware protocol."],
        [182201, "BSB NPS", "BSB NPS"],
        [181201, "BXB LTTS", "BXB LTTS"],
        [181601, "cdrcb town 2g", "Chengdu Rural Commercial Bank Town Two Generation protocol"],
        [105001, "CCB SNA APP", "China Constuction Bank SNA APP"],
        [105002, "CCB CMBI", "China Constuction Bank CMBI"],
        [105004, "CCB P6", "China Constuction Bank P6"],
        [103001, "CEB FRONT PLAT", "China Everbright Bank Glod General Front encryption platform"],
        [103002, "CEB FDI", "First Data International Shanghai Processing Centre To China Everbright Bank"],
        [103003, "CEB SOP", "ceb sop"],
        [103004, "CEB ESSC", "ceb essc"],
        [103005, "CEB MY", "ceb my"],
        [101108, "CEB CUPS", "China Union Pay System For CEB"],
        [102001, "CMB CAI", "China Merchants Bank Card Authorize Interface Protocol"],
        [102002, "CMB FEPI", "China Merchants Bank FEPI Protocol"],
        [102003, "CMB DEV", "China Merchants Bank DEV Protocol"],
        [102019, "CMB DEV OLD", "China Merchants Bank DEV Protocol"],
        [102004, "CMB FEPI P4W", "China Merchants Bank FEPI P4W Protocol"],
        [102006, "CMB MSGBINARY", "China Merchants Bank MSG Binary"],
        [102007, "CMB GOLD FRONT", "China Merchants Bank Glod Exchange Front Protocol"],
        [102008, "CMB GOLD SECOND STAGE",
         "China Merchants Bank Glod Exchange Second Stage Protocol"],
        [102017, "CMB GOLD FRONT LEVEL2",
         "China Merchants Bank Glod Exchange Second Stage Protocol"],
        [102009, "CMB CUSTODY", "China Merchants Bank store manager"],
        [102010, "CMB SNA CUSTODY", "China Merchants Bank Sna store manager"],
        [102011, "CMB SA", "China Merchants Bank sa"],
        [102012, "CMB CTG", "China Merchants Bank CTG"],
        [102018, "CMB CTGAPP", "China Merchants Bank CTG"],
        [102013, "CMB WASACE", "China Merchants Bank WASACE"],
        [102014, "CMB CMBA", "China Merchants Bank CMBA"],
        [102015, "CMB MBD", "China Merchants Bank MBD"],
        [102016, "CMB 10", "China Merchants Bank 10"],
        [102021, "CMB UNIGATEWAY", "China Merchants Bank UNIGATEWAY"],
        [102022, "CMB APP4", "China Merchants Bank APP4"],
        [102020, "CMB BSERVER", "China Merchants Bank BSERVERs"],
        [102023, "CMB APP3", "China Merchants Bank APP3"],
        [102024, "CMB EBCDIC4", "China Merchants Bank EBCDIC4"],
        [102025, "CMB ATMP", "China Merchants Bank ATMP"],
        [102026, "CMB SIGN", "China Merchants Bank SIGN"],
        [107002, "CUPDATA TIP BUFFER", "Buffer for multiple China UnionPay Tip messages"],
        [107003, "CUPDATA ISO", "China UnionPay Data ISO protocol"],
        [107004, "CUPDATA FIX", "China UnionPay Data FIX protocol"],
        [107005, "CUPDATA", "China UnionPay Data protocol V1"],
        [107006, "CUPDATA V1", "China UnionPay Data protocol V1"],
        [107007, "CUPDATA V2", "China UnionPay Data protocol V2"],
        [106001, "AS400", "ALIAS OF CITIC ABS TO AS400 COUNTER"],
        [106004, "AS400 COUNTER", "CITIC ABS TO AS400 COUNTER"],
        [106005, "AS400 MIDDLE", "CITIC ABS TO AS400 MIDDLE BUSINESS"],
        [106006, "AS400 CARD", "CITIC ABS TO AS400 CARD"],
        [106009, "AS400 SCT", "CITIC ABS TO AS400 SCT"],
        [106002, "citic acq ebcdic", "CITIC Acquiring Ebcdic"],
        [106003, "citic acq ascii", "CITIC Acquiring Ascii"],
        [106007, "citic cmb", "Citic Cascade Mailbox Protocol."],
        [106008, "citic abs", "Citic ABS Protocol."],
        [701001, "CNCC", "China National Clearing Center"],
        [701003, "CNCC 2G", "China National Clearing Center (second generation)"],
        [701002, "ECDS", "Electronic Commercial Draft System"],
        [701004, "CNCC FC", "CNCC Foreign Currency"],
        [701005, "CNCC SYS", "CNCC System"],
        [990301, "CUMS", "China Unionpay Merchant Service"],
        [101101, "CUPS", "China Union Pay System"],
        [101105, "CUPS POS", "China Union Pay System POS"],
        [101106, "CUPS NLC", "China Union Pay System No Length Check"],
        [101107, "CUPS NP", "China Union Pay System No Prolog"],
        [101103, "VISA", "Virtual InstrumentSoftware Architecture"],
        [101104, "MC", "Master Card"],
        [101102, "JIPS", "JCB Interface Processing System"],
        [990208, "GDNX BCSS", "GuangDong NongXin BCSS"],
        [990209, "GDNX SHUTONG", "GuangDong NongXin SHUTONG"],
        [990210, "CdrcbAtm", "ChengDu NongXin Atm"],
        [990206, "GZNSH SOP", "GuangZhou NongShangHang IPP TO Sop"],
        [990207, "GZNSH PAY", "GuangZhou NongShangHang IPP TO Sop"],
        [181701, "BOHN NX CORE", "Bank of Hainan Nongxin CORE Protocol"],
        [180901, "HKB 8583", "8583 Protocol of Han Kou Bank"],
        [180902, "HKB ATMPOS", "ATM/POS Protocol of Han Kou Bank"],
        [180903, "HKB WMP", "WMP Protocol of Han Kou Bank"],
        [180904, "HKB FILE TRANSFER", "AS400 File Transfer Protocol of Han Kou Bank"],
        [180905, "HKB FINANCE", "Finance Protocol of Han Kou Bank"],
        [180102, "HXB PRE APP", "HXB PRE APP"],
        [180101, "HXB MOBILE BANKING", "Hua xia bank mobile banking protocol"],
        [180103, "HXB TOPS", "Hua xia bank tops protocol"],
        [180104, "HXB ONLINE BANKING", "Hua xia bank online banking protocol"],
        [180105, "HXB DCMS", "Hua xia bank DCMS protocol"],
        [180106, "HXB EBS", "Hua xia bank EBS protocol"],
        [101201, "CBOD", "IBM Core Bank On Demand protocol"],
        [101202, "CBOD MM", "IBM Core Bank On Demand protocol base CICS BUFFER support MPMM."],
        [180701, "ICBC DSR", "ICBC DSR"],
        [104001, "CMBC BPP", "China Mingsheng Banking Crop. BPP"],
        [104002, "CMBC EUSP", "China Mingsheng Banking Crop. eBIS Universal Server Platform"],
        [104004, "CMBC EUSP C4",
         "China Mingsheng Banking Crop. eBIS Universal Server Platform Union Pay"],
        [104005, "CMBC EUSP D5",
         "China Mingsheng Banking Crop. eBIS Universal Server Platform Big Pay"],
        [104006, "CMBC EUSP D7", "China Mingsheng Banking Crop. eBIS Universal Server Platform D7"],
        [104007, "CMBC EUSP D8", "China Mingsheng Banking Crop. eBIS Universal Server Platform D8"],
        [104008, "CMBC EUSP D9", "China Mingsheng Banking Crop. eBIS Universal Server Platform Gold"],
        [104009, "CMBC F3 HIGH", "China Mingsheng Banking Crop. F3 high value payment"],
        [104010, "CMBC F3 SMALL", "China Mingsheng Banking Crop. F3 small amount payment"],
        [104011, "CMBC THIRDPARTY PAY", "China Mingsheng Banking Crop. Third party payment"],
        [104012, "CMBC CSP", "China Mingsheng Banking Crop. Client Security Platform"],
        [104013, "CMBC IBP", "China Mingsheng Banking Crop. Intermediary Business Platform"],
        [104014, "CMBC H3 HIGH", "China Mingsheng Banking Crop. H3 high value payment"],
        [104015, "CMBC FUND", "China Mingsheng Banking Crop. fund"],
        [104016, "CMBC H3 HIGH CICS", "China Mingsheng Banking Crop. H3 high value payment"],
        [104017, "CMBC XBANK IC", "China Mingsheng Banking Crop. XBACK IC"],
        [104018, "CMBC HKNB", "China Mingsheng Banking Crop. HK Bank."],
        [101301, "FIX", "Financial Information eXchange protocol."],
        [104003, "CMBC FIX", "China Mingsheng Banking Crop. TSS-PBSS FIX Similar protocol."],
        [180501, "NBB 8583", "Ningbo Bank 8583 Decoder"],
        [180502, "NBB IS", "Ningbo Bank international settlements Decoder"],
        [181401, "EBANKUNION PAYFREE", "Ebank union pay free protocol."],
        [190001, "PINGAN F5", "PingAn Bank F5 protocol"],
        [190002, "ADTEC", "ADTEC"],
        [190003, "CSRCB AFE", "Changshu Rural Commercial Bank AFE"],
        [190004, "CREDIT", "CREDIT"],
        [180301, "SRCB ICBS", "Shenzhen Rural Commercial Bank IBCS"],
        [180302, "ATMP", "Shenzhen Rural Commercial ATMP."],
        [181301, "TGB TCZF", "ChongQing Three Gorges Bank TongChengZhiFu protocol."],
        [181302, "TGB TUXEDO", "ChongQing Three Gorges Bank TUXEDO protocol."],
        [990402, "XIB_CNPAS2", "XIB CNPAS2"],
        [990101, "XJ ATM POS", "XJ ATM POS"],
        [990102, "xj_nxy_zhyw", "XJ NXY ZongHeYeWu"],
        [990103, "caizheng", "XJ  CaiZheng"],
        [180401, "ZMCB PAY 2G", "Zhejiang Mintai Commercial Bank PAY2G Decoder"],
        [990201, "ATMS", "MULTIPLE ATM"],
        [990202, "JZCORE", "JZ CORE 8583"],
        [990203, "JZTelCharge", "JZ Tel Charge 8583"],
        [990204, "JZIC", "JZ IC System"],
        [800001, "CCB FUND", "CCB Principal Asset Management Protocol"],
        [300002, "MySQL", "MySQL Database"],
        [300003, "DRDA", "Distributed Relational Database Architecture"],
        [300001, "TDS", "Tabular Data Stream Protocol"],
        [300004, "TNS", "Transparent Network Substrate"],
        [300005, "SQLNET", "SQL*Net"],
        [300006, "MEMCACHE", "Memcached Distributed Memory Caching System Protocol"],
        [400001, "KRB5", "Kerberos Protocol"],
        [500001, "DX Smgp", "China Telecom Short Message Gateway Protocol"],
        [500101, "China Telecom Smgp", "China Telecom Short Message Gateway Protocol"],
        [500002, "LT Sgip", "China Unicom Short Message Gateway Interface Protocol"],
        [500102, "China Unicom Sgip", "China Unicom Short Message Gateway Interface Protocol"],
        [500003, "YD Cmpp", "China Mobile Peer to Peer Protocol"],
        [500103, "China Mobile Cmpp", "China Mobile Peer to Peer Protocol"],
        [500004, "YD Uasp", "China Mobile Unified Accounting Service Protocol"],
        [500104, "China Mobile Uasp", "China Mobile Unified Accounting Service Protocol"],
        [500005, "gmcc was crm", "Guangdong Mobile WAS CRM"],
        [500006, "gmcc billing", "Guangdong Mobile Billing Protocol"],
        [500007, "dcc", "Diameter Credit Control Protocol"],
        [200001, "SSEP", "Shanghai Stock Exchange Protocol"],
        [200002, "secj1", "Security J1"],
        [200003, "step", "Securities Trading Exchange Protocol"],
        [200004, "gtja sip", "Guotai Junan SIP"],
        [200005, "kcxp", "Kingdom Communication eXchange Platform protocol."],
        [200006, "kcbp", "Kingdom Communication Business Platform protocol."],
        [200007, "kcbp data", "Kingdom Communication Business Data Platform protocol."],
        [200008, "FTD", "Futures Trading Data Exchange Protocol"],
    ], }
# with open("twitterdata.json", "w+") as twitter_data_file:
#     json.dump(x, twitter_data_file, indent=4,
# sort_keys=True)