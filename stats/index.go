package main

import "sync"

func AsiicHandler(Num int) string{

	var(

		Str string = ""

		k int

		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z

	)

	//用来匹配的字符A-Z

	Slice := []string{"","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O",

		"P","Q","R","S","T","U","V","W","X","Y","Z"}

	if Num >26 { //数据大于26需要进行拆分

		for {

			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个

			if k == 0 {

				temp = append(temp, 26)

				k = 26

			} else {

				temp = append(temp, k)

			}

			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中

			if Num <= 26{ //小于等于26直接进行匹配，不需要进行数据拆分

				temp = append(temp, Num)

				break

			}

		}

	}else{

		return Slice[Num]

	}

	for _,value := range temp{

		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面

	}

	return Str

}
type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) *pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}
var CityCode = []int{11,
	12,
	13,
	14,
	15,
	21,
	22,
	23,
	31,
	32,
	33,
	34,
	35,
	36,
	37,
	41,
	42,
	43,
	44,
	45,
	46,
	50,
	51,
	52,
	53,
	54,
	61,
	62,
	63,
	64,
	65,
}


var Index = []string{
	"0111",
	"0112",
	"0113",
	"0119",
	"0121",
	"0122",
	"0123",
	"0131",
	"0132",
	"0133",
	"0134",
	"0141",
	"0142",
	"0143",
	"0149",
	"0151",
	"0152",
	"0153",
	"0154",
	"0159",
	"0161",
	"0162",
	"0163",
	"0164",
	"0169",
	"0171",
	"0179",
	"0181",
	"0182",
	"0190",
	"0211",
	"0212",
	"0220",
	"0231",
	"0232",
	"0241",
	"0242",
	"0251",
	"0252",
	"0311",
	"0312",
	"0313",
	"0314",
	"0315",
	"0319",
	"0321",
	"0322",
	"0323",
	"0329",
	"0330",
	"0391",
	"0392",
	"0399",
	"0411",
	"0412",
	"0421",
	"0422",
	"0511",
	"0512",
	"0513",
	"0514",
	"0515",
	"0519",
	"0521",
	"0522",
	"0523",
	"0529",
	"0531",
	"0532",
	"0539",
	"0541",
	"0549",
	"0610",
	"0620",
	"0690",
	"0711",
	"0712",
	"0721",
	"0722",
	"0810",
	"0820",
	"0890",
	"0911",
	"0912",
	"0913",
	"0914",
	"0915",
	"0916",
	"0917",
	"0919",
	"0921",
	"0922",
	"0929",
	"0931",
	"0932",
	"0933",
	"0939",
	"1011",
	"1012",
	"1013",
	"1019",
	"1020",
	"1030",
	"1091",
	"1092",
	"1093",
	"1099",
	"1110",
	"1120",
	"1190",
	"1200",
	"1311",
	"1312",
	"1313",
	"1314",
	"1319",
	"1321",
	"1329",
	"1331",
	"1332",
	"1340",
	"1351",
	"1352",
	"1353",
	"1361",
	"1362",
	"1363",
	"1369",
	"1371",
	"1372",
	"1373",
	"1391",
	"1392",
	"1393",
	"1399",
	"1411",
	"1419",
	"1421",
	"1422",
	"1431",
	"1432",
	"1433",
	"1439",
	"1441",
	"1442",
	"1449",
	"1451",
	"1452",
	"1453",
	"1459",
	"1461",
	"1462",
	"1469",
	"1491",
	"1492",
	"1493",
	"1494",
	"1495",
	"1499",
	"1511",
	"1512",
	"1513",
	"1514",
	"1515",
	"1519",
	"1521",
	"1522",
	"1523",
	"1524",
	"1525",
	"1529",
	"1530",
	"1610",
	"1620",
	"1690",
	"1711",
	"1712",
	"1713",
	"1721",
	"1722",
	"1723",
	"1731",
	"1732",
	"1733",
	"1741",
	"1742",
	"1743",
	"1751",
	"1752",
	"1761",
	"1762",
	"1763",
	"1771",
	"1772",
	"1773",
	"1779",
	"1781",
	"1782",
	"1783",
	"1784",
	"1789",
	"1811",
	"1819",
	"1821",
	"1829",
	"1830",
	"1910",
	"1921",
	"1922",
	"1923",
	"1929",
	"1931",
	"1932",
	"1939",
	"1941",
	"1942",
	"1951",
	"1952",
	"1953",
	"1954",
	"1959",
	"2011",
	"2012",
	"2013",
	"2019",
	"2021",
	"2022",
	"2023",
	"2029",
	"2031",
	"2032",
	"2033",
	"2034",
	"2035",
	"2039",
	"2041",
	"2042",
	"2043",
	"2049",
	"2110",
	"2120",
	"2130",
	"2140",
	"2190",
	"2211",
	"2212",
	"2221",
	"2222",
	"2223",
	"2231",
	"2239",
	"2311",
	"2312",
	"2319",
	"2320",
	"2330",
	"2411",
	"2412",
	"2413",
	"2414",
	"2419",
	"2421",
	"2422",
	"2423",
	"2429",
	"2431",
	"2432",
	"2433",
	"2434",
	"2435",
	"2436",
	"2437",
	"2438",
	"2439",
	"2441",
	"2442",
	"2443",
	"2444",
	"2449",
	"2451",
	"2452",
	"2453",
	"2454",
	"2455",
	"2456",
	"2459",
	"2461",
	"2462",
	"2469",
	"2511",
	"2519",
	"2521",
	"2522",
	"2523",
	"2524",
	"2529",
	"2530",
	"2541",
	"2542",
	"2611",
	"2612",
	"2613",
	"2614",
	"2619",
	"2621",
	"2622",
	"2623",
	"2624",
	"2625",
	"2629",
	"2631",
	"2632",
	"2641",
	"2642",
	"2643",
	"2644",
	"2645",
	"2646",
	"2651",
	"2652",
	"2653",
	"2659",
	"2661",
	"2662",
	"2663",
	"2664",
	"2665",
	"2666",
	"2667",
	"2669",
	"2671",
	"2672",
	"2681",
	"2682",
	"2683",
	"2684",
	"2689",
	"2710",
	"2720",
	"2730",
	"2740",
	"2750",
	"2761",
	"2762",
	"2770",
	"2780",
	"2811",
	"2812",
	"2821",
	"2822",
	"2823",
	"2824",
	"2825",
	"2826",
	"2829",
	"2831",
	"2832",
	"2911",
	"2912",
	"2913",
	"2914",
	"2915",
	"2916",
	"2919",
	"2921",
	"2922",
	"2923",
	"2924",
	"2925",
	"2926",
	"2927",
	"2928",
	"2929",
	"3011",
	"3012",
	"3021",
	"3022",
	"3023",
	"3024",
	"3029",
	"3031",
	"3032",
	"3033",
	"3034",
	"3039",
	"3041",
	"3042",
	"3049",
	"3051",
	"3052",
	"3053",
	"3054",
	"3055",
	"3056",
	"3057",
	"3059",
	"3061",
	"3062",
	"3071",
	"3072",
	"3073",
	"3074",
	"3075",
	"3076",
	"3079",
	"3081",
	"3082",
	"3089",
	"3091",
	"3099",
	"3110",
	"3120",
	"3130",
	"3140",
	"3211",
	"3212",
	"3213",
	"3214",
	"3215",
	"3216",
	"3217",
	"3218",
	"3219",
	"3221",
	"3222",
	"3229",
	"3231",
	"3232",
	"3239",
	"3240",
	"3251",
	"3252",
	"3253",
	"3254",
	"3259",
	"3311",
	"3312",
	"3321",
	"3322",
	"3323",
	"3324",
	"3329",
	"3331",
	"3332",
	"3333",
	"3340",
	"3351",
	"3352",
	"3353",
	"3359",
	"3360",
	"3371",
	"3372",
	"3373",
	"3379",
	"3381",
	"3382",
	"3383",
	"3389",
	"3391",
	"3392",
	"3393",
	"3394",
	"3399",
	"3411",
	"3412",
	"3413",
	"3414",
	"3415",
	"3419",
	"3421",
	"3422",
	"3423",
	"3424",
	"3425",
	"3429",
	"3431",
	"3432",
	"3433",
	"3434",
	"3435",
	"3436",
	"3437",
	"3439",
	"3441",
	"3442",
	"3443",
	"3444",
	"3445",
	"3446",
	"3451",
	"3452",
	"3453",
	"3459",
	"3461",
	"3462",
	"3463",
	"3464",
	"3465",
	"3466",
	"3467",
	"3471",
	"3472",
	"3473",
	"3474",
	"3475",
	"3479",
	"3481",
	"3482",
	"3483",
	"3484",
	"3489",
	"3491",
	"3492",
	"3493",
	"3499",
	"3511",
	"3512",
	"3513",
	"3514",
	"3515",
	"3516",
	"3517",
	"3521",
	"3522",
	"3523",
	"3524",
	"3525",
	"3529",
	"3531",
	"3532",
	"3533",
	"3534",
	"3541",
	"3542",
	"3543",
	"3544",
	"3545",
	"3546",
	"3549",
	"3551",
	"3552",
	"3553",
	"3554",
	"3561",
	"3562",
	"3563",
	"3569",
	"3571",
	"3572",
	"3573",
	"3574",
	"3575",
	"3576",
	"3577",
	"3579",
	"3581",
	"3582",
	"3583",
	"3584",
	"3585",
	"3586",
	"3587",
	"3589",
	"3591",
	"3592",
	"3593",
	"3594",
	"3595",
	"3596",
	"3597",
	"3599",
	"3611",
	"3612",
	"3620",
	"3630",
	"3640",
	"3650",
	"3660",
	"3670",
	"3711",
	"3712",
	"3713",
	"3714",
	"3715",
	"3716",
	"3719",
	"3720",
	"3731",
	"3732",
	"3733",
	"3734",
	"3735",
	"3736",
	"3737",
	"3739",
	"3741",
	"3742",
	"3743",
	"3744",
	"3749",
	"3751",
	"3752",
	"3761",
	"3762",
	"3770",
	"3780",
	"3791",
	"3792",
	"3799",
	"3811",
	"3812",
	"3813",
	"3819",
	"3821",
	"3822",
	"3823",
	"3824",
	"3825",
	"3829",
	"3831",
	"3832",
	"3833",
	"3834",
	"3839",
	"3841",
	"3842",
	"3843",
	"3844",
	"3849",
	"3851",
	"3852",
	"3853",
	"3854",
	"3855",
	"3856",
	"3857",
	"3859",
	"3861",
	"3862",
	"3869",
	"3871",
	"3872",
	"3873",
	"3874",
	"3879",
	"3891",
	"3899",
	"3911",
	"3912",
	"3913",
	"3914",
	"3915",
	"3919",
	"3921",
	"3922",
	"3931",
	"3932",
	"3933",
	"3934",
	"3939",
	"3940",
	"3951",
	"3952",
	"3953",
	"3961",
	"3962",
	"3963",
	"3964",
	"3969",
	"3971",
	"3972",
	"3973",
	"3974",
	"3975",
	"3976",
	"3979",
	"3981",
	"3982",
	"3983",
	"3984",
	"3985",
	"3989",
	"3990",
	"4011",
	"4012",
	"4013",
	"4014",
	"4015",
	"4016",
	"4019",
	"4021",
	"4022",
	"4023",
	"4024",
	"4025",
	"4026",
	"4027",
	"4028",
	"4029",
	"4030",
	"4040",
	"4050",
	"4090",
	"4111",
	"4119",
	"4120",
	"4190",
	"4210",
	"4220",
	"4310",
	"4320",
	"4330",
	"4341",
	"4342",
	"4343",
	"4349",
	"4350",
	"4360",
	"4390",
	"4411",
	"4412",
	"4413",
	"4414",
	"4415",
	"4416",
	"4417",
	"4419",
	"4420",
	"4430",
	"4511",
	"4512",
	"4513",
	"4520",
	"4610",
	"4620",
	"4630",
	"4690",
	"4710",
	"4720",
	"4790",
	"4811",
	"4812",
	"4813",
	"4814",
	"4819",
	"4821",
	"4822",
	"4823",
	"4831",
	"4832",
	"4833",
	"4834",
	"4839",
	"4840",
	"4851",
	"4852",
	"4853",
	"4861",
	"4862",
	"4863",
	"4871",
	"4872",
	"4873",
	"4874",
	"4875",
	"4879",
	"4891",
	"4892",
	"4893",
	"4899",
	"4910",
	"4920",
	"4991",
	"4999",
	"5011",
	"5012",
	"5013",
	"5021",
	"5022",
	"5030",
	"5090",
	"5111",
	"5112",
	"5113",
	"5114",
	"5115",
	"5116",
	"5117",
	"5119",
	"5121",
	"5122",
	"5123",
	"5124",
	"5125",
	"5126",
	"5127",
	"5128",
	"5129",
	"5131",
	"5132",
	"5133",
	"5134",
	"5135",
	"5136",
	"5137",
	"5138",
	"5139",
	"5141",
	"5142",
	"5143",
	"5144",
	"5145",
	"5146",
	"5147",
	"5149",
	"5151",
	"5152",
	"5153",
	"5154",
	"5161",
	"5162",
	"5163",
	"5164",
	"5165",
	"5166",
	"5167",
	"5168",
	"5169",
	"5171",
	"5172",
	"5173",
	"5174",
	"5175",
	"5176",
	"5177",
	"5178",
	"5179",
	"5181",
	"5182",
	"5183",
	"5184",
	"5189",
	"5191",
	"5192",
	"5193",
	"5199",
	"5211",
	"5212",
	"5213",
	"5219",
	"5221",
	"5222",
	"5223",
	"5224",
	"5225",
	"5226",
	"5227",
	"5229",
	"5231",
	"5232",
	"5233",
	"5234",
	"5235",
	"5236",
	"5237",
	"5238",
	"5239",
	"5241",
	"5242",
	"5243",
	"5244",
	"5245",
	"5246",
	"5247",
	"5248",
	"5249",
	"5251",
	"5252",
	"5253",
	"5254",
	"5255",
	"5261",
	"5262",
	"5263",
	"5264",
	"5265",
	"5266",
	"5267",
	"5271",
	"5272",
	"5273",
	"5274",
	"5279",
	"5281",
	"5282",
	"5283",
	"5284",
	"5285",
	"5286",
	"5287",
	"5289",
	"5291",
	"5292",
	"5293",
	"5294",
	"5295",
	"5296",
	"5297",
	"5299",
	"5311",
	"5312",
	"5313",
	"5320",
	"5331",
	"5332",
	"5333",
	"5339",
	"5411",
	"5412",
	"5413",
	"5414",
	"5419",
	"5421",
	"5422",
	"5429",
	"5431",
	"5432",
	"5433",
	"5434",
	"5435",
	"5436",
	"5437",
	"5438",
	"5439",
	"5441",
	"5442",
	"5443",
	"5449",
	"5511",
	"5512",
	"5513",
	"5521",
	"5522",
	"5523",
	"5531",
	"5532",
	"5539",
	"5611",
	"5612",
	"5621",
	"5622",
	"5623",
	"5629",
	"5631",
	"5632",
	"5639",
	"5710",
	"5720",
	"5810",
	"5821",
	"5822",
	"5829",
	"5910",
	"5920",
	"5930",
	"5941",
	"5942",
	"5949",
	"5951",
	"5952",
	"5959",
	"5960",
	"5990",
	"6010",
	"6020",
	"6090",
	"6110",
	"6121",
	"6129",
	"6130",
	"6140",
	"6190",
	"6210",
	"6220",
	"6231",
	"6232",
	"6233",
	"6239",
	"6241",
	"6242",
	"6291",
	"6299",
	"6311",
	"6312",
	"6319",
	"6321",
	"6322",
	"6331",
	"6339",
	"6410",
	"6421",
	"6422",
	"6429",
	"6431",
	"6432",
	"6433",
	"6434",
	"6439",
	"6440",
	"6450",
	"6490",
	"6511",
	"6512",
	"6513",
	"6519",
	"6520",
	"6531",
	"6532",
	"6540",
	"6550",
	"6560",
	"6571",
	"6572",
	"6579",
	"6591",
	"6599",
	"6610",
	"6621",
	"6622",
	"6623",
	"6624",
	"6629",
	"6631",
	"6632",
	"6633",
	"6634",
	"6635",
	"6636",
	"6637",
	"6639",
	"6640",
	"6650",
	"6711",
	"6712",
	"6720",
	"6731",
	"6732",
	"6739",
	"6741",
	"6749",
	"6750",
	"6760",
	"6790",
	"6811",
	"6812",
	"6813",
	"6814",
	"6820",
	"6830",
	"6840",
	"6851",
	"6852",
	"6853",
	"6860",
	"6870",
	"6890",
	"6911",
	"6919",
	"6920",
	"6930",
	"6940",
	"6950",
	"6991",
	"6999",
	"7010",
	"7020",
	"7030",
	"7040",
	"7090",
	"7111",
	"7112",
	"7113",
	"7114",
	"7115",
	"7119",
	"7121",
	"7122",
	"7123",
	"7124",
	"7125",
	"7129",
	"7130",
	"7211",
	"7212",
	"7213",
	"7214",
	"7215",
	"7219",
	"7221",
	"7222",
	"7223",
	"7224",
	"7229",
	"7231",
	"7232",
	"7239",
	"7241",
	"7242",
	"7243",
	"7244",
	"7245",
	"7246",
	"7249",
	"7251",
	"7259",
	"7261",
	"7262",
	"7263",
	"7264",
	"7269",
	"7271",
	"7272",
	"7279",
	"7281",
	"7282",
	"7283",
	"7284",
	"7289",
	"7291",
	"7292",
	"7293",
	"7294",
	"7295",
	"7296",
	"7297",
	"7298",
	"7299",
	"7310",
	"7320",
	"7330",
	"7340",
	"7350",
	"7410",
	"7420",
	"7431",
	"7432",
	"7439",
	"7441",
	"7449",
	"7451",
	"7452",
	"7453",
	"7454",
	"7455",
	"7459",
	"7461",
	"7462",
	"7463",
	"7471",
	"7472",
	"7473",
	"7474",
	"7475",
	"7481",
	"7482",
	"7483",
	"7484",
	"7485",
	"7486",
	"7491",
	"7492",
	"7493",
	"7499",
	"7511",
	"7512",
	"7513",
	"7514",
	"7515",
	"7516",
	"7517",
	"7519",
	"7520",
	"7530",
	"7540",
	"7590",
	"7610",
	"7620",
	"7630",
	"7640",
	"7690",
	"7711",
	"7712",
	"7713",
	"7714",
	"7715",
	"7716",
	"7719",
	"7721",
	"7722",
	"7723",
	"7724",
	"7725",
	"7726",
	"7727",
	"7729",
	"7810",
	"7820",
	"7830",
	"7840",
	"7850",
	"7861",
	"7862",
	"7869",
	"7910",
	"7920",
	"7930",
	"7940",
	"7990",
	"8010",
	"8020",
	"8030",
	"8040",
	"8051",
	"8052",
	"8053",
	"8060",
	"8070",
	"8080",
	"8090",
	"8111",
	"8112",
	"8113",
	"8114",
	"8121",
	"8122",
	"8129",
	"8131",
	"8132",
	"8191",
	"8192",
	"8193",
	"8199",
	"8211",
	"8219",
	"8221",
	"8222",
	"8223",
	"8224",
	"8229",
	"8290",
	"8310",
	"8321",
	"8322",
	"8331",
	"8332",
	"8333",
	"8334",
	"8335",
	"8336",
	"8341",
	"8342",
	"8350",
	"8391",
	"8392",
	"8393",
	"8394",
	"8399",
	"8411",
	"8412",
	"8413",
	"8414",
	"8415",
	"8416",
	"8421",
	"8422",
	"8423",
	"8424",
	"8425",
	"8431",
	"8432",
	"8433",
	"8434",
	"8435",
	"8436",
	"8491",
	"8492",
	"8499",
	"8511",
	"8512",
	"8513",
	"8514",
	"8515",
	"8516",
	"8519",
	"8521",
	"8522",
	"8529",
	"8610",
	"8621",
	"8622",
	"8623",
	"8624",
	"8625",
	"8626",
	"8629",
	"8710",
	"8720",
	"8730",
	"8740",
	"8750",
	"8760",
	"8770",
	"8810",
	"8820",
	"8831",
	"8832",
	"8840",
	"8850",
	"8860",
	"8870",
	"8890",
	"8911",
	"8912",
	"8919",
	"8921",
	"8929",
	"8930",
	"8991",
	"8992",
	"8999",
	"9011",
	"9012",
	"9013",
	"9019",
	"9020",
	"9030",
	"9041",
	"9042",
	"9049",
	"9051",
	"9052",
	"9053",
	"9054",
	"9059",
	"9090",
	"9100",
	"9210",
	"9221",
	"9222",
	"9223",
	"9224",
	"9225",
	"9226",
	"9231",
	"9232",
	"9291",
	"9299",
	"9310",
	"9320",
	"9411",
	"9412",
	"9413",
	"9414",
	"9415",
	"9419",
	"9420",
	"9490",
	"9511",
	"9512",
	"9513",
	"9519",
	"9521",
	"9522",
	"9529",
	"9530",
	"9541",
	"9542",
	"9610",
	"9620",
	"-1",
}