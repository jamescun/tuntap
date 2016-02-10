// +build darwin

package tuntap

import (
	"os"
)

func newTUN(name string) (Interface, error) {
	file, err := os.OpenFile(name, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	return &Device{r: file, n: name}, nil
}

func newTAP(name string) (Interface, error) {
	return newTUN(name)
}
