# ARP Poisoning Tool

Una herramienta CLI en Go para realizar ataques de ARP spoofing enviando respuestas ARP falsas a una víctima en la red.

## Estructura del Proyecto

```
go-arp-poisoning/
├── cmd/
│   └── go-arp-poisoning/
│       └── main.go              # Main limpio y modular
├── internal/
│   ├── arp/
│   │   └── arp.go              # Lógica ARP (discovery, spoofing)
│   ├── network/
│   │   └── interface.go        # Utilidades de red
│   └── ui/
│       └── output.go           # Mensajes y formateo
├── pkg/
│   └── config/
│       └── config.go           # Configuración y flags
├── bin/                        # Binarios compilados
├── go.mod
├── go.sum
└── README.md
```

## Características

- **Descubrimiento automático de MACs**: Obtiene automáticamente tu MAC local y la MAC de la víctima mediante ARP requests
- **Modo infinito**: Envía paquetes de forma continua con intervalos configurables
- **Interfaz limpia**: Mensajes claros que muestran cada paso del proceso
- **Modular**: Código organizado en módulos reutilizables

## Instalación

```bash
git clone <repository-url>
cd go-arp-poisoning
go mod tidy
go build -o bin/go-arp-poisoning ./cmd/go-arp-poisoning
```

## Uso

### Uso básico (MACs automáticas)
```bash
sudo ./bin/go-arp-poisoning -victim-ip 192.168.1.100 -spoof-ip 192.168.1.1
```

### Uso con parámetros específicos
```bash
sudo ./bin/go-arp-poisoning \
  -i en0 \
  -victim-ip 192.168.1.100 \
  -spoof-ip 192.168.1.1 \
  -count 500
```

### Modo infinito
```bash
sudo ./bin/go-arp-poisoning \
  -victim-ip 192.168.1.100 \
  -spoof-ip 192.168.1.1 \
  -infinite \
  -interval 2000
```

## Parámetros

- `-i string`: Interfaz de red a usar (default "en0")
- `-victim-ip string`: IP de la víctima (requerido)
- `-spoof-ip string`: IP que quieres suplantar (requerido)
- `-src-mac string`: MAC de origen (automática si no se especifica)
- `-dst-mac string`: MAC de destino (automática si no se especifica)
- `-count int`: Cantidad de paquetes a enviar (default 1000)
- `-infinite`: Enviar paquetes de forma infinita
- `-interval int`: Intervalo entre paquetes en milisegundos (default 1000)

## Ejemplo de Salida

```
=== ARP Spoofing Tool ===
[*] Interfaz seleccionada: en0
[*] Buscando interfaz de red...
[+] Interfaz encontrada: en0
[*] Obteniendo MAC local...
[+] MAC local: aa:bb:cc:dd:ee:ff
[*] Buscando MAC de la víctima...
[+] MAC de la víctima encontrada: 11:22:33:44:55:66
[*] Preparando paquete ARP spoof...
[*] Enviando 1000 paquetes ARP spoof...
[+] Paquete ARP enviado 1000 veces
=== Fin del ataque ARP Spoofing ===
```

## Seguridad y Uso Responsable

⚠️ **Advertencia**: Esta herramienta es solo para fines educativos y pruebas de seguridad autorizadas. 

- Solo úsala en redes que poseas o tengas autorización explícita para probar
- El ARP spoofing puede interrumpir el tráfico de red y causar problemas de conectividad
- Asegúrate de cumplir con todas las leyes locales e internacionales

## Dependencias

- [gopacket](https://github.com/google/gopacket): Para construcción y envío de paquetes de red
- Requiere permisos de administrador para acceso a interfaces de red

## Licencia

Este proyecto es solo para fines educativos. Úsalo bajo tu propia responsabilidad.
