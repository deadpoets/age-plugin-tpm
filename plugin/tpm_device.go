//go:build !windows

package plugin

import (
	"github.com/google/go-tpm/tpm2/transport"
	"github.com/google/go-tpm/tpm2/transport/linuxtpm"
)

// openDevice opens /dev/tpmrm0, or tpmPath if one is given.
func openDevice(tpmPath string) (transport.TPMCloser, error) {
	if tpmPath != "" {
		return linuxtpm.Open(tpmPath)
	}
	return linuxtpm.Open("/dev/tpmrm0")
}
