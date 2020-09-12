package cowweb

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/hhiroshell/cowweb/pkg/api"
	"github.com/hhiroshell/cowweb/pkg/domain/service"
	"github.com/hhiroshell/cowweb/pkg/infrastructure/cowsay"
)

var port *int32
var slow *bool
var load *int32

func init() {
	rootCmd.AddCommand(serve)
	port = serve.Flags().Int32P("port", "p", 8080, "Port number for the cowweb http server (default: 8080)")
	slow = serve.Flags().BoolP("slow", "s", false, "Start the cowweb http server with slow mode")
	load = serve.Flags().Int32P("load", "l", 1024, "CPU load in slow mode")
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start the cowweb http server",
	Long:  `Start the cowweb http server`,
	Run: func(cmd *cobra.Command, args []string) {
		var svc service.CowService
		if *slow {
			svc = cowsay.NewSlowCowsay(int(*load))
		} else {
			svc = cowsay.NewCowsay()
		}
		server := api.NewAPIServer(svc)
		log.Fatal(server.ListenAndServe())
	},
}
