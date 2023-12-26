package server


type Server struct {
	sm      *http.ServeMux
	handler *handlers.Handler
	server  *http.Server
}

/*
**
======================================================
Replace the storage type with actual db driver type
=====================================================
**
*/
func NewServer(storage any) *Server {
	sm := http.NewServeMux()
	h := handlers.NewHandler(storage)
	return &Server{
		sm:      sm,
		handler: h,
		server: &http.Server{
			Addr:    ":8080",
			Handler: sm,
			/*
			   =================================================
			   For more config please refer to http.Server type
			   ================================================
			*/
		},
	}
}

func (s *Server) Run() {

	go s.start()
	s.sm.Handle("/", s.handler)
	slog.Info("Server Sucessully Running")
	/* Handle GraceFullShudown */
	s.graceFullShutDown()
}

func (s *Server) start() {
	if err := s.server.ListenAndServe(); err != nil {
		slog.Error("Internal Error", "Failed to start Server", err.Error())
		os.Exit(1)
	}
}

func (s *Server) graceFullShutDown() {

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	<-signalChan

	slog.Info("Gracefully shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()
	s.server.Shutdown(ctx)
}

