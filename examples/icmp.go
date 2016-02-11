package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jamescun/tuntap"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("help: icmp.go <interface name|device path>")
		return
	}

	tun, err := tuntap.Tun(os.Args[1])
	if err != nil {
		fmt.Println("error: tun:", err)
		return
	}
	defer tun.Close()

	buf := make([]byte, 1500)
	for {
		n, err := tun.Read(buf)
		if err == tuntap.ErrNotReady {
			fmt.Println("warning: tun: interface not ready, waiting 1s...")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			fmt.Println("error: read:", err)
			break
		}

		fmt.Printf("received: %d [%x]\n", n, buf[:n])
	}
}
