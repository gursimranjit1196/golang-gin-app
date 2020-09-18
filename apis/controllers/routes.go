package controllers

func (s *Server) initializeRoutes() {
	r := s.Router

	r.GET("/ping", s.Ping)
}
