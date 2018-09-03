package tracing

import (
	"github.com/abonec/file_downloader/config"
	"os"
	"fmt"
	"runtime/trace"
	"os/signal"
	"syscall"
)

func Start(cfg config.Config) error {
	if !cfg.Tracing() {
		return nil
	}
	file, err := os.Create("trace.out")
	if err != nil {
		return err
	}
	trace.Start(file)

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigs
		trace.Stop()
		fmt.Println("Interrupting...")
		os.Exit(0)
	}()
	return nil
}
