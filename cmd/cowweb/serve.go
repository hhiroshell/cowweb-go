package cowweb

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		server, error := api.NewAPIServer(svc, *port)
		if error != nil {
			log.Fatal(error)
		}

		if *shutdownGracefully {
			sig := make(chan os.Signal)
			signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
			go func() {
				defer close(sig)
				<-sig
				ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
				server.Shutdown(ctx)
			}()
		}
		log.Fatal(server.ListenAndServe())
	},
}
