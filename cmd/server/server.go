package main

import (
	"context"

	"sygnux.transaction/pkg/infra/startup"
)

func main() {
	ctx := context.Background()

	startUP := startup.NewAPIServerStartup(ctx)
	startUP.Initialize()
}
