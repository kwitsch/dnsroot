package model

import "net"

// RootServer is a struct that represents a root server
type RootServer struct {
	// Official name of the root server
	Name string
	// Former name of the root server(if any)
	Formerly string
	// Organization that operates the root server(if any)
	OperatedBy string
	// IPv4 address of the root server
	A string
	// IPv6 address of the root server
	AAAA string
}

// IPA returns the IPv4 address of the root server
func (r *RootServer) IPA() net.IP {
	return net.ParseIP(r.A)
}

// IPAAAA returns the IPv6 address of the root server
func (r *RootServer) IPAAAA() net.IP {
	return net.ParseIP(r.AAAA)
}
