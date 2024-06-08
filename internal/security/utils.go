package security

import (
	"crypto/tls"
	"errors"
	"net"
	"time"
)

func isDisallowedIP(hostIP string) bool {
	ip := net.ParseIP(hostIP)

	return ip.IsMulticast() || ip.IsUnspecified() || ip.IsLoopback() || ip.IsPrivate()
}

func checkDisallowedIP(conn net.Conn) error {
	ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	if isDisallowedIP(ip) {
		conn.Close()

		return errors.New("ip address is not allowed")
	}
	return nil
}

func dialFunc(network, addr string, timeout time.Duration, tlsConfig *tls.Config) (net.Conn, error) {
	var conn net.Conn
	var err error

	dialer := &net.Dialer{Timeout: timeout}

	if tlsConfig != nil {
		conn, err = tls.DialWithDialer(dialer, network, addr, tlsConfig)
		if err != nil {
			return nil, err
		}
		if err := checkDisallowedIP(conn); err != nil {
			return nil, err
		}
		return conn, nil
	} else {
		conn, err = dialer.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		if err := checkDisallowedIP(conn); err != nil {
			return nil, err
		}
		return conn, nil
	}
}
