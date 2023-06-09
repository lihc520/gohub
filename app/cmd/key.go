package cmd

import (
	"github.com/lihc520/gohub/pkg/console"
	"github.com/lihc520/gohub/pkg/helpers"
	"github.com/spf13/cobra"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated Key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs, // 不允许传参
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomNumber(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
