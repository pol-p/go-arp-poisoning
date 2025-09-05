package ui

import "fmt"

func PrintBanner() {
	fmt.Println("=== ARP Spoofing Tool ===")
}

func PrintInterface(iface string) {
	fmt.Println("[*] Interfaz seleccionada:", iface)
}

func PrintSearchingInterface() {
	fmt.Println("[*] Buscando interfaz de red...")
}

func PrintInterfaceFound(name string) {
	fmt.Println("[+] Interfaz encontrada:", name)
}

func PrintGettingLocalMAC() {
	fmt.Println("[*] Obteniendo MAC local...")
}

func PrintLocalMAC(mac string) {
	fmt.Println("[+] MAC local:", mac)
}

func PrintUsingProvidedMAC(mac string) {
	fmt.Println("[*] Usando MAC local proporcionada:", mac)
}

func PrintSearchingVictimMAC() {
	fmt.Println("[*] Buscando MAC de la víctima...")
}

func PrintVictimMACFound(mac string) {
	fmt.Println("[+] MAC de la víctima encontrada:", mac)
}

func PrintUsingProvidedVictimMAC(mac string) {
	fmt.Println("[*] Usando MAC de la víctima proporcionada:", mac)
}

func PrintPreparingPacket() {
	fmt.Println("[*] Preparando paquete ARP spoof...")
}

func PrintSendingInfinite(interval int) {
	fmt.Printf("[*] Enviando paquetes ARP spoof cada %d ms (infinito, Ctrl+C para parar)...\n", interval)
}

func PrintSendingCount(count int) {
	fmt.Printf("[*] Enviando %d paquetes ARP spoof...\n", count)
}

func PrintPacketsSent(count int) {
	fmt.Printf("[+] Paquete ARP enviado %d veces\n", count)
}

func PrintFinish() {
	fmt.Println("=== Fin del ataque ARP Spoofing ===")
}

func PrintError(msg string, err error) {
	fmt.Printf("[!] %s: %v\n", msg, err)
}
