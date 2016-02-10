package tuntap

import (
	"errors"
)

var (
	ErrUnsupported = errors.New("device is unsupported on this platform")
)

// device identifiers
const (
	TUN = 0x0001
	TAP = 0x0002
)

// Interface represents a TUN/TAP network interface
type Interface interface {
	// return name of TUN/TAP interface
	Name() string

	// implement io.Reader interface, read bytes into p from TUN/TAP interface
	Read(p []byte) (n int, err error)

	// implement io.Writer interface, write bytes from p to TUN/TAP interface
	Write(p []byte) (n int, err error)

	// implement io.Closer interface, must be called when done with TUN/TAP interface
	Close() error

	// return string representation of TUN/TAP interface
	String() string
}

// return TUN/TAP interface ready for use
func Open(dev uint16, name string) (Interface, error) {
	switch dev {
	case TUN:
		return newTUN(name)

	case TAP:
		return newTAP(name)
	}

	return nil, ErrUnsupported
}
