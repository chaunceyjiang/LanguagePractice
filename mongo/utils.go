package main

func (topov *Topov)GetTwoWayTrandeIntfIPS( intfName string) map[string][]string {
	intfs := topov.GetCaptures()
	var intfs_list []Intf
	for _, intf := range intfs {
		if intf.IsTwowayTrade {
			intfs_list = append(intfs_list, intf)
		}
	}
	var master Intf
	var slave Intf // 目前只处理一个slave
	for _, intf := range intfs_list {
		if intf.IsMaster && intf.Name == intfName {
			master = intf
			for _, i := range intfs {
				if intf.CorrelatedTradeGroupName == i.CorrelatedTradeGroupName && i.IsSlave {
					slave = i

				}
			}
		}
		if intf.IsSlave && intf.Name == intfName {
			for _, i := range intfs {
				if intf.CorrelatedTradeGroupName == i.CorrelatedTradeGroupName && i.IsMaster {
					master = i
				}
			}
			slave = intf
		}
	}

	result := make(map[string][]string)
	result["master"] = topov.GetServerIPS(master)
	result["slave"] = topov.GetServerIPS(slave)
	return result
}

func (topov Topov) GetServerIPS(intf Intf) []string {
	var ips_list []string
	for _, a := range topov.Areas {
		for _, l := range a.Levels {
			if l.ID == intf.ID {
				for _, comp := range l.Components {
					for _, connp := range comp.Connprops {
						for _, ips := range connp.IPAddress {
							ips_list = append(ips_list, ips.Ips...)
						}
					}
				}
			}

		}
	}
	return ips_list
}
