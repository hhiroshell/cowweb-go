package cowweb

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cowweb",
	Short: "The cowweb is a web API version of cowsay",
	Long:  `The cowweb is a web API version of cowsay`,
}

// Execute his is called by main.main().
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
