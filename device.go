package tuntap

import (
	"fmt"
	"io"
)

// Device represents a generic TUN/TAP device
type Device struct {
	d uint16
	r io.ReadWriteCloser
	n string
}

func (d *Device) Name() string                { return d.n }
func (d *Device) Close() error                { return d.r.Close() }
func (d *Device) Read(p []byte) (int, error)  { return d.r.Read(p) }
func (d *Device) Write(p []byte) (int, error) { return d.r.Write(p) }

func (d *Device) String() string {
	if d.d == TUN {
		return fmt.Sprintf("TUN{%s}", d.n)
	} else if d.d == TAP {
		return fmt.Sprintf("TAP{%s}", d.n)
	}

	return ""
}
