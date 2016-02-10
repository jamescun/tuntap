package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/jamescun/tuntap"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("help: icmp.go <interface name|device path>")
		return
	}

	tun, err := tuntap.Open(tuntap.TUN, os.Args[1])
	if err != nil {
		fmt.Println("error: tun:", err)
		return
	}
	defer tun.Close()

	buf := make([]byte, 1500)
	for {
		n, err := tun.Read(buf)
		if perr, ok := err.(*os.PathError); ok {
			if code, ok := perr.Err.(syscall.Errno); ok {
				if code == 0x5 {
					// interface is not ready (no address assigned)
					fmt.Println("warning: tun: interface not ready, waiting 1s...")

					time.Sleep(1 * time.Second)
					continue
				}
			}

			fmt.Println("error: read:", err)
			return
		} else if err != nil {
			fmt.Printf("error: read: %#v\n", err)
			fmt.Println("error: read:", err)
			return
		}

		fmt.Printf("received: %d [%x]\n", n, buf[:n])
	}
}
