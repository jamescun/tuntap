TUN/TAP
=======

tuntap is a native wrapper for interfacing with TUN/TAP network devices in an idiomatic fashion.

Currently supported are Linux and Mac OS X.

    go get github.com/jamescun/tuntap


Configuration
-------------

The configuration required to open a TUN/TAP device varies by platform. The differences are noted below.

### Linux

When creating a TUN/TAP device, Linux expects to be given a name for the new interface, and a new interface will be allocated for it by the kernel module.

    tap, err := Open(tuntap.TAP, "tap0")
    tun, err := Open(tuntap.TUN, "foo1")


### Mac OS X

On startup, the Mac OS X TUN/TAP kernel extension will allocate multiple TUN/TAP devices, up to the maximum number of each. When creating a TUN/TAP device, Mac OS X expects to be given a path to an unused device.

Currently this package will not attempt to find an unused device, however this is planned to be implemented.

    tap, err := Open(tuntap.TAP, "/dev/tap0")
	tun, err := Open(tuntap.TUN, "/dev/tun15")
