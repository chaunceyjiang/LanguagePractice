package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.LUTC)
}

func main() {
	var ethN string
	flag.StringVar(&ethN, "i", "", "抓捕的网卡")
	flag.Parse()
	if ethN == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	var (
		ethLayer layers.Ethernet
		ipLayer  layers.IPv4
		tcpLayer layers.TCP
	)
	devs, err := pcap.FindAllDevs()
	if err != nil {
		log.Println(err)
	}
	var device pcap.Interface
	for _, dev := range devs {
		if dev.Name == ethN {
			device = dev
		}
		//fmt.Println("Addresses",dev.Addresses)
		//fmt.Println("Description",dev.Description)
		//fmt.Println("Flags",dev.Flags)
		//fmt.Println()

	}
	handler, err := pcap.OpenLive(device.Name, 65535, true, 10*time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	packetSource := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &ethLayer, &ipLayer, &tcpLayer)
		foundLayerTypes := []gopacket.LayerType{}
		parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		for _, layerType := range foundLayerTypes {
			if layerType == layers.LayerTypeIPv4 {
				fmt.Println(ethLayer.SrcMAC, ethLayer.DstMAC)
				fmt.Println(ipLayer.SrcIP, ipLayer.DstIP)
				fmt.Println(tcpLayer.SrcPort, tcpLayer.DstPort, tcpLayer.Seq)
				fmt.Println()
				fmt.Println(hex.Dump(tcpLayer.Payload))
				fmt.Println()
			}

		}
	}
}
