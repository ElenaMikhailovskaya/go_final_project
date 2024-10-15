package service

func New(opts ...func(server *Server)) (*Server, error) {
	c := new(Server)
	//c.userMu = new(sync.Map)
	for _, o := range opts {
		o(c)
	}
	return c, nil
}

func WithDatabase(db DB) func(*Server) {
	return func(s *Server) {
		s.DBase = db
	}
}
