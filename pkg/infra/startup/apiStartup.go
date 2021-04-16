package startup

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sygnux.transaction/pkg/infra/api"
	"sygnux.transaction/utils"
)

type apiServerStatup struct {
	ctx context.Context
}

func NewAPIServerStartup(ctx context.Context) *apiServerStatup {
	return &apiServerStatup{
		ctx: ctx,
	}
}

func (s *apiServerStatup) Initialize() {
	cancelContext, cancelFunc := context.WithCancel(s.ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := api.NewServer(cancelContext)
	port := utils.GetEnv("PORT", "3000")
	server.Run(fmt.Sprintf(":%s", port))

	<-c
	if err := server.Shutdown(); err != nil {
		log.Printf("error to shutdown API server %v\n", err)
	}
	cancelFunc()

	log.Println("Waiting 10 seconds to turning Server off")
	time.Sleep(10 * time.Second)
	log.Println("Server turned off")
}
