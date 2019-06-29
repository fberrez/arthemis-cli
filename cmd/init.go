package cmd

import (
	"github.com/arthemis-minecraft/arthemis-cli/controller"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes Arthemis server",
	Long: `init initializes the Arthemis Minecraft server.
	
	- Downloads running config
	- Installs required plugins
	`,
	Run: controller.Initialize,
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("target", "plugins", "Directory where plugins will be installed")
	initCmd.Flags().Bool("force", false, "If files already exist, they are replaced")
}
