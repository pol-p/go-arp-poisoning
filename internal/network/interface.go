package network

import (
	"fmt"
	"net"
)

func GetInterface(name string) (*net.Interface, error) {
	return net.InterfaceByName(name)
}

func GetLocalIPv4(iface *net.Interface) (net.IP, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			return ipnet.IP.To4(), nil
		}
	}

	return nil, fmt.Errorf("no se encontró IP local en la interfaz")
}
