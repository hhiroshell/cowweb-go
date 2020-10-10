package cowweb

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/hhiroshell/cowweb/pkg/api"
	"github.com/hhiroshell/cowweb/pkg/domain/service"
	"github.com/hhiroshell/cowweb/pkg/infrastructure/cowsay"
)

var port *int
var slow *bool
var load *int
var shutdownGracefully *bool

func init() {
	rootCmd.AddCommand(serve)
	port = serve.Flags().IntP("port", "p", 8080, "Port number for the cowweb http server (default: 8080)")
	slow = serve.Flags().BoolP("slow", "s", false, "Start the cowweb http server with slow mode")
	load = serve.Flags().IntP("load", "l", 1024, "CPU load in slow mode")
	shutdownGracefully = serve.Flags().BoolP("shutdown-gracefully", "g", true, "Gracefully shutting down a cowweb http server")
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start the cowweb http server",
	Long:  `Start the cowweb http server`,
	Run: func(cmd *cobra.Command, args []string) {
		var svc service.CowService
		if *slow {
			svc = cowsay.NewSlowCowsay(*load)
		} else {
			svc = cowsay.NewCowsay()
		}
		server, err := api.NewAPIServer(svc, *port)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()

		sig := make(chan os.Signal)
		defer close(sig)
		signal.Notify(sig, syscall.SIGTERM, os.Interrupt)
		<-sig
		if *shutdownGracefully {
			if err := server.Shutdown(context.Background()); err != nil {
				log.Print(err)
			}
		}
	},
}
