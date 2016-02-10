// +build linux

package tuntap

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const cIFF_NO_PI = 0x1000

func newTUN(name string) (Interface, error) {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	ifName, err := createInterface(file.Fd(), name, TUN|cIFF_NO_PI)
	if err != nil {
		return nil, err
	}

	return &Device{r: file, n: ifName}, nil
}

func newTAP(name string) (Interface, error) {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	ifName, err := createInterface(file.Fd(), name, TAP|cIFF_NO_PI)
	if err != nil {
		return nil, err
	}

	return &Device{r: file, n: ifName}, nil
}

type ifReq struct {
	Name  [0x10]byte
	Flags uint16
	pad   [0x28 - 0x10 - 2]byte
}

func createInterface(fd uintptr, ifName string, flags uint16) (createdIFName string, err error) {
	var req ifReq
	req.Flags = flags
	copy(req.Name[:], ifName)
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(&req)))
	if errno != 0 {
		err = errno
		return
	}
	createdIFName = strings.Trim(string(req.Name[:]), "\x00")
	return
}
