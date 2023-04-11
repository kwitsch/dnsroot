package rootfile

type RootFile struct {
	Version    string
	LastUpdate string
	Servers    []RootServer
}

type RootServer struct {
	Name       string
	Formerly   string
	OperatedBy string
	A          string
	AAAA       string
}

func (f *RootFile) AppendServer(server *RootServer) {
	if server != nil && len(server.Name) > 0 {
		f.Servers = append(f.Servers, *server)
	}
}
