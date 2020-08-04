package tlsreuse

import (
	"crypto/tls"
	"errors"
	"net"

	"github.com/libp2p/go-reuseport"
)

func Listen(network, address string, config *tls.Config) (net.Listener, error) {
	if config == nil || len(config.Certificates) == 0 &&
		config.GetCertificate == nil && config.GetConfigForClient == nil {
		return nil, errors.New("tls: neither Certificates, GetCertificate, nor GetConfigForClient set in Config")
	}
	l, err := reuseport.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(l, config), nil
}
