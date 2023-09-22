package bootstrap

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type RunFunc func()

func GracefulShutdown(bootstrap RunFunc) {
	bootstrap()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGSEGV)
	<-stop
	// FIXME: case handling
	log.Println("Application stopped.")
}
