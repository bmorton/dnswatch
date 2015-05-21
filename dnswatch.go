package main

import (
	"flag"
	"fmt"
	"time"

	"code.google.com/p/gopacket"
	"code.google.com/p/gopacket/layers"
	"code.google.com/p/gopacket/pcap"
)

func main() {
	var deviceInterface string
	flag.StringVar(&deviceInterface, "interface", "eth0", "device interface to listen for DNS packets")
	flag.Parse()

	handle, err := pcap.OpenLive(deviceInterface, 65535, true, 1000)
	if err != nil {
		panic(err)
	}

	err = handle.SetBPFFilter("port 53")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Watching for DNS packets on %s...\n", deviceInterface)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetChan := packetSource.Packets()
	for {
		select {
		case packet := <-packetChan:
			handlePacket(packet)
		}
	}
}

func handlePacket(packet gopacket.Packet) {
	metadata := packet.Metadata()
	dnsLayer := packet.Layer(layers.LayerTypeDNS)
	dns, _ := dnsLayer.(*layers.DNS)
	if len(dns.Answers) == 0 {
		return
	}
	for _, question := range dns.Questions {
		if question.Type == layers.DNSTypeA || question.Type == layers.DNSTypeAAAA {
			fmt.Printf("%s - %s [%d]\n", metadata.Timestamp.Format(time.ANSIC), string(question.Name), dns.ID)
		}
	}
	for _, answer := range dns.Answers {
		if answer.Type == layers.DNSTypeA {
			fmt.Printf("\t\t\t    -> A    %s\n", answer.String())
		} else if answer.Type == layers.DNSTypeAAAA {
			fmt.Printf("\t\t\t    -> AAAA %s\n", answer.String())
		}
	}
}
