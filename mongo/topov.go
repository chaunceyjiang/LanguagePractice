package main

type Intf struct {
	ServerID string `bson:"server_id"`
	Name string `bson:"name"`
	Completed bool `bson:"completed"`
	CorrelatedTradeGroupName string `bson:"correlated_trade_group_name"`
	IsTwowayTrade bool `bson:"is_twoway_trade"`
	DispName string `bson:"disp_name"`
	IsSlave bool `bson:"is_slave"`
	ID int `bson:"id"`
	GroupID int `bson:"groupId"`
	Weight int `bson:"weight"`
	CorrelatedGroupName string `bson:"correlated_group_name"`
	Captype string `bson:"captype"`
	IsMaster bool `bson:"is_master"`
	IsDoubleLive bool `bson:"is_double_live"`
	NicAddr string `bson:"nic_addr"`
}
type IPS struct {
	Ips []string `bson:"ips"`
	Ports []string `bson:"ports"`
	Name string `bson:"name"`
}
type Unit struct {
	Completed bool `bson:"completed"`
	DispName string `bson:"disp_name"`
	Pos int `bson:"pos"`
	Name string `bson:"name"`
	Area string `bson:"area"`
}
type Connprop struct {
	Imgname string `bson:"imgname"`
	MatchType bool `bson:"match_type"`
	Name string `bson:"name"`
	Completed bool `bson:"completed"`
	Pos int `bson:"pos"`
	IPAddress []IPS `bson:"ip_address"`
	Side string `bson:"side"`
	Protocol string `bson:"protocol"`
}
type Component struct {
	Comtype string `bson:"comtype"`
	Imgname string `bson:"imgname"`
	Name string `bson:"name"`
	Top int `bson:"top"`
	Left int `bson:"left"`
	Completed bool `bson:"completed"`
	Units []Unit `bson:"units"`
	Connprops []Connprop `bson:"connprops"`
}
type Level struct {
	Name string `bson:"name"`
	DispName string `bson:"disp_name"`
	ID int `bson:"id"`
	GroupID int `bson:"groupId"`
	Weight int `bson:"weight"`
	Top int `bson:"top"`
	Left int `bson:"left"`
	Completed bool `bson:"completed"`
	Components []Component `bson:"components"`
}
type Area struct {
	Name string `bson:"name"`
	DispName string `bson:"disp_name"`
	Top int `bson:"top"`
	Left int `bson:"left"`
	Completed bool `bson:"completed"`
	Levels []Level `bson:"levels"`
}


type Topov struct {
	Name string `bson:"name"`
	State string `bson:"state"`
	ApplyTs int `bson:"apply_ts"`
	Areas []Area `bson:"areas"`
	Connections []struct {
		Intf string `bson:"intf"`
		Connprops []struct {
			Level struct {
				Top int `bson:"top"`
				Left int `bson:"left"`
			} `bson:"level"`
			Component struct {
				Top int `bson:"top"`
				Left int `bson:"left"`
			} `bson:"component"`
			Connprop struct {
				Side string `bson:"side"`
				Pos int `bson:"pos"`
			} `bson:"connprop"`
			Area struct {
				Top int `bson:"top"`
				Left int `bson:"left"`
			} `bson:"area"`
		} `bson:"connprops"`
	} `bson:"connections"`
	Creater string `bson:"creater"`
	Cts float64 `bson:"cts"`
	Desc string `bson:"desc"`
	DispName string `bson:"disp_name"`
	Groups []interface{} `bson:"groups"`
	HasDimension bool `bson:"has_dimension"`
	HasGroup bool `bson:"has_group"`
	HasSite bool `bson:"has_site"`
	Icon string `bson:"icon"`
	Intfs []Intf `bson:"intfs"`
	BaselineTime string `bson:"baseline_time"`
	DataPrecison string `bson:"data_precison"`
	EnabledBaseline bool `bson:"enabled_baseline"`
}

func (p Topov) GetCaptures() []Intf {
	return p.Intfs
}



type Instances struct {
	InstanceID uint64    `bson:"instance_id"`
	AppName    string `bson:"app_name"`
	IntfName   string `bson:"intf_name"`
}

