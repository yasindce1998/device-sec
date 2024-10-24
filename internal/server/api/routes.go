package api

func (s *Server) setupRoutes() {
    api := s.router.Group("/api/v1")
    {
        api.POST("/commands", s.createCommand)
        api.GET("/commands/:id", s.getCommand)
        api.GET("/ws", s.handleWebSocket)
    }
}