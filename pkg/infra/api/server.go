package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"sygnux.transaction/pkg/infra/api/gin_router"
	"sygnux.transaction/pkg/infra/api/mux_router"
	"sygnux.transaction/utils"
)

type server struct {
	cancelContext context.Context
	httpServer    *http.Server
	routes        map[string][]string // key: route and values scope allowed
}

func NewServer(
	cancelContext context.Context,
) *server {
	return &server{
		cancelContext: cancelContext,
		routes:        map[string][]string{},
	}
}

func (s *server) setPersonalHandlerTimeout(enginer http.Handler) {
	timeOutInSeconds, err := strconv.Atoi(utils.GetEnv("HANDLER_TIMEOUT_SECONDS", "0"))
	if err == nil && timeOutInSeconds > 0 {
		logrus.Println("Default handler timeout is", timeOutInSeconds, "seconds")
		s.httpServer.Handler = http.TimeoutHandler(enginer, (time.Duration(timeOutInSeconds) * time.Second), "Request get timeout...")
	}
}

func (s *server) newEnginer() http.Handler {
	var handler http.Handler

	confHandler := utils.GetEnv("ROUTER", "GIN")
	switch confHandler {
	case "GIN":
		handler = gin_router.NewEnginerHandler().InitializeRoutes()
	case "MUX":
		handler = mux_router.NewEnginerHandler().InitializeRoutes()

	default:
		logrus.Warnf("router (%s) not found", confHandler)
		confHandler = "GIN"
		handler = gin_router.NewEnginerHandler().InitializeRoutes()
	}

	logrus.Printf("Setting (%s) as router...", confHandler)
	return handler
}

func (s *server) Run(addr string) {
	engine := s.newEnginer()
	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			logrus.Errorf("error to start server %v", err)
		}
	}()
}

func (s *server) Shutdown() error {
	return s.httpServer.Shutdown(s.cancelContext)
}
