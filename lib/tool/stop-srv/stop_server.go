package stop_srv

import (
	"os"
	"os/signal"
	"syscall"
)

func SafelyStopServer(quitFunc func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case s := <-ch:
			switch s {
			case syscall.SIGHUP:
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				quitFunc()
				return
			}
		}
	}
}
