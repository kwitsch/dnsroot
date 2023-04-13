package model

import "net"

type RootServer struct {
	Name       string
	Formerly   string
	OperatedBy string
	A          string
	AAAA       string
}

func (r *RootServer) IPA() net.IP {
	return net.ParseIP(r.A)
}

func (r *RootServer) IPAAAA() net.IP {
	return net.ParseIP(r.AAAA)
}
