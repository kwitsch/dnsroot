package model

type DNSRoot struct {
	servers []*RootServer
	allA    []string
	allAAAA []string
}

func NewRoot(servers []*RootServer) *DNSRoot {
	allA := []string{}
	allAAAA := []string{}
	for _, s := range servers {
		if len(s.A) > 0 {
			allA = append(allA, s.A)
		}
		if len(s.AAAA) > 0 {
			allAAAA = append(allAAAA, s.AAAA)
		}
	}

	res := DNSRoot{
		servers: servers,
		allA:    allA,
		allAAAA: allAAAA,
	}

	return &res
}

func (dr *DNSRoot) AllServers() []*RootServer {
	return dr.servers
}

func (dr *DNSRoot) AllA() []string {
	return dr.allA
}

func (dr *DNSRoot) AllAAAA() []string {
	return dr.allAAAA
}
