package controllers

func (s *Server) initializeRoutes() {
	r := s.Router

	r.GET("/ping", s.Ping)

	// User Routes
	r.POST("/user", s.CreateUser)
	r.GET("/users", s.GetUsers)
	r.PUT("/users/:id", s.UpdateUser)
}
