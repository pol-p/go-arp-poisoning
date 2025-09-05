package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Interface string
	VictimIP  string
	SpoofIP   string
	SrcMAC    string
	DstMAC    string
	Count     int
	Infinite  bool
	Interval  int
}

func ParseFlags() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Interface, "i", "en0", "Interfaz de red a usar")
	flag.StringVar(&cfg.VictimIP, "victim-ip", "", "IP de la víctima")
	flag.StringVar(&cfg.SpoofIP, "spoof-ip", "", "IP que quieres suplantar")
	flag.StringVar(&cfg.SrcMAC, "src-mac", "", "MAC de origen (la del atacante)")
	flag.StringVar(&cfg.DstMAC, "dst-mac", "", "MAC de destino (la de la víctima)")
	flag.IntVar(&cfg.Count, "count", 1000, "Cantidad de paquetes a enviar")
	flag.BoolVar(&cfg.Infinite, "infinite", false, "Enviar paquetes ARP de forma infinita")
	flag.IntVar(&cfg.Interval, "interval", 1000, "Intervalo entre paquetes (milisegundos)")

	flag.Parse()

	if cfg.VictimIP == "" || cfg.SpoofIP == "" {
		fmt.Println("[!] Uso: go-arp-poisoning -i INTERFACE -victim-ip IP -spoof-ip IP [-count N]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return cfg
}
