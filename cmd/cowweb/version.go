package cowweb

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(version)
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cowweb",
	Long:  `Print the version number of cowweb`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(cowsay.Say(cowsay.Phrase("v1.0.0"), cowsay.Type("default")))
	},
}
